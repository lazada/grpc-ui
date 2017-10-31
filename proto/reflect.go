/*
gRPC reflection from soa_manager helpers
*/

package proto

import (
	"context"
	"fmt"
	"github.com/lazada/grpc-ui/reflection"
	"strings"
	"time"
)

func _treeSearch(typeInfo map[string]*reflection.TypeInfo, typeName string, fieldPath []int) (fields []*reflection.FieldInfo, err error) {
	fields = typeInfo[typeName].Fields
	for _, fieldNumber := range fieldPath {

		var found bool

		for _, fi := range fields {
			if fi.Number == fieldNumber {
				fields = typeInfo[fi.TypeName].Fields
				found = true
				break
			}
		}

		if !found {
			err = fmt.Errorf("Invalid field path: %v", fieldPath)
			return
		}
	}

	return
}

func findFieldInfoByName(typeInfo map[string]*reflection.TypeInfo, typeName, fieldName string, fieldPath []int) (fieldInfo *reflection.FieldInfo, err error) {
	var fields []*reflection.FieldInfo
	fields, err = _treeSearch(typeInfo, typeName, fieldPath)
	if err != nil {
		return
	}

	var found bool
	for _, fi := range fields {
		if fi.Name == fieldName {
			fieldInfo = fi
			found = true
			break
		}
	}

	if !found {
		err = fmt.Errorf("Invalid field name: %v", fieldName)
		return
	}

	return
}

func findFieldInfoByNumber(typeInfo map[string]*reflection.TypeInfo, typeName string, fieldNumber int, fieldPath []int) (fieldInfo *reflection.FieldInfo, err error) {
	var fields []*reflection.FieldInfo
	fields, err = _treeSearch(typeInfo, typeName, fieldPath)
	if err != nil {
		return
	}

	var found bool
	for _, fi := range fields {
		if fi.Number == fieldNumber {
			fieldInfo = fi
			found = true
			break
		}
	}

	if !found {
		err = fmt.Errorf("Invalid field number: %v", fieldNumber)
		return
	}

	return
}

func FetchMethodReflection(serverAddr string, methodPath string) (result *reflection.Method, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	services_list, err := reflection.GetInfo(ctx, serverAddr)
	if err != nil {
		return
	}

	_methodPath := strings.Split(methodPath, "/")
	if len(_methodPath) != 3 || _methodPath[0] != "" {
		err = fmt.Errorf("Invalid methodPath param: %v", methodPath)
		return
	}
	serviceName := _methodPath[1]
	funcName := _methodPath[2]

	_serviceName := strings.Split(serviceName, ".")
	if len(_serviceName) != 2 {
		err = fmt.Errorf("Invalid methodPath param: %v", methodPath)
		return
	}
	packageName := _serviceName[0]
	serviceName = _serviceName[1]

	for pName, services := range services_list.Packages{
		for _, s := range services {
			if pName == packageName && s.Name == serviceName {
				for _, m := range s.Methods {
					if m.Name == funcName {
						result = m
					}
				}
			}
		}

	}

	return
}
