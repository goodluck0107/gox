package outhttpor

import (
	"github.com/goodluck0107/gox/code/protocol"
	"github.com/goodluck0107/gox/handler"

	"github.com/goodluck0107/gona/boot/channel"
)

func NewChannelInitializerRaw(mMessageFactory protocol.ProtocolFactory) channel.ChannelInitializer {
	return &ChannelInitializerRaw{mMessageFactory: mMessageFactory}
}

type ChannelInitializerRaw struct {
	mMessageFactory protocol.ProtocolFactory
}

func (initializer *ChannelInitializerRaw) InitChannel(pipeline channel.ChannelPipeline) {
	if pipeline == nil {
		return
	}
	// UpHandleOnRoutineSync--CTS SecurityDecoder -->  MessageDecoder-->  ExecutionHandleOnRoutineSync
	// pipeline.AddLast("MessageDecoder", code.NewMessageDecoderHandleOnRoutineSync(initializer.mMessageFactory)) // 消息解码处理器
	pipeline.AddLast("MessageDecoder", handler.NewProtocolDecoder(initializer.mMessageFactory))

	pipeline.AddLast("ExecutionHandler", NewExecutionHandler()) // 消息逻辑处理
	// DownHandleOnRoutineSync--STS or STC  MessageEncoder -->  SecurityEncoder
	pipeline.AddLast("MessageEncoder", handler.NewProtocolRawEncoder()) // 消息编码处理器
}
