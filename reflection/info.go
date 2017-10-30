package reflection

import (
	"context"
	"strings"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

type InfoResp struct {
	Types    map[string]*TypeInfo  `json:"types"`
	Packages map[string][]*Service `json:"packages"`
}

type Package struct {
	Name     string     `json:"name"`
	Services []*Service `json:"services"`
}

type Service struct {
	Name    string    `json:"name"`
	Methods []*Method `json:"methods"`
}

type Method struct {
	Name string `json:"name"`

	In  string `json:"in"`
	Out string `json:"out"`

	InStream  bool `json:"in_stream"`
	OutStream bool `json:"out_stream"`
}

type TypeInfo struct {
	Id     int          `json:"id"`
	Fields []*FieldInfo `json:"fields"`
}

type FieldInfo struct {
	Name       string `json:"name"`
	Number     int    `json:"number"`
	IsRepeated bool   `json:"is_repeated"`
	IsRequired bool   `json:"is_required"`
	TypeName   string `json:"type_name"`
	TypeID     int    `json:"type_id"`
}

func GetInfo(ctx context.Context, addr string) (*InfoResp, error) {
	pool := &descPool{}
	if err := pool.connect(ctx, addr); err != nil {
		return nil, err
	}

	defer pool.disconnect()

	services, err := pool.getServicesDescriptors()
	if err != nil {
		return nil, err
	}

	res := &InfoResp{
		Packages: make(map[string][]*Service),
		Types:    make(map[string]*TypeInfo),
	}

	for sname, descr := range services {
		packageName := strings.Split(sname, "/")[0]
		if packageName == "grpc.reflection.v1alpha" {
			continue
		}

		s := &Service{
			Name:    *descr.Name,
			Methods: make([]*Method, 0),
		}
		for _, method := range descr.Method {
			s.Methods = append(s.Methods, &Method{
				Name: method.GetName(),
				In:   method.GetInputType(),
				Out:  method.GetOutputType(),

				InStream:  method.GetClientStreaming(),
				OutStream: method.GetServerStreaming(),
			})
		}

		res.Packages[packageName] = append(res.Packages[packageName], s)
	}

	res.Types = make(map[string]*TypeInfo)
	for k := range pool.getTypes() {
		res.Types[k] = GetTypeInfo(pool, k)
	}

	return res, nil
}

func GetTypeInfo(pool *descPool, typeName string) *TypeInfo {
	desc := pool.getTypeDescriptor(typeName)
	if desc == nil {
		return nil
	}

	info := &TypeInfo{
		Fields: make([]*FieldInfo, 0),
	}

	for _, field := range desc.GetField() {
		label := field.GetLabel()
		info.Fields = append(info.Fields, &FieldInfo{
			Name:       field.GetName(),
			Number:     int(field.GetNumber()),
			TypeName:   field.GetTypeName(),
			TypeID:     int(field.GetType()),
			IsRepeated: label == descriptor.FieldDescriptorProto_LABEL_REPEATED,
			IsRequired: label == descriptor.FieldDescriptorProto_LABEL_REQUIRED,
		})

	}
	return info
}
