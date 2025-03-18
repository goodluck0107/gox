package message

type CustomMessage interface {
	Decode(b []byte) error

	Encode() ([]byte, error)
}
