package message

import "fmt"

type IMessage interface {
	Decode(b []byte) error

	Encode() ([]byte, error)
}

func Unmarshal(b []byte, m IMessage) error {
	fmt.Println("Unmarshal b:", b, "m:", m)
	return m.Decode(b)
}

func Marshal(m IMessage) ([]byte, error) {

	return nil, nil
}
