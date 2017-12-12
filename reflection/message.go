package reflection

type Message struct {
	Payload []byte
}

func (m *Message) Reset() {}

func (m *Message) String() string {
	return ""
}

func (m *Message) ProtoMessage() {}

func (m *Message) Marshal() ([]byte, error) {
	return m.Payload, nil
}

func (m *Message) Unmarshal(buf []byte) error {
	m.Payload = buf
	return nil
}
