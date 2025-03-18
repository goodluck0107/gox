package message

type CustomMessage interface {
	Decode(b []byte) error

	Encode() ([]byte, error)
}

func Unmarshal(b []byte, m CustomMessage) error {
	return m.Decode(b)
}

func Marshal(m CustomMessage) ([]byte, error) {

	return nil, nil
}
