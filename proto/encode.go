/*
Dynamic map -> protobuf generator
*/

package proto

import (
	"encoding/binary"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/lazada/grpc-ui/reflection"
)

func Encode(typeInfo map[string]*reflection.TypeInfo, typeName string, data []FieldValue, fieldPath []int) (resBuf []byte, err error) {
	for _, field := range data {
		// get field info

		var fieldInfo *reflection.FieldInfo
		for _, fi := range typeInfo[typeName].Fields {
			if fi.Number == field.Number {
				break
			}
		}
		if fieldInfo == nil {
			return nil, fmt.Errorf("no such field with number: %v", field.Number)
		}


		typeId := descriptor.FieldDescriptorProto_Type(fieldInfo.TypeID)

		wireType := FieldTypeToWireType(typeId)
		if fieldInfo.IsRepeated {
			wireType = proto.WireBytes
		}

		// encode value

		var fieldBuf []byte

		switch typeId {
		case descriptor.FieldDescriptorProto_TYPE_FLOAT:
			fieldBuf, err = encodeFloat(
				field.Value,
				fieldInfo.Number, wireType,
				fieldInfo.IsRepeated,
			)

		case descriptor.FieldDescriptorProto_TYPE_DOUBLE:
			fieldBuf, err = encodeDouble(
				field.Value,
				fieldInfo.Number, wireType,
				fieldInfo.IsRepeated,
			)

		case descriptor.FieldDescriptorProto_TYPE_SINT32:
			fieldBuf, err = encodeSInt32(
				field.Value,
				fieldInfo.Number, wireType,
				fieldInfo.IsRepeated,
			)

		case descriptor.FieldDescriptorProto_TYPE_SINT64:
			fieldBuf, err = encodeSInt64(
				field.Value,
				fieldInfo.Number, wireType,
				fieldInfo.IsRepeated,
			)

		case descriptor.FieldDescriptorProto_TYPE_FIXED32,
			descriptor.FieldDescriptorProto_TYPE_SFIXED32:
			fieldBuf, err = encodeFixed32(
				field.Value,
				fieldInfo.Number, wireType,
				fieldInfo.IsRepeated,
			)

		case descriptor.FieldDescriptorProto_TYPE_FIXED64,
			descriptor.FieldDescriptorProto_TYPE_SFIXED64:
			fieldBuf, err = encodeFixed64(
				field.Value,
				fieldInfo.Number, wireType,
				fieldInfo.IsRepeated,
			)

		case descriptor.FieldDescriptorProto_TYPE_BOOL:
			fieldBuf, err = encodeBool(
				field.Value,
				fieldInfo.Number, wireType,
				fieldInfo.IsRepeated,
			)

		case descriptor.FieldDescriptorProto_TYPE_INT32,
			descriptor.FieldDescriptorProto_TYPE_INT64,
			descriptor.FieldDescriptorProto_TYPE_UINT32,
			descriptor.FieldDescriptorProto_TYPE_UINT64,
			descriptor.FieldDescriptorProto_TYPE_ENUM:
			fieldBuf, err = encodeInt(
				field.Value,
				fieldInfo.Number, wireType,
				fieldInfo.IsRepeated,
			)

		case descriptor.FieldDescriptorProto_TYPE_BYTES:
			fieldBuf, err = encodeBytes(
				field.Value,
				fieldInfo.Number, wireType,
				fieldInfo.IsRepeated,
			)

		case descriptor.FieldDescriptorProto_TYPE_STRING:
			fieldBuf, err = encodeString(
				field.Value,
				fieldInfo.Number, wireType,
				fieldInfo.IsRepeated,
			)

		//case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
		//	if fieldInfo.IsMap {
		//		fieldBuf, err = encodeMap(
		//			field.Value,
		//			fieldInfo.Number, wireType,
		//			typeName,
		//			typeInfo, fieldPath,
		//		)
		//	} else {
		//		fieldBuf, err = encodeMessage(
		//			field.Value,
		//			fieldInfo.Number, wireType,
		//			typeInfo, typeName,fieldPath,
		//			fieldInfo.IsRepeated,
		//		)
		//	}
		//
		}

		if err != nil {
			return
		}

		 // write encoded data to buffer

		if len(fieldBuf) > 0 {
			resBuf = append(resBuf, fieldBuf...)
		}

	}

	return
}

func EncodeTag(n int, wt int) []byte {
	/*
		set field tag: field number and wire type
	*/
	return EncodeVarint(uint64((uint32(n) << 3) | uint32(wt)))
}

func EncodeVarint(x uint64) []byte {
	/*
		encode types:
		- int32
		- int64
		- uint32
		- uint64
	*/
	return proto.EncodeVarint(x)
}

func EncodeFixed32(x uint32) []byte {
	/*
		encode types:
		- fixed32
		- sfixed32
		- float
	*/
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, x)
	return buf
}

func EncodeFixed64(x uint64) []byte {
	/*
		encode types:
		- fixed64
		- sfixed64
		- double
	*/
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, x)
	return buf
}

func EncodeZigzag32(x uint64) []byte {
	/*
		encode types:
		- sint32
	*/
	return EncodeVarint(uint64((uint32(x) << 1) ^ uint32((int32(x) >> 31))))
}

func EncodeZigzag64(x uint64) []byte {
	/*
		encode types:
		- sint64
	*/
	return EncodeVarint(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func EncodeBytes(x []byte) (buf []byte) {
	/*
		encode types:
		- []byte
		- message
	*/
	buf = append(buf, EncodeVarint(uint64(len(x)))...)
	buf = append(buf, x...)
	return
}
