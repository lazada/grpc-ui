package proto

import (
	"testing"

	"math"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

const TOLERANCE = 0.001

func TestDecodeFloat(t *testing.T) {
	VALUE := -10.01

	wt := FieldTypeToWireType(descriptor.FieldDescriptorProto_TYPE_FLOAT)
	n := 1

	var buf []byte
	var err error

	buf, err = encodeFloat(VALUE, n, wt, false)

	if err != nil {
		t.Fatalf("Error with encodeFloat: %#v", err)
	}

	index := 0
	shift := 0

	var wireType int
	var fieldNumber int

	wireType, fieldNumber, shift = DecodeTag(buf[index:])
	index += shift

	if wt != wireType {
		t.Fatalf("Invalid DecodeTag wireType: %v != %v", wt, wireType)
	}
	if n != fieldNumber {
		t.Fatalf("Invalid DecodeTag fieldNumber: %v != %v", n, fieldNumber)
	}

	var value interface{} = nil

	value, shift, err = decodeFloat(buf[index:], false)
	index += shift

	if err != nil {
		t.Fatalf("Error with decodeFloat: %#v", err)
	}

	_value, _ := value.(float64)

	if diff := math.Abs(_value - VALUE); diff > TOLERANCE {
		t.Fatalf("Invalid decodeFloat result: %#v != %#v", value, VALUE)
	}
}

func TestDecodeFloatRepeated(t *testing.T) {
	VALUE_1 := -20.02
	VALUE_2 := -30.03

	wt := FieldTypeToWireType(descriptor.FieldDescriptorProto_TYPE_FLOAT)
	n := 1

	var buf []byte
	var err error

	buf, err = encodeFloat([]interface{}{VALUE_1, VALUE_2}, n, wt, true)

	if err != nil {
		t.Fatalf("Error with encodeFloat (repeated): %#v", err)
	}

	index := 0
	shift := 0

	var wireType int
	var fieldNumber int

	wireType, fieldNumber, shift = DecodeTag(buf[index:])
	index += shift

	if wt != wireType {
		t.Fatalf("Invalid DecodeTag wireType: %v != %v", wt, wireType)
	}
	if n != fieldNumber {
		t.Fatalf("Invalid DecodeTag fieldNumber: %v != %v", n, fieldNumber)
	}

	var value interface{} = nil

	value, shift, err = decodeFloat(buf[index:], true)
	index += shift

	if err != nil {
		t.Fatalf("Error with decodeFloat (repeated): %#v", err)
	}

	_value, _ := value.([]float64)

	if diff := math.Abs(_value[0] - VALUE_1); diff > TOLERANCE {
		t.Fatalf("Invalid decodeFloat (repeated) result: %#v != %#v", _value[0], VALUE_1)
	}
	if diff := math.Abs(_value[1] - VALUE_2); diff > TOLERANCE {
		t.Fatalf("Invalid decodeFloat (repeated) result: %#v != %#v", _value[1], VALUE_2)
	}
}

func TestDecodeDouble(t *testing.T) {
	VALUE := -100.01

	wt := FieldTypeToWireType(descriptor.FieldDescriptorProto_TYPE_DOUBLE)
	n := 1

	var buf []byte
	var err error

	buf, err = encodeDouble(VALUE, n, wt, false)

	if err != nil {
		t.Fatalf("Error with encodeDouble: %#v", err)
	}

	index := 0
	shift := 0

	var wireType int
	var fieldNumber int

	wireType, fieldNumber, shift = DecodeTag(buf[index:])
	index += shift

	if wt != wireType {
		t.Fatalf("Invalid DecodeTag wireType: %v != %v", wt, wireType)
	}
	if n != fieldNumber {
		t.Fatalf("Invalid DecodeTag fieldNumber: %v != %v", n, fieldNumber)
	}

	var value interface{} = nil

	value, shift, err = decodeDouble(buf[index:], false)
	index += shift

	if err != nil {
		t.Fatalf("Error with decodeDouble: %#v", err)
	}

	_value, _ := value.(float64)

	if diff := math.Abs(_value - VALUE); diff > TOLERANCE {
		t.Fatalf("Invalid decodeDouble result: %#v != %#v", value, VALUE)
	}
}

func TestDecodeDoubleRepeated(t *testing.T) {
	VALUE_1 := -200.02
	VALUE_2 := -300.03

	wt := FieldTypeToWireType(descriptor.FieldDescriptorProto_TYPE_DOUBLE)
	n := 1

	var buf []byte
	var err error

	buf, err = encodeDouble([]interface{}{VALUE_1, VALUE_2}, n, wt, true)

	if err != nil {
		t.Fatalf("Error with encodeDouble (repeated): %#v", err)
	}

	index := 0
	shift := 0

	var wireType int
	var fieldNumber int

	wireType, fieldNumber, shift = DecodeTag(buf[index:])
	index += shift

	if wt != wireType {
		t.Fatalf("Invalid DecodeTag wireType: %v != %v", wt, wireType)
	}
	if n != fieldNumber {
		t.Fatalf("Invalid DecodeTag fieldNumber: %v != %v", n, fieldNumber)
	}

	var value interface{} = nil

	value, shift, err = decodeDouble(buf[index:], true)
	index += shift

	if err != nil {
		t.Fatalf("Error with decodeDouble (repeated): %#v", err)
	}

	_value, _ := value.([]float64)

	if diff := math.Abs(_value[0] - VALUE_1); diff > TOLERANCE {
		t.Fatalf("Invalid decodeDouble (repeated) result: %#v != %#v", _value[0], VALUE_1)
	}
	if diff := math.Abs(_value[1] - VALUE_2); diff > TOLERANCE {
		t.Fatalf("Invalid decodeDouble (repeated) result: %#v != %#v", _value[1], VALUE_2)
	}
}

func TestDecodeBool(t *testing.T) {
	VALUE := true

	wt := FieldTypeToWireType(descriptor.FieldDescriptorProto_TYPE_FLOAT)
	n := 1

	var buf []byte
	var err error

	buf, err = encodeBool(VALUE, n, wt, false)

	if err != nil {
		t.Fatalf("Error with encodeBool: %#v", err)
	}

	index := 0
	shift := 0

	var wireType int
	var fieldNumber int

	wireType, fieldNumber, shift = DecodeTag(buf[index:])
	index += shift

	if wt != wireType {
		t.Fatalf("Invalid DecodeTag wireType: %v != %v", wt, wireType)
	}
	if n != fieldNumber {
		t.Fatalf("Invalid DecodeTag fieldNumber: %v != %v", n, fieldNumber)
	}

	var value interface{} = nil

	value, shift, err = decodeBool(buf[index:], false)
	index += shift

	if err != nil {
		t.Fatalf("Error with decodeBool: %#v", err)
	}

	_value, _ := value.(bool)

	if _value != VALUE {
		t.Fatalf("Invalid decodeBool result: %#v != %#v", value, VALUE)
	}
}

func TestDecodeBoolRepeated(t *testing.T) {
	VALUE_1 := false
	VALUE_2 := true

	wt := FieldTypeToWireType(descriptor.FieldDescriptorProto_TYPE_FLOAT)
	n := 1

	var buf []byte
	var err error

	buf, err = encodeBool([]interface{}{VALUE_1, VALUE_2}, n, wt, true)

	if err != nil {
		t.Fatalf("Error with encodeBool (repeated): %#v", err)
	}

	index := 0
	shift := 0

	var wireType int
	var fieldNumber int

	wireType, fieldNumber, shift = DecodeTag(buf[index:])
	index += shift

	if wt != wireType {
		t.Fatalf("Invalid DecodeTag wireType: %v != %v", wt, wireType)
	}
	if n != fieldNumber {
		t.Fatalf("Invalid DecodeTag fieldNumber: %v != %v", n, fieldNumber)
	}

	var value interface{} = nil

	value, shift, err = decodeBool(buf[index:], true)
	index += shift

	if err != nil {
		t.Fatalf("Error with decodeBool (repeated): %#v", err)
	}

	_value, _ := value.([]bool)

	if _value[0] != VALUE_1 {
		t.Fatalf("Invalid decodeBool (repeated) result: %#v != %#v", _value[0], VALUE_1)
	}
	if _value[1] != VALUE_2 {
		t.Fatalf("Invalid decodeBool (repeated) result: %#v != %#v", _value[1], VALUE_2)
	}
}

func TestDecodeSInt32(t *testing.T) {
	VALUE := -1000

	wt := FieldTypeToWireType(descriptor.FieldDescriptorProto_TYPE_DOUBLE)
	n := 1

	var buf []byte
	var err error

	buf, err = encodeSInt32(float64(VALUE), n, wt, false)

	if err != nil {
		t.Fatalf("Error with encodeSInt32: %#v", err)
	}

	index := 0
	shift := 0

	var wireType int
	var fieldNumber int

	wireType, fieldNumber, shift = DecodeTag(buf[index:])
	index += shift

	if wt != wireType {
		t.Fatalf("Invalid DecodeTag wireType: %v != %v", wt, wireType)
	}
	if n != fieldNumber {
		t.Fatalf("Invalid DecodeTag fieldNumber: %v != %v", n, fieldNumber)
	}

	var value interface{} = nil

	value, shift, err = decodeSInt32(buf[index:], false)
	index += shift

	if err != nil {
		t.Fatalf("Error with decodeSInt32: %#v", err)
	}

	_value, _ := value.(int)

	if _value != VALUE {
		t.Fatalf("Invalid decodeSInt32 result: %#v != %#v", value, VALUE)
	}
}

func TestDecodeSInt32Repeated(t *testing.T) {
	VALUE_1 := -2000
	VALUE_2 := -3000

	wt := FieldTypeToWireType(descriptor.FieldDescriptorProto_TYPE_DOUBLE)
	n := 1

	var buf []byte
	var err error

	buf, err = encodeSInt32([]interface{}{float64(VALUE_1), float64(VALUE_2)}, n, wt, true)

	if err != nil {
		t.Fatalf("Error with encodeSInt32 (repeated): %#v", err)
	}

	index := 0
	shift := 0

	var wireType int
	var fieldNumber int

	wireType, fieldNumber, shift = DecodeTag(buf[index:])
	index += shift

	if wt != wireType {
		t.Fatalf("Invalid DecodeTag wireType: %v != %v", wt, wireType)
	}
	if n != fieldNumber {
		t.Fatalf("Invalid DecodeTag fieldNumber: %v != %v", n, fieldNumber)
	}

	var value interface{} = nil

	value, shift, err = decodeSInt32(buf[index:], true)
	index += shift

	if err != nil {
		t.Fatalf("Error with decodeSInt32 (repeated): %#v", err)
	}

	_value, _ := value.([]int)

	if _value[0] != VALUE_1 {
		t.Fatalf("Invalid decodeSInt32 (repeated) result: %#v != %#v", _value[0], VALUE_1)
	}
	if _value[1] != VALUE_2 {
		t.Fatalf("Invalid decodeSInt32 (repeated) result: %#v != %#v", _value[1], VALUE_2)
	}
}

func TestDecodeSInt64(t *testing.T) {
	VALUE := -1000

	wt := FieldTypeToWireType(descriptor.FieldDescriptorProto_TYPE_DOUBLE)
	n := 1

	var buf []byte
	var err error

	buf, err = encodeSInt64(float64(VALUE), n, wt, false)

	if err != nil {
		t.Fatalf("Error with encodeSInt64: %#v", err)
	}

	index := 0
	shift := 0

	var wireType int
	var fieldNumber int

	wireType, fieldNumber, shift = DecodeTag(buf[index:])
	index += shift

	if wt != wireType {
		t.Fatalf("Invalid DecodeTag wireType: %v != %v", wt, wireType)
	}
	if n != fieldNumber {
		t.Fatalf("Invalid DecodeTag fieldNumber: %v != %v", n, fieldNumber)
	}

	var value interface{} = nil

	value, shift, err = decodeSInt64(buf[index:], false)
	index += shift

	if err != nil {
		t.Fatalf("Error with decodeSInt64: %#v", err)
	}

	_value, _ := value.(int)

	if _value != VALUE {
		t.Fatalf("Invalid decodeSInt64 result: %#v != %#v", value, VALUE)
	}
}

func TestDecodeSInt64Repeated(t *testing.T) {
	VALUE_1 := -2000
	VALUE_2 := -3000

	wt := FieldTypeToWireType(descriptor.FieldDescriptorProto_TYPE_DOUBLE)
	n := 1

	var buf []byte
	var err error

	buf, err = encodeSInt64([]interface{}{float64(VALUE_1), float64(VALUE_2)}, n, wt, true)

	if err != nil {
		t.Fatalf("Error with encodeSInt64 (repeated): %#v", err)
	}

	index := 0
	shift := 0

	var wireType int
	var fieldNumber int

	wireType, fieldNumber, shift = DecodeTag(buf[index:])
	index += shift

	if wt != wireType {
		t.Fatalf("Invalid DecodeTag wireType: %v != %v", wt, wireType)
	}
	if n != fieldNumber {
		t.Fatalf("Invalid DecodeTag fieldNumber: %v != %v", n, fieldNumber)
	}

	var value interface{} = nil

	value, shift, err = decodeSInt64(buf[index:], true)
	index += shift

	if err != nil {
		t.Fatalf("Error with decodeSInt64 (repeated): %#v", err)
	}

	_value, _ := value.([]int)

	if _value[0] != VALUE_1 {
		t.Fatalf("Invalid decodeSInt64 (repeated) result: %#v != %#v", _value[0], VALUE_1)
	}
	if _value[1] != VALUE_2 {
		t.Fatalf("Invalid decodeSInt64 (repeated) result: %#v != %#v", _value[1], VALUE_2)
	}
}

func TestDecodeFixed32(t *testing.T) {
	var VALUE uint32 = 4294967295

	wt := FieldTypeToWireType(descriptor.FieldDescriptorProto_TYPE_DOUBLE)
	n := 1

	var buf []byte
	var err error

	buf, err = encodeFixed32(float64(VALUE), n, wt, false)

	if err != nil {
		t.Fatalf("Error with encodeFixed32: %#v", err)
	}

	index := 0
	shift := 0

	var wireType int
	var fieldNumber int

	wireType, fieldNumber, shift = DecodeTag(buf[index:])
	index += shift

	if wt != wireType {
		t.Fatalf("Invalid DecodeTag wireType: %v != %v", wt, wireType)
	}
	if n != fieldNumber {
		t.Fatalf("Invalid DecodeTag fieldNumber: %v != %v", n, fieldNumber)
	}

	var value interface{} = nil

	value, shift, err = decodeFixed32(buf[index:], false)
	index += shift

	if err != nil {
		t.Fatalf("Error with decodeFixed32: %#v", err)
	}

	_value, _ := value.(uint32)

	if _value != VALUE {
		t.Fatalf("Invalid decodeFixed32 result: %#v != %#v", _value, VALUE)
	}
}

func TestDecodeFixed32Repeated(t *testing.T) {
	var VALUE_1 uint32 = 4294967295
	var VALUE_2 uint32 = 4294967295 - 1

	wt := FieldTypeToWireType(descriptor.FieldDescriptorProto_TYPE_DOUBLE)
	n := 1

	var buf []byte
	var err error

	buf, err = encodeFixed32([]interface{}{float64(VALUE_1), float64(VALUE_2)}, n, wt, true)

	if err != nil {
		t.Fatalf("Error with encodeFixed32 (repeated): %#v", err)
	}

	index := 0
	shift := 0

	var wireType int
	var fieldNumber int

	wireType, fieldNumber, shift = DecodeTag(buf[index:])
	index += shift

	if wt != wireType {
		t.Fatalf("Invalid DecodeTag wireType: %v != %v", wt, wireType)
	}
	if n != fieldNumber {
		t.Fatalf("Invalid DecodeTag fieldNumber: %v != %v", n, fieldNumber)
	}

	var value interface{} = nil

	value, shift, err = decodeFixed32(buf[index:], true)
	index += shift

	if err != nil {
		t.Fatalf("Error with decodeFixed32 (repeated): %#v", err)
	}

	_value, _ := value.([]uint32)

	if _value[0] != VALUE_1 {
		t.Fatalf("Invalid decodeFixed32 (repeated) result: %#v != %#v", _value[0], VALUE_1)
	}
	if _value[1] != VALUE_2 {
		t.Fatalf("Invalid decodeFixed32 (repeated) result: %#v != %#v", _value[1], VALUE_2)
	}
}

func TestDecodeFixed64(t *testing.T) {
	/*
		uint64 -> float64 -> uint64 loss
		==
		var BIG uint64 = 18446744073709551615
		fmt.Printf("%#v %#v", BIG, uint64(float64(BIG)))
		==
		0xffffffffffffffff 0x8000000000000000
	*/

	var VALUE uint64 = 4294967295

	wt := FieldTypeToWireType(descriptor.FieldDescriptorProto_TYPE_DOUBLE)
	n := 1

	var buf []byte
	var err error

	buf, err = encodeFixed64(float64(VALUE), n, wt, false)

	if err != nil {
		t.Fatalf("Error with encodeFixed64: %#v", err)
	}

	index := 0
	shift := 0

	var wireType int
	var fieldNumber int

	wireType, fieldNumber, shift = DecodeTag(buf[index:])
	index += shift

	if wt != wireType {
		t.Fatalf("Invalid DecodeTag wireType: %v != %v", wt, wireType)
	}
	if n != fieldNumber {
		t.Fatalf("Invalid DecodeTag fieldNumber: %v != %v", n, fieldNumber)
	}

	var value interface{} = nil

	value, shift, err = decodeFixed64(buf[index:], false)
	index += shift

	if err != nil {
		t.Fatalf("Error with decodeFixed64: %#v", err)
	}

	_value, _ := value.(uint64)

	if _value != VALUE {
		t.Fatalf("Invalid decodeFixed64 result: %#v != %#v", _value, VALUE)
	}
}

func TestDecodeFixed64Repeated(t *testing.T) {
	var VALUE_1 uint64 = 4294967295
	var VALUE_2 uint64 = 4294967295 - 1

	wt := FieldTypeToWireType(descriptor.FieldDescriptorProto_TYPE_DOUBLE)
	n := 1

	var buf []byte
	var err error

	buf, err = encodeFixed64([]interface{}{float64(VALUE_1), float64(VALUE_2)}, n, wt, true)

	if err != nil {
		t.Fatalf("Error with encodeFixed64 (repeated): %#v", err)
	}

	index := 0
	shift := 0

	var wireType int
	var fieldNumber int

	wireType, fieldNumber, shift = DecodeTag(buf[index:])
	index += shift

	if wt != wireType {
		t.Fatalf("Invalid DecodeTag wireType: %v != %v", wt, wireType)
	}
	if n != fieldNumber {
		t.Fatalf("Invalid DecodeTag fieldNumber: %v != %v", n, fieldNumber)
	}

	var value interface{} = nil

	value, shift, err = decodeFixed64(buf[index:], true)
	index += shift

	if err != nil {
		t.Fatalf("Error with decodeFixed64 (repeated): %#v", err)
	}

	_value, _ := value.([]uint64)

	if _value[0] != VALUE_1 {
		t.Fatalf("Invalid decodeFixed64 (repeated) result: %#v != %#v", _value[0], VALUE_1)
	}
	if _value[1] != VALUE_2 {
		t.Fatalf("Invalid decodeFixed64 (repeated) result: %#v != %#v", _value[1], VALUE_2)
	}
}
