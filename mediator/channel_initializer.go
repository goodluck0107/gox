package mediator

import (
	"gitee.com/andyxt/gox/message"

	"gitee.com/andyxt/gona/boot/channel"
	"gitee.com/andyxt/gox/handler/code"
	"gitee.com/andyxt/gox/handler/schedule"
)

func NewChannelInitializer(mInboundChannelCommandMaker schedule.IChannelInboundCommandMaker, mMessageFactory message.IMessageFactory) channel.ChannelInitializer {
	return newChannelInitializer(mInboundChannelCommandMaker, mMessageFactory)
}

func newChannelInitializer(mInboundChannelCommandMaker schedule.IChannelInboundCommandMaker, mMessageFactory message.IMessageFactory) channel.ChannelInitializer {
	instance := new(channelInitializer)
	instance.mInboundChannelCommandMaker = mInboundChannelCommandMaker
	instance.mMessageFactory = mMessageFactory
	return instance
}

type channelInitializer struct {
	mInboundChannelCommandMaker schedule.IChannelInboundCommandMaker
	mMessageFactory             message.IMessageFactory
}

func (initializer *channelInitializer) InitChannel(pipeline channel.ChannelPipeline) {
	if pipeline == nil {
		return
	}
	// UpHandleOnRoutineSync--CTS SecurityDecoder -->  MessageDecoder-->  ExecutionHandleOnRoutineSync
	pipeline.AddLast("MessageDecoder", code.NewMessageDecoderHandleOnRoutineSync(initializer.mMessageFactory))
	pipeline.AddLast("InBoundExecutionHandler", schedule.NewInBoundExecutionHandler(initializer.mInboundChannelCommandMaker))
	// DownHandleOnRoutineSync--STS or STC  MessageEncoder -->  SecurityEncoder
	pipeline.AddLast("MessageEncoder", code.NewMessageEncoderHandleOnRoutineSync())
}
