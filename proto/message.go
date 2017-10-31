package proto

import (
	"fmt"
	"github.com/lazada/grpc-ui/reflection"
)

// Dynamic proto Message

type FieldValue struct{
	Number int `json:"number"`
	Value string `json:"val"`
}

type Message struct {
	TypeInfo map[string]*reflection.TypeInfo
	TypeName string
	Data []FieldValue
	PB map[string]interface{}
}

// proto.Message interface

func (m *Message) Reset() {}

func (m *Message) String() string {
	return  "TODO"
}

func (m *Message) ProtoMessage() {}

// Marshaler and Unarshaler interfaces

func (m *Message) Marshal() ([]byte, error) {
	return Encode(m.TypeInfo, m.TypeName, m.Data)
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
