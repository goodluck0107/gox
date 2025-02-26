package services

import (
	"fmt"

	"gitee.com/andyxt/gona/logger"
	"gitee.com/andyxt/gox/executor"
	"gitee.com/andyxt/gox/mediator/rpc/mid"
	"gitee.com/andyxt/gox/mediator/rpc/pb/rpc"
	"gitee.com/andyxt/gox/mediator/rpc/rpcServer/center"
	"gitee.com/andyxt/gox/service"
)

const (
	unsubscribePath = "/UnsubscribeRequest"
)

// UnsubscribeRequest 取消订阅
func (*RpcService) RouteForUnsubscribeRequest() (string, uint32, uint32) {
	return unsubscribePath, uint32(mid.UnsubscribeRequest), service.ProtoTypePB
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
