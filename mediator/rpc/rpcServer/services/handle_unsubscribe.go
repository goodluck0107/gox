package services

import (
	"fmt"

	"github.com/goodluck0107/gox/code/message"
	"github.com/goodluck0107/gox/executor"
	"github.com/goodluck0107/gox/internal/logger"
	"github.com/goodluck0107/gox/mediator/rpc/mid"
	"github.com/goodluck0107/gox/mediator/rpc/pb/rpc"
	"github.com/goodluck0107/gox/mediator/rpc/rpcServer/center"
	"github.com/goodluck0107/gox/service"
)

const (
	unsubscribePath = "/UnsubscribeRequest"
)

// UnsubscribeRequest 取消订阅
func (*RpcService) RouteForUnsubscribeRequest() (string, uint32, message.MessageType) {
	return unsubscribePath, uint32(mid.UnsubscribeRequest), message.MessageTypePB
}

func (*RpcService) UnsubscribeRequest(request service.IServiceRequest, msg *rpc.UnsubscribeRequest) error {
	logger.Info(fmt.Sprintf("UnsubscribeRequest Topic:%v", msg.Topic))
	executor.FireEvent(newUnsubscribeEvent(request.ChannelContext(), msg))
	return nil
}

type unsubscribeEvent struct {
	ctx service.IChannelContext
	msg *rpc.UnsubscribeRequest
}

func newUnsubscribeEvent(ctx service.IChannelContext, msg *rpc.UnsubscribeRequest) (this *unsubscribeEvent) {
	return &unsubscribeEvent{ctx: ctx, msg: msg}
}

func (recvEvent *unsubscribeEvent) QueueId() int64 {
	return stringToInt64(recvEvent.ctx.ID())
}

func (recvEvent *unsubscribeEvent) Wait() (interface{}, bool) {
	return nil, true
}

func (recvEvent *unsubscribeEvent) Exec() {
	center.DelSub(recvEvent.msg.Topic, recvEvent.ctx)
}
