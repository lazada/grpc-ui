package proto

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

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
