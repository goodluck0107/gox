package services

import (
	"fmt"

	"gitee.com/andyxt/gona/logger"
	"gitee.com/andyxt/gox/executor"
	"gitee.com/andyxt/gox/extends"
	"gitee.com/andyxt/gox/mediator/rpc/mid"
	"gitee.com/andyxt/gox/mediator/rpc/pb/rpc"
	"gitee.com/andyxt/gox/mediator/server"
	"gitee.com/andyxt/gox/service"
	"gitee.com/andyxt/gox/session"
)

// RouteForRPCCallRequest RPC调用
func (*RpcService) RouteForRPCCallRequest() (string, uint32) {
	return "/RPCCallRequest", uint32(mid.RPCCallRequest)
}

func (*RpcService) RPCCallRequest(request service.IServiceRequest, msg *rpc.RPCCallRequest) error {
	logger.Info(fmt.Sprintf("RPCCallRequest NodeID:%v PlayerID:%v FuncCode:%v", msg.NodeID, msg.PlayerID, msg.FuncCode))
	executor.FireEvent(newRpcCallEvent(msg.NodeID, msg))
	return nil
}

type rpcCallEvent struct {
	nodeID int64
	msg    *rpc.RPCCallRequest
}

func newRpcCallEvent(nodeID int64, msg *rpc.RPCCallRequest) (this *rpcCallEvent) {
	this = new(rpcCallEvent)
	this.nodeID = nodeID
	this.msg = msg
	return this
}

func (recvEvent *rpcCallEvent) QueueId() int64 {
	return recvEvent.nodeID
}

func (recvEvent *rpcCallEvent) Wait() (interface{}, bool) {
	return nil, true
}

func (recvEvent *rpcCallEvent) Exec() {
	playerSession := session.GetSession(recvEvent.nodeID)
	if playerSession == nil {
		return
	}
	Ctx := extends.GetChlCtx(playerSession)
	if Ctx == nil {
		return
	}
	server.Push(Ctx, mid.RPCCallPush, &rpc.RPCCallPush{
		PlayerID: recvEvent.msg.PlayerID,
		FuncCode: recvEvent.msg.FuncCode,
		FuncData: recvEvent.msg.FuncData,
	})
}
