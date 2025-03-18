package services

import (
	"fmt"

	"gitee.com/andyxt/gox/code/message"
	"gitee.com/andyxt/gox/executor"
	"gitee.com/andyxt/gox/internal/logger"
	"gitee.com/andyxt/gox/mediator/rpc/mid"
	"gitee.com/andyxt/gox/mediator/rpc/pb/rpc"
	"gitee.com/andyxt/gox/mediator/rpc/rpcServer/center"
	"gitee.com/andyxt/gox/service"
)

const (
	subscribePath = "/SubscribeRequest"
)

// SubscribeRequest 发起订阅
func (*RpcService) RouteForSubscribeRequest() (string, uint32, message.MessageType) {
	return subscribePath, uint32(mid.SubscribeRequest), message.MessageTypePB
}

func (*RpcService) SubscribeRequest(request service.IServiceRequest, msg *rpc.SubscribeRequest) error {
	logger.Info(fmt.Sprintf("SubscribeRequest Topic:%v", msg.Topic))
	executor.FireEvent(newSubscribeEvent(request.ChannelContext(), msg))
	return nil
}

type subscribeEvent struct {
	ctx service.IChannelContext
	msg *rpc.SubscribeRequest
}

func newSubscribeEvent(ctx service.IChannelContext, msg *rpc.SubscribeRequest) *subscribeEvent {
	return &subscribeEvent{ctx: ctx, msg: msg}
}

func (recvEvent *subscribeEvent) QueueId() int64 {
	return stringToInt64(recvEvent.ctx.ID())
}

func (recvEvent *subscribeEvent) Wait() (interface{}, bool) {
	return nil, true
}

func (recvEvent *subscribeEvent) Exec() {
	center.AddSub(recvEvent.msg.Topic, recvEvent.ctx)
}
