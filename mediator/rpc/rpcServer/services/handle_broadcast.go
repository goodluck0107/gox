package services

import (
	"fmt"

	"gitee.com/andyxt/gona/logger"
	"gitee.com/andyxt/gox/executor"
	"gitee.com/andyxt/gox/extends"
	"gitee.com/andyxt/gox/mediator/rpc/mid"
	"gitee.com/andyxt/gox/mediator/rpc/pb/rpc"
	"gitee.com/andyxt/gox/messageImpl"
	"gitee.com/andyxt/gox/service"
	"gitee.com/andyxt/gox/session"
)

// RouteForRPCBroadcastRequest RPC广播
func (*RpcService) RouteForRPCBroadcastRequest() (string, uint32, uint32) {
	return "/RPCBroadcastRequest", uint32(mid.RPCBroadcastRequest), service.ProtoTypePB
}

func (*RpcService) RPCBroadcastRequest(request service.IServiceRequest, msg *rpc.RPCBroadcastRequest) error {
	logger.Info(fmt.Sprintf("RPCBroadcastRequest PlayerID:%v FuncCode:%v", msg.PlayerID, msg.FuncCode))
	session.TraverseDo(func(i1 session.ISession, i2 interface{}) {
		executor.FireEvent(newRpcBroadcastEvent(i1.UID(), msg))
	}, nil)
	return nil
}

type rpcBroadcastEvent struct {
	nodeID int64
	msg    *rpc.RPCBroadcastRequest
}

func newRpcBroadcastEvent(nodeID int64, msg *rpc.RPCBroadcastRequest) (this *rpcBroadcastEvent) {
	this = new(rpcBroadcastEvent)
	this.nodeID = nodeID
	this.msg = msg
	return this
}

func (recvEvent *rpcBroadcastEvent) QueueId() int64 {
	return recvEvent.nodeID
}

func (recvEvent *rpcBroadcastEvent) Wait() (interface{}, bool) {
	return nil, true
}

func (recvEvent *rpcBroadcastEvent) Exec() {
	playerSession := session.GetSession(recvEvent.nodeID)
	if playerSession == nil {
		return
	}
	Ctx := extends.GetChlCtx(playerSession)
	if Ctx == nil {
		return
	}
	messageImpl.Push(Ctx, mid.RPCCallPush, &rpc.RPCCallPush{
		PlayerID: recvEvent.msg.PlayerID,
		FuncCode: recvEvent.msg.FuncCode,
		FuncData: recvEvent.msg.FuncData,
	})
}
