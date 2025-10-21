package mediator

import (
	"github.com/goodluck0107/gona/boot/channel"
	"github.com/goodluck0107/gox/code/protocol"
	"github.com/goodluck0107/gox/handler"
)

func NewChannelInitializer(mInboundChannelCommandMaker handler.IChannelInboundCommandMaker, mMessageFactory protocol.ProtocolFactory) channel.ChannelInitializer {
	return newChannelInitializer(mInboundChannelCommandMaker, mMessageFactory)
}

func newChannelInitializer(mInboundChannelCommandMaker handler.IChannelInboundCommandMaker, mMessageFactory protocol.ProtocolFactory) channel.ChannelInitializer {
	instance := new(channelInitializer)
	instance.mInboundChannelCommandMaker = mInboundChannelCommandMaker
	instance.mMessageFactory = mMessageFactory
	return instance
}

type channelInitializer struct {
	mInboundChannelCommandMaker handler.IChannelInboundCommandMaker
	mMessageFactory             protocol.ProtocolFactory
}

func (initializer *channelInitializer) InitChannel(pipeline channel.ChannelPipeline) {
	if pipeline == nil {
		return
	}
	// UpHandleOnRoutineSync--CTS SecurityDecoder -->  MessageDecoder-->  ExecutionHandleOnRoutineSync
	pipeline.AddLast("MessageDecoder", handler.NewProtocolDecoder(initializer.mMessageFactory))
	pipeline.AddLast("InBoundExecutionHandler", handler.NewInBoundExecutionHandler(initializer.mInboundChannelCommandMaker))
	// DownHandleOnRoutineSync--STS or STC  MessageEncoder -->  SecurityEncoder
	pipeline.AddLast("MessageEncoder", handler.NewProtocolEncoder())
}
