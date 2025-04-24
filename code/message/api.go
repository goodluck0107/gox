package message

import (
	"encoding/json"
	"errors"
	"fmt"

	"google.golang.org/protobuf/proto"
)

func Unmarshal(msgType MessageType, b []byte, v any) error {
	if msgType == MessageTypePB { // ProtoBuffer解码(二进制->消息)
		pm, ok := v.(proto.Message)
		if !ok {
			return errors.New("param:%v is not proto.Message while protoType is ProtoTypePB")
		}
		err := proto.Unmarshal(b, pm)
		if err != nil {
			return fmt.Errorf("param:%v can not unmarshal to proto.Message while protoType is ProtoTypePB:%v", b, err)
		}
		return nil
	} else if msgType == MessageTypeJson { // json解码(二进制->消息)
		err := json.Unmarshal(b, v)
		if err != nil {
			return fmt.Errorf("param:%v can not unmarshal to json while protoType is ProtoTypeJson:%v", b, err)
		}
		return nil
	} else if msgType == MessageTypeCustom { // 自定义解码(二进制->消息)
		bm, ok := v.(CustomMessage)
		if !ok {
			return fmt.Errorf("param:%v is not message.CustomMessage while protoType is ProtoTypeBN:%v", b, "not CustomMessage")
		}
		err := bm.Decode(b)
		if err != nil {
			return fmt.Errorf("param:%v can not unmarshal to message.CustomMessage while protoType is ProtoTypeBN:%v", b, err)
		}
		return nil
	}
	return errors.New("param can not unmarshal to message while protoType is invalid")
}
