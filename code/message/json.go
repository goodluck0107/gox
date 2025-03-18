package message

import "encoding/json"

func Json(v any) CustomMessage {
	return &jsonMessage{v: v}
}

type jsonMessage struct {
	v any
}

func (bean *jsonMessage) Decode(e []byte) error {
	return json.Unmarshal(e, bean.v)
}

func (bean *jsonMessage) Encode() ([]byte, error) {
	return json.Marshal(bean.v)
}
