package proto

import (
	"testing"

	"bytes"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

func TestEncodeFloat(t *testing.T) {
	wt := FieldTypeToWireType(descriptor.FieldDescriptorProto_TYPE_FLOAT)
	n := 1

	var buf []byte
	var err error

	buf, err = encodeFloat(10.01, n, wt, false)

	if err != nil {
		t.Fatalf("Error with encodeFloat: %#v", err)
	}

	if !bytes.Equal(buf, []byte{0xd, 0xf6, 0x28, 0x20, 0x41}) {
		t.Fatalf("Invalid encodeFloat result: %#v", buf)
	}
}

func TestEncodeFloatRepeated(t *testing.T) {
	wt := FieldTypeToWireType(descriptor.FieldDescriptorProto_TYPE_FLOAT)
	n := 1

	var buf []byte
	var err error

	buf, err = encodeFloat([]interface{}{20.02, 30.03}, n, wt, true)

	if err != nil {
		t.Fatalf("Error with encodeFloat (repeated): %#v", err)
	}

	if !bytes.Equal(buf, []byte{0xd, 0x8, 0xf6, 0x28, 0xa0, 0x41, 0x71, 0x3d, 0xf0, 0x41}) {
		t.Fatalf("Invalid encodeFloat (repeated) result: %#v", buf)
	}
}

func TestEncodeDouble(t *testing.T) {
	wt := FieldTypeToWireType(descriptor.FieldDescriptorProto_TYPE_DOUBLE)
	n := 1

	var buf []byte
	var err error

	buf, err = encodeDouble(10.01, n, wt, false)

	if err != nil {
		t.Fatalf("Error with encodeDouble: %#v", err)
	}

	if !bytes.Equal(buf, []byte{0x9, 0x85, 0xeb, 0x51, 0xb8, 0x1e, 0x5, 0x24, 0x40}) {
		t.Fatalf("Invalid encodeDouble result: %#v", buf)
	}
}

func TestEncodeDoubleRepeated(t *testing.T) {
	wt := FieldTypeToWireType(descriptor.FieldDescriptorProto_TYPE_DOUBLE)
	n := 1

	var buf []byte
	var err error

	buf, err = encodeDouble([]interface{}{20.02, 30.03}, n, wt, true)

	if err != nil {
		t.Fatalf("Error with encodeDouble (repeated): %#v", err)
	}

	if !bytes.Equal(buf, []byte{0x9, 0x10, 0x85, 0xeb, 0x51, 0xb8, 0x1e, 0x5, 0x34, 0x40, 0x48, 0xe1, 0x7a, 0x14, 0xae, 0x7, 0x3e, 0x40}) {
		t.Fatalf("Invalid encodeDouble (repeated) result: %#v", buf)
	}
}

func TestEncodeBool(t *testing.T) {
	wt := FieldTypeToWireType(descriptor.FieldDescriptorProto_TYPE_BOOL)
	n := 1

	var buf []byte
	var err error

	buf, err = encodeBool(true, n, wt, false)

	if err != nil {
		t.Fatalf("Error with encodeBool: %#v", err)
	}

	if !bytes.Equal(buf, []byte{0x8, 0x1}) {
		t.Fatalf("Invalid encodeBool result: %#v", buf)
	}
}

func TestEncodeBoolRepeated(t *testing.T) {
	wt := FieldTypeToWireType(descriptor.FieldDescriptorProto_TYPE_BOOL)
	n := 1

	var buf []byte
	var err error

	buf, err = encodeBool([]interface{}{true, false}, n, wt, true)

	if err != nil {
		t.Fatalf("Error with encodeBool (repeated): %#v", err)
	}

	if !bytes.Equal(buf, []byte{0x8, 0x2, 0x1, 0x0}) {
		t.Fatalf("Invalid encodeBool (repeated) result: %#v", buf)
	}
}

func TestEncodeSInt32(t *testing.T) {
	wt := FieldTypeToWireType(descriptor.FieldDescriptorProto_TYPE_SINT32)
	n := 1

	var buf []byte
	var err error

	buf, err = encodeSInt32(float64(-100), n, wt, false)

	if err != nil {
		t.Fatalf("Error with encodeSInt32: %#v", err)
	}

	if !bytes.Equal(buf, []byte{0x8, 0xc7, 0x1}) {
		t.Fatalf("Invalid encodeSInt32 result: %#v", buf)
	}
}

func TestEncodeSInt32Repeated(t *testing.T) {
	wt := FieldTypeToWireType(descriptor.FieldDescriptorProto_TYPE_SINT32)
	n := 1

	var buf []byte
	var err error

	buf, err = encodeSInt32([]interface{}{float64(-200), float64(300)}, n, wt, true)

	if err != nil {
		t.Fatalf("Error with encodeSInt32 (repeated): %#v", err)
	}

	if !bytes.Equal(buf, []byte{0x8, 0x4, 0x8f, 0x3, 0xd8, 0x4}) {
		t.Fatalf("Invalid encodeSInt32 (repeated) result: %#v", buf)
	}
}

func TestEncodeSInt64(t *testing.T) {
	wt := FieldTypeToWireType(descriptor.FieldDescriptorProto_TYPE_SINT64)
	n := 1

	var buf []byte
	var err error

	buf, err = encodeSInt64(float64(-1000), n, wt, false)

	if err != nil {
		t.Fatalf("Error with encodeSInt64: %#v", err)
	}

	if !bytes.Equal(buf, []byte{0x8, 0xcf, 0xf}) {
		t.Fatalf("Invalid encodeSInt64 result: %#v", buf)
	}
}

func TestEncodeSInt64Repeated(t *testing.T) {
	wt := FieldTypeToWireType(descriptor.FieldDescriptorProto_TYPE_SINT64)
	n := 1

	var buf []byte
	var err error

	buf, err = encodeSInt64([]interface{}{float64(-2000), float64(3000)}, n, wt, true)

	if err != nil {
		t.Fatalf("Error with encodeSInt64 (repeated): %#v", err)
	}

	if !bytes.Equal(buf, []byte{0x8, 0x4, 0x9f, 0x1f, 0xf0, 0x2e}) {
		t.Fatalf("Invalid encodeSInt64 (repeated) result: %#v", buf)
	}
}

func TestEncodeFixed32(t *testing.T) {
	wt := FieldTypeToWireType(descriptor.FieldDescriptorProto_TYPE_FIXED32)
	n := 1

	var buf []byte
	var err error

	buf, err = encodeFixed32(float64(10000), n, wt, false)

	if err != nil {
		t.Fatalf("Error with encodeFixed32: %#v", err)
	}

	if !bytes.Equal(buf, []byte{0xd, 0x10, 0x27, 0x0, 0x0}) {
		t.Fatalf("Invalid encodeFixed32 result: %#v", buf)
	}
}

func TestEncodeFixed32Repeated(t *testing.T) {
	wt := FieldTypeToWireType(descriptor.FieldDescriptorProto_TYPE_FIXED32)
	n := 1

	var buf []byte
	var err error

	buf, err = encodeFixed32([]interface{}{float64(20000), float64(30000)}, n, wt, true)

	if err != nil {
		t.Fatalf("Error with encodeFixed32 (repeated): %#v", err)
	}

	if !bytes.Equal(buf, []byte{0xd, 0x8, 0x20, 0x4e, 0x0, 0x0, 0x30, 0x75, 0x0, 0x0}) {
		t.Fatalf("Invalid encodeFixed32 (repeated) result: %#v", buf)
	}
}

func TestEncodeFixed64(t *testing.T) {
	wt := FieldTypeToWireType(descriptor.FieldDescriptorProto_TYPE_FIXED64)
	n := 1

	var buf []byte
	var err error

	buf, err = encodeFixed64(float64(100000), n, wt, false)

	if err != nil {
		t.Fatalf("Error with encodeFixed64: %#v", err)
	}

	if !bytes.Equal(buf, []byte{0x9, 0xa0, 0x86, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0}) {
		t.Fatalf("Invalid encodeFixed64 result: %#v", buf)
	}
}

func TestEncodeFixed64Repeated(t *testing.T) {
	wt := FieldTypeToWireType(descriptor.FieldDescriptorProto_TYPE_FIXED64)
	n := 1

	var buf []byte
	var err error

	buf, err = encodeFixed64([]interface{}{float64(200000), float64(300000)}, n, wt, true)

	if err != nil {
		t.Fatalf("Error with encodeFixed64 (repeated): %#v", err)
	}

	if !bytes.Equal(buf, []byte{0x9, 0x10, 0x40, 0xd, 0x3, 0x0, 0x0, 0x0, 0x0, 0x0, 0xe0, 0x93, 0x4, 0x0, 0x0, 0x0, 0x0, 0x0}) {
		t.Fatalf("Invalid encodeFixed64 (repeated) result: %#v", buf)
	}
}

func TestEncodeInt(t *testing.T) {
	wt := FieldTypeToWireType(descriptor.FieldDescriptorProto_TYPE_INT64)
	n := 1

	var buf []byte
	var err error

	buf, err = encodeInt(float64(500), n, wt, false)

	if err != nil {
		t.Fatalf("Error with encodeInt: %#v", err)
	}

	if !bytes.Equal(buf, []byte{0x8, 0xf4, 0x3}) {
		t.Fatalf("Invalid encodeInt result: %#v", buf)
	}
}

func TestEncodeIntRepeated(t *testing.T) {
	wt := FieldTypeToWireType(descriptor.FieldDescriptorProto_TYPE_INT64)
	n := 1

	var buf []byte
	var err error

	buf, err = encodeInt([]interface{}{float64(1000), float64(1500)}, n, wt, true)

	if err != nil {
		t.Fatalf("Error with encodeInt (repeated): %#v", err)
	}

	if !bytes.Equal(buf, []byte{0x8, 0x4, 0xe8, 0x7, 0xdc, 0xb}) {
		t.Fatalf("Invalid encodeInt (repeated) result: %#v", buf)
	}
}
