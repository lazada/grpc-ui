package main

import (
	"github.com/lazada/grpc-ui/test_server/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"strings"
)

type impl struct {
	pb.ExampleServer
}

func (i *impl) Test(ctx context.Context, req *pb.Req) (*pb.Res, error) {
	return &pb.Res{
		Int32Field:    req.Int32Field,
		Int64Field:    req.Int64Field,
		Floatfield:    req.Floatfield,
		Doublefield:   req.Doublefield,
		Uint32Field:   req.Uint32Field,
		Uint64Field:   req.Uint64Field,
		Sint32Field:   req.Sint32Field,
		Sint64Field:   req.Sint64Field,
		Fixed32Field:  req.Fixed32Field,
		Fixed64Field:  req.Fixed64Field,
		Sfixed32Field: req.Sfixed32Field,
		Sfixed64Field: req.Sfixed64Field,
		Boolfield:     req.Boolfield,
		Stringfield:   req.Stringfield,
		Bytesfield:    req.Bytesfield,
	}, nil
}

func (i *impl) GetUser(ctx context.Context, req *pb.GetUserReq) (*pb.UserResp, error) {
	return &pb.UserResp{
		FirstName: req.FirstName + "[server]",
		LastName: req.LastName + " [server]",
		Active: len(req.FirstName) > 3,
		Age: req.Age,
		Skills: strings.Join(req.Skills, ", "),
	}, nil
}

func start() error {
	ln, err := net.Listen("tcp", ":3001")
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	pb.RegisterExampleServer(s, &impl{})

	reflection.Register(s)

	return s.Serve(ln)
}

func main() {
	if err := start(); err != nil {
		panic(err)
	}
}
