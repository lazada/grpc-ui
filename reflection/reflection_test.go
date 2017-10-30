package reflection

import (
	"context"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/lazada/grpc-ui/reflection/fixtures/simple"
	"google.golang.org/grpc"
	grpcr "google.golang.org/grpc/reflection"
	"net"
	"testing"
	"time"
)

func TestSimpleTypes(t *testing.T) {
	var stub struct {
		simple.SimpleServer
	}

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}

	defer ln.Close()

	s := grpc.NewServer()
	simple.RegisterSimpleServer(s, stub)
	grpcr.Register(s)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go func() {
		if err := s.Serve(ln); err != nil {
			t.Fatal(err)
		}
		defer s.Stop()

		<-ctx.Done()
	}()

	info, err := GetInfo(ctx, ln.Addr().String())
	if err != nil {
		t.Fatalf("GetInfo err: %v", err)
	}
	if _, ok := info.Packages["simple"]; !ok {
		t.Fatalf("Should have package `simple`, got: %+v", info.Packages)
	}

	if len(info.Packages["simple"]) != 1 {
		t.Fatalf("Should have 1 servce, got: %v", len(info.Packages["simple"]))
	}

	simpleService := info.Packages["simple"][0]

	if len(simpleService.Methods) != 1 {
		t.Fatalf("Should have 1 method, got: %v", simpleService.Methods)
	}

	if simpleService.Methods[0].Name != "Test" {
		t.Fatalf("Service `Simple` have `Test` method, got: %v", simpleService.Methods[0].Name)
	}

	if simpleService.Methods[0].In != ".simple.Req" {
		t.Fatalf("`Test` method shoud have `.simple.Req` type , got: %v", simpleService.Methods[0].In)
	}

	if simpleService.Methods[0].Out != ".simple.Res" {
		t.Fatalf("`Test` method shoud have `.simple.Res` type , got: %v", simpleService.Methods[0].Out)
	}

	for _, tp := range []string{".simple.Req", ".simple.Res"} {
		if _, ok := info.Types[tp]; !ok {
			t.Fatalf("Should have `%v` type, got: %+v", tp, info.Types)
		}
	}

	req := info.Types[".simple.Req"]

	fields := []struct {
		Name       string
		Number     int
		TypeID     descriptor.FieldDescriptorProto_Type
		IsRepeated bool
		IsRequired bool
	}{
		{"int32field", 1, descriptor.FieldDescriptorProto_TYPE_INT32, false, false},
		{"int64field", 2, descriptor.FieldDescriptorProto_TYPE_INT64, false, false},
		{"floatfield", 3, descriptor.FieldDescriptorProto_TYPE_FLOAT, false, false},
		{"doublefield", 4, descriptor.FieldDescriptorProto_TYPE_DOUBLE, false, false},
		{"uint32field", 5, descriptor.FieldDescriptorProto_TYPE_UINT32, false, false},
		{"uint64field", 6, descriptor.FieldDescriptorProto_TYPE_UINT64, false, false},
		{"sint32field", 7, descriptor.FieldDescriptorProto_TYPE_SINT32, false, false},
		{"sint64field", 8, descriptor.FieldDescriptorProto_TYPE_SINT64, false, false},
		{"fixed32field", 9, descriptor.FieldDescriptorProto_TYPE_FIXED32, false, false},
		{"fixed64field", 10, descriptor.FieldDescriptorProto_TYPE_FIXED64, false, false},
		{"sfixed32field", 11, descriptor.FieldDescriptorProto_TYPE_SFIXED32, false, false},
		{"sfixed64field", 12, descriptor.FieldDescriptorProto_TYPE_SFIXED64, false, false},
		{"boolfield", 13, descriptor.FieldDescriptorProto_TYPE_BOOL, false, false},
		{"stringfield", 14, descriptor.FieldDescriptorProto_TYPE_STRING, false, false},
		{"bytesfield", 15, descriptor.FieldDescriptorProto_TYPE_BYTES, false, false},
		{"int32Repeated", 16, descriptor.FieldDescriptorProto_TYPE_INT32, true, false},
	}

	if len(fields) != len(req.Fields) {
		t.Fatalf("Should have `%v` fields, got `%v`", len(fields), len(req.Fields) )

	}

	for i, f := range fields {
		if req.Fields[i].Name != f.Name {
			t.Fatalf("Should names `%v`, got: `%v`", f.Name, req.Fields[i].Name)
		}
		if req.Fields[i].Name != f.Name {
			t.Fatalf("Should have type `%v`, got: `%v`", descriptor.FieldDescriptorProto_Type_name[int32(f.TypeID)], descriptor.FieldDescriptorProto_Type_name[int32(req.Fields[i].TypeID)])
		}

		if req.Fields[i].IsRepeated != f.IsRepeated {
			t.Fatalf("Field `%v` should have repeated flag `%v`, got: `%v`", f.Name, f.IsRepeated, req.Fields[i].IsRepeated)
		}
		if req.Fields[i].IsRequired != f.IsRequired {
			t.Fatalf("Should have required flag `%v`, got: `%v`", f.IsRequired, req.Fields[i].IsRequired)
		}
	}
}
