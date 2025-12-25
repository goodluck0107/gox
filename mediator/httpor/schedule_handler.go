package httpor

import (
	"fmt"
	"github.com/goodluck0107/gona/utils/cast"
	"github.com/goodluck0107/gox/eventBus"
	"github.com/goodluck0107/gox/extends"
	"github.com/goodluck0107/gox/internal/logger"
	"github.com/goodluck0107/gox/mediator/server/evts"
	"github.com/goodluck0107/gox/service"
	"net/http"

	"github.com/goodluck0107/gona/boot"
	"github.com/goodluck0107/gona/boot/channel"
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

	httpRequest := ctx.ContextAttr().Get(channel.KeyForRequest).(*http.Request)
	if httpRequest != nil {
		playerID := cast.ToInt64(httpRequest.Header.Get("X-Inner-PlayerID"))
		extends.PutInUserID(ctx, playerID)
		msgID := service.Code(servicePath)
		extends.SetMsgID(request, msgID)

		eventBus.Trigger(evts.EVT_ServiceBefore, request)
	}
	serviceErr := service.DispatchByPath(servicePath, request, msg)
	if serviceErr != nil {
		logger.Error(fmt.Sprintf("%v callService:%v error:%v ", ctx.ID(), servicePath, serviceErr))
	}
	if httpRequest != nil {
		eventBus.Trigger(evts.EVT_ServiceAfter, request, serviceErr)
	}
	return
}
