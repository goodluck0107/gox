package handler

import (
	"github.com/goodluck0107/gona/boot/channel"
	"github.com/goodluck0107/gox/code/message"
	"github.com/goodluck0107/gox/code/protocol"
	"github.com/goodluck0107/gox/internal/logger"
)

func NewProtocolRawEncoder() (this *ProtocolRawEncoder) {
	this = new(ProtocolRawEncoder)
	return
}

// json ---> []byte
type ProtocolRawEncoder struct {
}

func (encoder *ProtocolRawEncoder) ExceptionCaught(ctx channel.ChannelContext, err error) {
	//	logger.Debug("MessageEncoder ExceptionCaught")
}

func (encoder *ProtocolRawEncoder) Write(ctx channel.ChannelContext, e interface{}) interface{} {
	//	logger.Debug("MessageEncoder Write")
	msg := e.(message.CustomMessage)
	proto := protocol.Raw(msg)
	buf, err := proto.Encode()
	if err != nil {
		logger.Error("ProtocolRawEncoder Write err=", err)
		return nil
	}
	return buf
}

func (encoder *ProtocolRawEncoder) Close(ctx channel.ChannelContext) {
	//	logger.Debug("MessageEncoder Close")
}
