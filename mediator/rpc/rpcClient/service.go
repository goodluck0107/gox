package rpcClient

import (
	"fmt"
	"time"

	"github.com/goodluck0107/gox/code/message"
	"github.com/goodluck0107/gox/executor"
	"github.com/goodluck0107/gox/internal/logger"
	"github.com/goodluck0107/gox/mediator/rpc/mid"
	"github.com/goodluck0107/gox/mediator/rpc/pb/rpc"
	"github.com/goodluck0107/gox/messageImpl"
	"github.com/goodluck0107/gox/service"
)

type RpcClientService struct{}

func NewService() *RpcClientService {
	return &RpcClientService{}
}

// RouteForRPCHeartbeatResponse 心跳
func (*RpcClientService) RouteForRPCHeartbeatResponse() (string, uint32, message.MessageType) {
	return "/RPCHeartbeatResponse", uint32(mid.HeartbeatResponse), message.MessageTypePB
}

func (*RpcClientService) RPCHeartbeatResponse(request service.IServiceRequest, msg *rpc.HeartbeatResponse) error {
	logger.Info("RPCHeartbeatResponse")
	return nil
}

// RouteForMessagePush 处理服务器推送的RPC调用
func (*RpcClientService) RouteForMessagePush() (string, uint32, message.MessageType) {
	return "/MessagePush", uint32(mid.MessagePush), message.MessageTypePB
}

func (*RpcClientService) MessagePush(request service.IServiceRequest, msg *rpc.MessagePush) error {
	logger.Info(fmt.Sprintf("MessagePush Topic:%v MsgCode:%v", msg.Topic, msg.MsgCode))
	executor.FireEvent(newRpcCallPushEvent(request.ChannelContext(), msg))
	return nil
}

type rpcCallPushEvent struct {
	ctx service.IChannelContext
	msg *rpc.MessagePush
}

func newRpcCallPushEvent(ctx service.IChannelContext, msg *rpc.MessagePush) (this *rpcCallPushEvent) {
	this = new(rpcCallPushEvent)
	this.ctx = ctx
	this.msg = msg
	return this
}

func (recvEvent *rpcCallPushEvent) QueueId() int64 {
	return time.Now().Unix()
}

func (recvEvent *rpcCallPushEvent) Wait() (interface{}, bool) {
	return nil, true
}

func (recvEvent *rpcCallPushEvent) Exec() {
	funcMsg := messageImpl.NewMessageDirect(1, 0, 1, 1, uint16(recvEvent.msg.MsgCode), recvEvent.msg.MsgData)
	callService(recvEvent.ctx, funcMsg)
}
