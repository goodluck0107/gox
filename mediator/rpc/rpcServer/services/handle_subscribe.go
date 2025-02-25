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

// SubscribeRequest 发起订阅
func (*RpcService) RouteForSubscribeRequest() (string, uint32, uint32) {
	return "/SubscribeRequest", uint32(mid.SubscribeRequest), service.ProtoTypePB
}

func (*RpcService) SubscribeRequest(request service.IServiceRequest, msg *rpc.SubscribeRequest) error {
	logger.Info(fmt.Sprintf("SubscribeRequest Topic:%v", msg.Topic))
	ctx := request.ChannelContext()
	executor.FireEvent(newSubscribeEvent(ctx, msg))
	return nil
}

type subscribeEvent struct {
	ctx service.IChannelContext
	msg *rpc.SubscribeRequest
}

func newSubscribeEvent(ctx service.IChannelContext, msg *rpc.SubscribeRequest) (this *subscribeEvent) {
	this = new(subscribeEvent)
	this.ctx = ctx
	this.msg = msg
	return this
}

func (recvEvent *subscribeEvent) QueueId() int64 {
	return time.Now().UnixNano()
}

func (recvEvent *subscribeEvent) Wait() (interface{}, bool) {
	return nil, true
}

func (recvEvent *subscribeEvent) Exec() {
	center.AddSub(recvEvent.msg.Topic, recvEvent.ctx)
}
