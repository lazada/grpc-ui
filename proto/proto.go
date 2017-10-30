/*
Dynamic proto.Message
*/

package proto

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/lazada/grpc-ui/reflection"
)

// Dynamic proto Message

type Message struct {
	TypeInfo map[string]*reflection.TypeInfo
	TypeName string
	PB       map[string]interface{}
}

// proto.Message interface

func (m *Message) Reset() {
	m.PB = map[string]interface{}{}
}

func (m *Message) String() string {
	return fmt.Sprintf("%#v", m.PB)
}

func (m *Message) ProtoMessage() {
}

// Marshaler and Unarshaler interfaces

func (m *Message) Marshal() ([]byte, error) {
	return Encode(m.TypeInfo, m.TypeName, m.PB, nil)
}

func (m *Message) Unmarshal(buf []byte) error {
	res, err := Decode(m.TypeInfo, m.TypeName, buf, nil)
	_res, ok := res.(map[string]interface{})
	if !ok {
		err = fmt.Errorf("Invalid unmarshaled message, err: %v", err)
	}
	m.PB = _res
	return err
}

func FieldTypeToWireType(typeId descriptor.FieldDescriptorProto_Type) (wireType int) {
	switch typeId {

	case descriptor.FieldDescriptorProto_TYPE_FLOAT:
		wireType = proto.WireFixed32

	case descriptor.FieldDescriptorProto_TYPE_DOUBLE:
		wireType = proto.WireFixed64

	case descriptor.FieldDescriptorProto_TYPE_FIXED32,
		descriptor.FieldDescriptorProto_TYPE_SFIXED32:
		wireType = proto.WireFixed32

	case descriptor.FieldDescriptorProto_TYPE_FIXED64,
		descriptor.FieldDescriptorProto_TYPE_SFIXED64:
		wireType = proto.WireFixed64

	case descriptor.FieldDescriptorProto_TYPE_BOOL,
		descriptor.FieldDescriptorProto_TYPE_INT32,
		descriptor.FieldDescriptorProto_TYPE_INT64,
		descriptor.FieldDescriptorProto_TYPE_SINT32,
		descriptor.FieldDescriptorProto_TYPE_SINT64,
		descriptor.FieldDescriptorProto_TYPE_UINT32,
		descriptor.FieldDescriptorProto_TYPE_UINT64,
		descriptor.FieldDescriptorProto_TYPE_ENUM:
		wireType = proto.WireVarint

	case descriptor.FieldDescriptorProto_TYPE_BYTES,
		descriptor.FieldDescriptorProto_TYPE_STRING,
		descriptor.FieldDescriptorProto_TYPE_MESSAGE:
		wireType = proto.WireBytes

	}

	return
}
