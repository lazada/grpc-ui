/*
Dynamic protobuf -> map generator
*/

package proto

import (
	"encoding/binary"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"

	"github.com/lazada/grpc-ui/reflection"
)

func Decode(typeInfo map[string]*reflection.TypeInfo, typeName string, buf []byte, fieldPath []int) (msg interface{}, err error) {
	// https://developers.google.com/protocol-buffers/docs/encoding
	// Base on:
	// github.com/golang/protobuf/proto/decode.go
	// unmarshalType

	message := map[string]interface{}{}

	index := 0
	shift := 0

	for err == nil && index < len(buf) {

		// parse tag

		var wireType int
		var fieldNumber int
		wireType, fieldNumber, shift = DecodeTag(buf[index:])
		index += shift

		if fieldNumber <= 0 {
			err = fmt.Errorf("Invalid field number: %v", fieldNumber)
			return
		}

		// get field info

		var fieldInfo *reflection.FieldInfo
		fieldInfo, err = findFieldInfoByNumber(typeInfo, typeName, fieldNumber, fieldPath)
		if err != nil {
			return
		}

		// set repeated flag

		var repeatedField bool
		var mapField bool
		var append bool
		if fieldInfo.IsRepeated {
			repeatedField = true
			append = true
		}

		if fieldInfo.IsMap {
			mapField = true
			append = true
		}

		typeId := descriptor.FieldDescriptorProto_Type(fieldInfo.TypeID)

		// validate wire type

		_wireType := FieldTypeToWireType(typeId)
		if repeatedField {
			_wireType = proto.WireBytes
		}

		if wireType != _wireType {
			err = fmt.Errorf("tag wire type %v != field wire type %v", wireType, _wireType)
			return
		}

		// decode message

		var value interface{} = nil
		shift = 0

		switch typeId {
		case descriptor.FieldDescriptorProto_TYPE_FLOAT:
			value, shift, err = decodeFloat(buf[index:], repeatedField)
			index += shift

		case descriptor.FieldDescriptorProto_TYPE_DOUBLE:
			value, shift, err = decodeDouble(buf[index:], repeatedField)
			index += shift

		case descriptor.FieldDescriptorProto_TYPE_SINT32:
			value, shift, err = decodeSInt32(buf[index:], repeatedField)
			index += shift

		case descriptor.FieldDescriptorProto_TYPE_SINT64:
			value, shift, err = decodeSInt64(buf[index:], repeatedField)
			index += shift

		case descriptor.FieldDescriptorProto_TYPE_FIXED32:
			value, shift, err = decodeFixed32(buf[index:], repeatedField)
			index += shift

		case descriptor.FieldDescriptorProto_TYPE_FIXED64:
			value, shift, err = decodeFixed64(buf[index:], repeatedField)
			index += shift

		case descriptor.FieldDescriptorProto_TYPE_SFIXED32:
			value, shift, err = decodeSFixed32(buf[index:], repeatedField)
			index += shift

		case descriptor.FieldDescriptorProto_TYPE_SFIXED64:
			value, shift, err = decodeSFixed64(buf[index:], repeatedField)
			index += shift

		case descriptor.FieldDescriptorProto_TYPE_BOOL:
			value, shift, err = decodeBool(buf[index:], repeatedField)
			index += shift

		case descriptor.FieldDescriptorProto_TYPE_UINT32,
			descriptor.FieldDescriptorProto_TYPE_UINT64:
			value, shift, err = decodeUInt(buf[index:], repeatedField)
			index += shift

		case descriptor.FieldDescriptorProto_TYPE_INT32,
			descriptor.FieldDescriptorProto_TYPE_INT64,
			descriptor.FieldDescriptorProto_TYPE_ENUM:
			value, shift, err = decodeInt(buf[index:], repeatedField)
			index += shift

		case descriptor.FieldDescriptorProto_TYPE_BYTES:
			value, shift, err = decodeBytes(buf[index:])
			index += shift

		case descriptor.FieldDescriptorProto_TYPE_STRING:
			value, shift, err = decodeString(buf[index:])
			index += shift

		case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
			value, shift, err = decodeMessage(buf[index:], fieldInfo.Number, typeName, typeInfo, fieldPath)
			index += shift

		}

		if err != nil {
			return
		}

		if append {
			err = appendValue(&message, fieldInfo.Name, value, mapField)
			if err != nil {
				return
			}
		} else {
			message[fieldInfo.Name] = value
		}

	}

	msg = interface{}(message)

	return
}

func DecodeTag(buf []byte) (wireType int, fieldNumber int, i int) {
	/*
		Decode wire type and field number from tag
	*/

	var tagValue uint64
	tagValue, i = DecodeVarint(buf)

	wireType = int(tagValue & 0x7)
	fieldNumber = int(tagValue >> 3)

	return
}

func DecodeVarint(buf []byte) (x uint64, i int) {
	/*
		decode types:
		- int32
		- int64
		- uint32
		- uint64
	*/
	return proto.DecodeVarint(buf)
}

func DecodeFixed32(buf []byte) (x uint32, i int) {
	/*
		decode types:
		- fixed32
		- sfixed32
		- float
	*/
	x = binary.LittleEndian.Uint32(buf)
	i = 4
	return
}

func DecodeFixed64(buf []byte) (x uint64, i int) {
	/*
		decode types:
		- fixed64
		- sfixed64
		- double
	*/
	x = binary.LittleEndian.Uint64(buf)
	i = 8
	return
}

func DecodeZigzag32(buf []byte) (x uint64, i int) {
	/*
		decode types:
		- sint32
	*/
	var v uint64
	v, i = DecodeVarint(buf)
	x = uint64((uint32(v) >> 1) ^ uint32((int32(v&1)<<31)>>31))
	return
}

func DecodeZigzag64(buf []byte) (x uint64, i int) {
	/*
		decode types:
		- sint64
	*/
	x, i = DecodeVarint(buf)
	x = (x >> 1) ^ uint64((int64(x&1)<<63)>>63)
	return
}

func DecodeBytes(buf []byte) (x []byte, i int) {
	/*
		decode types:
		- []byte
		- message
	*/
	shift := 0

	var len uint64
	len, shift = DecodeVarint(buf[i:])
	i += shift

	x = make([]byte, len)
	copy(x, buf[i:])
	i += int(len)

	return
}
