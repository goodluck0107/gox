package httpor

import (
	"fmt"

	"gitee.com/andyxt/gox/internal/logger"
	"gitee.com/andyxt/gox/service"

	"gitee.com/andyxt/gona/boot"
	"gitee.com/andyxt/gona/boot/channel"
)

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
	if e == nil {
		e = []byte{}
	}
	msg := e.([]byte)
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
