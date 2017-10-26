/*
grpc.Invoke wrapper
Dynamic in and out structures
*/

package reflection

import (
	"fmt"

	"github.com/jhump/protoreflect/grpcreflect"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	grpcr "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"

	"github.com/golang/protobuf/jsonpb"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic"
	"github.com/jhump/protoreflect/dynamic/grpcdynamic"
	"time"
)

func Invoke(ctx context.Context, serverAddr string, methodPath string, data []byte) (result interface{}, err error) {
	conn, err := grpc.DialContext(ctx, serverAddr, grpc.WithInsecure())
	if err != nil {
		return
	}
	defer conn.Close()

	cl := grpcreflect.NewClient(ctx, grpcr.NewServerReflectionClient(conn))

	services, err := cl.ListServices()
	if err != nil {
		return nil, err
	}

	var currentMd *desc.MethodDescriptor
	for _, s := range services {
		sd, err := cl.ResolveService(s)
		if err != nil {
			return nil, err
		}

		for _, md := range sd.GetMethods() {
			name := md.GetFullyQualifiedName()
			if methodPath == name {
				currentMd = md
				break
			}
		}
	}

	if currentMd == nil {
		return nil, fmt.Errorf("No such method: %v", methodPath)
	}

	stub := grpcdynamic.NewStub(conn)

	msg := dynamic.NewMessage(currentMd.GetInputType())
	err = msg.UnmarshalJSON(data)
	if err != nil {
		return nil, err
	}

	resp, err := stub.InvokeRpc(context.Background(), currentMd, msg)
	if err != nil {
		return nil, err
	}
	m := jsonpb.Marshaler{}
	s, err := m.MarshalToString(resp)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func InvokeStream(ctx context.Context, serverAddr string, methodPath string, data []byte, cb func(data string)) error {
	conn, err := grpc.DialContext(ctx, serverAddr, grpc.WithInsecure(), grpc.WithTimeout(5 * time.Second), grpc.WithBlock())
	if err != nil {
		return err
	}
	defer conn.Close()

	cl := grpcreflect.NewClient(ctx, grpcr.NewServerReflectionClient(conn))

	services, err := cl.ListServices()
	if err != nil {
		return err
	}

	var currentMd *desc.MethodDescriptor
	for _, s := range services {
		sd, err := cl.ResolveService(s)
		if err != nil {
			return err
		}

		for _, md := range sd.GetMethods() {
			name := md.GetFullyQualifiedName()
			if methodPath == name {
				currentMd = md
				break
			}
		}
	}
	if currentMd == nil {
		return fmt.Errorf("No such method: %v", methodPath)
	}

	stub := grpcdynamic.NewStub(conn)

	msg := dynamic.NewMessage(currentMd.GetInputType())
	err = msg.UnmarshalJSON(data)
	if err != nil {
		return err
	}

	m := jsonpb.Marshaler{
		OrigName: true,
	}

	ss, err := stub.InvokeRpcServerStream(ctx, currentMd, msg)
	if err != nil {
		return err
	}

	for {

		msg, err := ss.RecvMsg()
		if err != nil {
			return err
		}
		data, err := m.MarshalToString(msg)
		if err != nil {
			return err
		}

		cb(data)
	}

	return nil

}
