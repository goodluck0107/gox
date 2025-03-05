package httpor

import (
	"fmt"

	"gitee.com/andyxt/gona/logger"
	"gitee.com/andyxt/gox/handler/code"
	"gitee.com/andyxt/gox/service"

	"gitee.com/andyxt/gona/boot"
	"gitee.com/andyxt/gona/boot/channel"
)

type ChannelInitializer struct {
}

func NewChannelInitializer() (instance *ChannelInitializer) {
	instance = new(ChannelInitializer)
	return
}

func (initializer *ChannelInitializer) InitChannel(pipeline channel.ChannelPipeline) {
	if pipeline == nil {
		return
	}
	// UpHandleOnRoutineSync--CTS SecurityDecoder -->  MessageDecoder-->  ExecutionHandleOnRoutineSync
	// pipeline.AddLast("MessageDecoder", code.NewMessageDecoderHandleOnRoutineSync(initializer.mMessageFactory)) // 消息解码处理器
	pipeline.AddLast("ExecutionHandler", NewExecutionHandler()) // 消息逻辑处理
	// DownHandleOnRoutineSync--STS or STC  MessageEncoder -->  SecurityEncoder
	pipeline.AddLast("MessageEncoder", code.NewMessageEncoder()) // 消息编码处理器
}

type ExecutionHandler struct {
}

func NewExecutionHandler() (this *ExecutionHandler) {
	this = new(ExecutionHandler)
	return
}

func (handler *ExecutionHandler) ExceptionCaught(ctx channel.ChannelContext, err error) {
	logger.Error(ctx.ID(), "ExecutionHandler 链接处理异常:", err)

}

func (handler *ExecutionHandler) ChannelActive(ctx channel.ChannelContext) (goonNext bool) {
	logger.Info(ctx.ID(), "ExecutionHandler 链接建立", "URLPath", ctx.ContextAttr().GetString(boot.KeyURLPath))
	return
}

func (handler *ExecutionHandler) ChannelInactive(ctx channel.ChannelContext) (goonNext bool) {
	logger.Info(ctx.ID(), "ExecutionHandler 链接中断")
	return
}

func (handler *ExecutionHandler) MessageReceived(ctx channel.ChannelContext, e interface{}) (ret interface{}, goonNext bool) {
	msg := e.([]byte)
	if msg == nil {
		msg = []byte{}
	}
	logger.Info(ctx.ID(), "ExecutionHandler 链接收到消息", "URLPath", ctx.ContextAttr().GetString(boot.KeyURLPath), "msg", string(msg))
	reqPath := ctx.ContextAttr().GetString(channel.KeyForReqPath)
	servicePath := fmt.Sprintf("%v", reqPath)
	request := service.NewSessionRequest(ctx, service.NewAttr(nil))
	serviceErr := service.DispatchByPath(servicePath, request, msg)
	if serviceErr != nil {
		logger.Error(fmt.Sprintf("%v callService:%v error:%v ", ctx.ID(), servicePath, serviceErr))
	}
	return
}
