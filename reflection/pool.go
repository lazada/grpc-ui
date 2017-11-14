package reflection

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
)

type descPool struct {
	stream     pb.ServerReflection_ServerReflectionInfoClient
	services   map[string]*descriptor.ServiceDescriptorProto
	types      map[string]*descriptor.DescriptorProto
	enum_types map[string]*descriptor.EnumDescriptorProto
	conn       *grpc.ClientConn
}

func (p *descPool) connect(ctx context.Context, addr string) error {

	conn, err := grpc.DialContext(ctx, addr, grpc.WithInsecure())
	if err != nil {
		return err
	}
	p.conn = conn

	c := pb.NewServerReflectionClient(conn)

	p.stream, err = c.ServerReflectionInfo(ctx)
	if err != nil {
		p.disconnect()
		return err
	}

	return nil
}

func (p *descPool) disconnect() {
	if p.conn != nil {
		go p.conn.Close()
	}
}

func (p *descPool) getServicesDescriptors() (map[string]*descriptor.ServiceDescriptorProto, error) {
	err := p.stream.SendMsg(&pb.ServerReflectionRequest{
		MessageRequest: &pb.ServerReflectionRequest_ListServices{},
	})
	if err != nil {
		return nil, err
	}

	serviceResp, err := p.stream.Recv()
	if err != nil {
		return nil, err
	}

	if err := serviceResp.GetErrorResponse(); err != nil {
		return nil, errors.New(err.String())
	}

	for _, service := range serviceResp.GetListServicesResponse().Service {
		p.updateSymbolFile(service.Name)
	}

	return p.services, nil
}

func (p *descPool) updateSymbolFile(name string) error {
	err := p.stream.SendMsg(&pb.ServerReflectionRequest{
		MessageRequest: &pb.ServerReflectionRequest_FileContainingSymbol{
			FileContainingSymbol: name,
		},
	})

	file, err := p.stream.Recv()
	if err != nil {
		return err
	}

	if err := file.GetErrorResponse(); err != nil {
		return errors.New(err.String())
	}

	for _, descData := range file.GetFileDescriptorResponse().FileDescriptorProto {
		if err := p.parseFileDescriptor(descData); err != nil {
			return err
		}
	}

	return nil
}

func (p *descPool) parseFileDescriptor(data []byte) error {
	d := &descriptor.FileDescriptorProto{}
	if err := proto.Unmarshal(data, d); err != nil {
		return err
	}

	if p.services == nil {
		p.services = make(map[string]*descriptor.ServiceDescriptorProto)
	}

	for _, service := range d.Service {
		sname := fmt.Sprintf("%v/%v", d.GetPackage(), service.GetName())
		p.services[sname] = service
	}

	if p.types == nil {
		p.types = make(map[string]*descriptor.DescriptorProto)
	}
	if p.enum_types == nil {
		p.enum_types = make(map[string]*descriptor.EnumDescriptorProto)
	}

	fullName := ""

	pack := d.GetPackage()
	if pack != "" {
		fullName = "."+d.GetPackage()

	}
	p.processEnumDescriptors(fullName, d.EnumType)

	p.processTypeDescriptors(fullName, d.MessageType)

	return nil
}

func (p *descPool) processTypeDescriptors(prefix string, desc []*descriptor.DescriptorProto) {
	for _, d := range desc {
		fullName := prefix + "." + d.GetName()
		p.types[fullName] = d
		p.processTypeDescriptors(fullName, d.NestedType)
		p.processEnumDescriptors(fullName, d.EnumType)
	}
}

func (p *descPool) processEnumDescriptors(prefix string, desc []*descriptor.EnumDescriptorProto) {
	for _, d := range desc {
		fullName := prefix + "." + d.GetName()
		p.enum_types[fullName] = d
	}
}

func (p *descPool) getTypeDescriptor(name string) *descriptor.DescriptorProto {

	// update symbol from protobuf standart
	if strings.HasPrefix(name, ".google.protobuf") {
		p.updateSymbolFile(name[1:])
	}

	if p.types == nil {
		return nil
	}
	if t, ok := p.types[name]; ok {
		return t
	}
	return nil
}

func (p *descPool) getEnumDescriptor(name string) *descriptor.EnumDescriptorProto {
	if p.enum_types == nil {
		return nil
	}
	if t, ok := p.enum_types[name]; ok {
		return t
	}
	return nil
}

func (p *descPool) getTypes() map[string]*descriptor.DescriptorProto {
	return p.types
}

func (p *descPool) getEnums() map[string]*descriptor.EnumDescriptorProto {
	return  p.enum_types
}