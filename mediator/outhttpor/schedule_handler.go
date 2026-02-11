package outhttpor

import (
	"fmt"
	"github.com/goodluck0107/gona/boot"
	"github.com/goodluck0107/gona/boot/channel"
	"github.com/goodluck0107/gox/code/protocol"
	"github.com/goodluck0107/gox/eventBus"
	"github.com/goodluck0107/gox/extends"
	"github.com/goodluck0107/gox/internal/logger"
	"github.com/goodluck0107/gox/mediator/server/evts"
	"github.com/goodluck0107/gox/service"
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
	protoMsg := e.(protocol.Protocol)
	logger.Info(ctx.ID(), "ExecutionHandler 链接收到消息", "URLPath", ctx.ContextAttr().GetString(boot.KeyURLPath), "msg", protoMsg)
	request := service.NewSessionRequest(ctx, service.NewAttr(nil))
	msgID := protoMsg.GetMsgID()
	extends.SetSeqID(request, protoMsg.GetSeqID())
	extends.SetMsgID(request, msgID)
	eventBus.Trigger(evts.EVT_ServiceBefore, request)
	serviceErr := service.DispatchByCode(uint32(msgID), request, protoMsg.GetMsgData())
	if serviceErr != nil {
		logger.Error(fmt.Sprintf("%v callService:%v error:%v ", ctx.ID(), protoMsg.GetMsgID(), serviceErr))
	}
	eventBus.Trigger(evts.EVT_ServiceAfter, request, serviceErr)
	return
}
