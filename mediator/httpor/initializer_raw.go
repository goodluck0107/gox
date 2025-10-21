package httpor

import (
	"github.com/goodluck0107/gox/handler"

	"github.com/goodluck0107/gona/boot/channel"
)

func NewChannelInitializerRaw() channel.ChannelInitializer {
	return new(ChannelInitializerRaw)
}

type ChannelInitializerRaw struct {
}

func (initializer *ChannelInitializerRaw) InitChannel(pipeline channel.ChannelPipeline) {
	if pipeline == nil {
		return
	}
	// UpHandleOnRoutineSync--CTS SecurityDecoder -->  MessageDecoder-->  ExecutionHandleOnRoutineSync
	// pipeline.AddLast("MessageDecoder", code.NewMessageDecoderHandleOnRoutineSync(initializer.mMessageFactory)) // 消息解码处理器
	pipeline.AddLast("ExecutionHandler", NewExecutionHandler()) // 消息逻辑处理
	// DownHandleOnRoutineSync--STS or STC  MessageEncoder -->  SecurityEncoder
	pipeline.AddLast("MessageEncoder", handler.NewProtocolRawEncoder()) // 消息编码处理器
}
