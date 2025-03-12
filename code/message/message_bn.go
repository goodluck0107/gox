package message

type IMessage interface {
	Decode(b []byte) error

	Encode() ([]byte, error)
}

func Unmarshal(b []byte, m IMessage) error {
	return m.Decode(b)
}

func Marshal(m IMessage) ([]byte, error) {

	return nil, nil
}
