package handler

import (
	"github.com/goodluck0107/gona/boot/channel"
	"github.com/goodluck0107/gox/code/protocol"
	"github.com/goodluck0107/gox/internal/logger"
)

// protocol.Protocol ---> []byte
type ProtocolEncoder struct {
}

func NewProtocolEncoder() (this *ProtocolEncoder) {
	this = new(ProtocolEncoder)
	return
}

func (encoder *ProtocolEncoder) ExceptionCaught(ctx channel.ChannelContext, err error) {
	//	logger.Debug("MessageEncoder ExceptionCaught")
}

func (encoder *ProtocolEncoder) Write(ctx channel.ChannelContext, e interface{}) interface{} {
	//	logger.Debug("MessageEncoder Write")
	proto := e.(protocol.Protocol)
	buf, err := proto.Encode()
	if err != nil {
		logger.Error("ProtocolEncoder Write err=", err)
		return nil
	}
	return buf
}

func (encoder *ProtocolEncoder) Close(ctx channel.ChannelContext) {
	//	logger.Debug("MessageEncoder Close")
}
