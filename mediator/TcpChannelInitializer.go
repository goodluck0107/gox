package mediator

import (
	"gitee.com/andyxt/gox/message"

	"gitee.com/andyxt/gona/boot/channel"
	"gitee.com/andyxt/gox/handler/code"
	"gitee.com/andyxt/gox/handler/schedule"
)

func NewTcpChannelInitializer(mInboundChannelCommandMaker schedule.IChannelInboundCommandMaker, mMessageFactory message.IMessageFactory) channel.ChannelInitializer {
	return newTcpChannelInitializer(mInboundChannelCommandMaker, mMessageFactory)
}

func newTcpChannelInitializer(mInboundChannelCommandMaker schedule.IChannelInboundCommandMaker, mMessageFactory message.IMessageFactory) channel.ChannelInitializer {
	instance := new(tcpChannelInitializer)
	instance.mInboundChannelCommandMaker = mInboundChannelCommandMaker
	instance.mMessageFactory = mMessageFactory
	return instance
}

type tcpChannelInitializer struct {
	mInboundChannelCommandMaker schedule.IChannelInboundCommandMaker
	mMessageFactory             message.IMessageFactory
}

func (initializer *tcpChannelInitializer) InitChannel(pipeline channel.ChannelPipeline) {
	if pipeline == nil {
		return
	}
	// UpHandleOnRoutineSync--CTS SecurityDecoder -->  MessageDecoder-->  ExecutionHandleOnRoutineSync
	pipeline.AddLast("MessageDecoder", code.NewMessageDecoderHandleOnRoutineSync(initializer.mMessageFactory))
	pipeline.AddLast("InBoundExecutionHandleOnRoutineSync", schedule.NewInBoundExecutionHandler(initializer.mInboundChannelCommandMaker))
	// DownHandleOnRoutineSync--STS or STC  MessageEncoder -->  SecurityEncoder
	pipeline.AddLast("MessageEncoder", code.NewMessageEncoderHandleOnRoutineSync())
}
