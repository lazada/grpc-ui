package proto

import (
	"context"
	"fmt"
	"github.com/lazada/grpc-ui/reflection"
	"google.golang.org/grpc"
)

type FieldData []struct{
	Val string `json:"val"`
	Number int `json:"number"`
}

func Invoke(ctx context.Context, addr string, packageName, serviceName, methodName string, data FieldData) (interface{}, error){
	info, err := reflection.GetInfo(ctx, addr)
	if err != nil {
		return nil, err
	}
	inType, outType, err := findInType(info, packageName, serviceName, methodName)
	if err != nil {
		return nil, err
	}

	valuesMap := make(map[string]interface{})

	for _, f := range data {
		for _, pf := range info.Types[inType].Fields {
			if f.Number == pf.Number {
				valuesMap[pf.Name] = f.Val
			}
		}
	}
	in := &Message{
		TypeInfo: info.Types,
		TypeName: inType,
		PB: valuesMap,
	}

	out := &Message{
		TypeInfo: info.Types,
		TypeName: outType,
	}

	conn, err := grpc.DialContext(ctx, addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	methodFqdn := fmt.Sprintf("/%s.%s/%s", packageName, serviceName, methodName)

	if err := grpc.Invoke(ctx, methodFqdn, in, out, conn); err != nil {
		return nil, err
	}

	return out.PB, nil
}


func findInType(info *reflection.InfoResp, packageName, serviceName, methodName string) (string, string, error) {
	for _, s := range info.Packages[packageName] {
		if s.Name == serviceName {
			for _, m := range s.Methods {
				if m.Name == methodName {
					return m.In, m.Out, nil
				}
			}
		}
	}
	return "", "", fmt.Errorf("no such method: %v", methodName)
}