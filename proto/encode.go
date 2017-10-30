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

func Encode(typeInfo map[string]*reflection.TypeInfo, typeName string, msg interface{}, fieldPath []int) (resBuf []byte, err error) {
	// https://developers.google.com/protocol-buffers/docs/encoding
	// Base on:
	// github.com/golang/protobuf/proto/encode.go
	// enc_struct

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Panic in Encode: %v", r)
		}
	}()

	_msg, ok := msg.(map[string]interface{})
	if !ok {
		err = fmt.Errorf("Encode: Conversion from interface{} to map[string]interface{} failed")
		return
	}

	for fieldName, value := range _msg {

		// get field info

		var fieldInfo *reflection.FieldInfo
		fieldInfo, err = findFieldInfoByName(typeInfo, typeName, fieldName, fieldPath)
		if err != nil {
			return
		}

		// set repeated flag

		var repeatedField bool
		var mapField bool

		if fieldInfo.IsRepeated {
			repeatedField = true
		}

		if fieldInfo.IsMap {
			mapField = true
		}

		typeId := descriptor.FieldDescriptorProto_Type(fieldInfo.TypeID)

		wireType := FieldTypeToWireType(typeId)
		if repeatedField {
			wireType = proto.WireBytes
		}

		// encode value

		var fieldBuf []byte

		switch typeId {
		case descriptor.FieldDescriptorProto_TYPE_FLOAT:
			fieldBuf, err = encodeFloat(
				value,
				fieldInfo.Number, wireType,
				repeatedField,
			)

		case descriptor.FieldDescriptorProto_TYPE_DOUBLE:
			fieldBuf, err = encodeDouble(
				value,
				fieldInfo.Number, wireType,
				repeatedField,
			)

		case descriptor.FieldDescriptorProto_TYPE_SINT32:
			fieldBuf, err = encodeSInt32(
				value,
				fieldInfo.Number, wireType,
				repeatedField,
			)

		case descriptor.FieldDescriptorProto_TYPE_SINT64:
			fieldBuf, err = encodeSInt64(
				value,
				fieldInfo.Number, wireType,
				repeatedField,
			)

		case descriptor.FieldDescriptorProto_TYPE_FIXED32,
			descriptor.FieldDescriptorProto_TYPE_SFIXED32:
			fieldBuf, err = encodeFixed32(
				value,
				fieldInfo.Number, wireType,
				repeatedField,
			)

		case descriptor.FieldDescriptorProto_TYPE_FIXED64,
			descriptor.FieldDescriptorProto_TYPE_SFIXED64:
			fieldBuf, err = encodeFixed64(
				value,
				fieldInfo.Number, wireType,
				repeatedField,
			)

		case descriptor.FieldDescriptorProto_TYPE_BOOL:
			fieldBuf, err = encodeBool(
				value,
				fieldInfo.Number, wireType,
				repeatedField,
			)

		case descriptor.FieldDescriptorProto_TYPE_INT32,
			descriptor.FieldDescriptorProto_TYPE_INT64,
			descriptor.FieldDescriptorProto_TYPE_UINT32,
			descriptor.FieldDescriptorProto_TYPE_UINT64,
			descriptor.FieldDescriptorProto_TYPE_ENUM:
			fieldBuf, err = encodeInt(
				value,
				fieldInfo.Number, wireType,
				repeatedField,
			)

		case descriptor.FieldDescriptorProto_TYPE_BYTES:
			fieldBuf, err = encodeBytes(
				value,
				fieldInfo.Number, wireType,
				repeatedField,
			)

		case descriptor.FieldDescriptorProto_TYPE_STRING:
			fieldBuf, err = encodeString(
				value,
				fieldInfo.Number, wireType,
				repeatedField,
			)

		case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
			if mapField {
				fieldBuf, err = encodeMap(
					value,
					fieldInfo.Number, wireType,
					typeName,
					typeInfo, fieldPath,
				)
			} else {
				fieldBuf, err = encodeMessage(
					value,
					fieldInfo.Number, wireType,
					typeInfo, typeName,fieldPath,
					repeatedField,
				)
			}

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
