package services

import (
	"fmt"
	"time"

	"gitee.com/andyxt/gona/logger"
	"gitee.com/andyxt/gox/executor"
	"gitee.com/andyxt/gox/mediator/rpc/center"
	"gitee.com/andyxt/gox/mediator/rpc/mid"
	"gitee.com/andyxt/gox/mediator/rpc/pb/rpc"
	"gitee.com/andyxt/gox/service"
)

// UnsubscribeRequest 取消订阅
func (*RpcService) RouteForUnsubscribeRequest() (string, uint32, uint32) {
	return "/UnsubscribeRequest", uint32(mid.UnsubscribeRequest), service.ProtoTypePB
}

func (*RpcService) UnsubscribeRequest(request service.IServiceRequest, msg *rpc.UnsubscribeRequest) error {
	logger.Info(fmt.Sprintf("UnsubscribeRequest Topic:%v", msg.Topic))
	ctx := request.ChannelContext()
	executor.FireEvent(newUnsubscribeEvent(ctx, msg))
	return nil
}

type unsubscribeEvent struct {
	ctx service.IChannelContext
	msg *rpc.UnsubscribeRequest
}

func newUnsubscribeEvent(ctx service.IChannelContext, msg *rpc.UnsubscribeRequest) (this *unsubscribeEvent) {
	this = new(unsubscribeEvent)
	this.ctx = ctx
	this.msg = msg
	return this
}

func (recvEvent *unsubscribeEvent) QueueId() int64 {
	return time.Now().UnixNano()
}

func (recvEvent *unsubscribeEvent) Wait() (interface{}, bool) {
	return nil, true
}

func (recvEvent *unsubscribeEvent) Exec() {
	center.DelSub(recvEvent.msg.Topic, recvEvent.ctx)
}
