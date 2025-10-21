package httpor

import (
	"github.com/goodluck0107/gox/handler"

	"github.com/goodluck0107/gona/boot/channel"
)

func NewChannelInitializerRawJson() channel.ChannelInitializer {
	return new(ChannelInitializerRawJson)
}

type ChannelInitializerRawJson struct {
}

func (initializer *ChannelInitializerRawJson) InitChannel(pipeline channel.ChannelPipeline) {
	if pipeline == nil {
		return
	}
	// UpHandleOnRoutineSync--CTS SecurityDecoder -->  MessageDecoder-->  ExecutionHandleOnRoutineSync
	// pipeline.AddLast("MessageDecoder", code.NewMessageDecoderHandleOnRoutineSync(initializer.mMessageFactory)) // 消息解码处理器
	pipeline.AddLast("ExecutionHandler", NewExecutionHandler()) // 消息逻辑处理
	// DownHandleOnRoutineSync--STS or STC  MessageEncoder -->  SecurityEncoder
	pipeline.AddLast("MessageEncoder", handler.NewProtocolRawJsonEncoder()) // 消息编码处理器
}
