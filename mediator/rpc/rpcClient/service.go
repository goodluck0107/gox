package rpcClient

import (
	"fmt"

	"gitee.com/andyxt/gona/logger"
	"gitee.com/andyxt/gox/executor"
	"gitee.com/andyxt/gox/extends"
	"gitee.com/andyxt/gox/mediator/client/clientkey"
	"gitee.com/andyxt/gox/mediator/rpc/mid"
	"gitee.com/andyxt/gox/mediator/rpc/pb/rpc"
	"gitee.com/andyxt/gox/message"
	"gitee.com/andyxt/gox/service"
	"gitee.com/andyxt/gox/session"
)

type RpcClientService struct{}

func NewService() *RpcClientService {
	return &RpcClientService{}
}

// RouteForRPCLoginResponse 登录
func (*RpcClientService) RouteForRPCLoginResponse() (string, uint32) {
	return "/RPCLoginResponse", uint32(mid.RPCLoginResponse)
}

func (*RpcClientService) RPCLoginResponse(request service.IServiceRequest, msg *rpc.LoginResponse) error {
	logger.Info(fmt.Sprintf("RPCLoginResponse Code:%v", msg.Code))
	chlCtx := request.ChannelContext()
	uID := chlCtx.ContextAttr().GetInt64(clientkey.KeyFireUser)
	extends.PutInUserInfo(chlCtx, uID, 0)
	return nil
}

// RouteForRPCLogoutResponse 登出
func (*RpcClientService) RouteForRPCLogoutResponse() (string, uint32) {
	return "/RPCLogoutResponse", uint32(mid.RPCLogoutResponse)
}

func (*RpcClientService) RPCLogoutResponse(request service.IServiceRequest, msg *rpc.LogoutResponse) error {
	logger.Info(fmt.Sprintf("RPCLogoutResponse Code:%v Message:%v", msg.Code, msg.Message))
	chlCtx := request.ChannelContext()
	extends.Logout(chlCtx)
	return nil
}

// RouteForRPCHeartbeatResponse 心跳
func (*RpcClientService) RouteForRPCHeartbeatResponse() (string, uint32) {
	return "/RPCHeartbeatResponse", uint32(mid.RPCHeartbeatResponse)
}

func (*RpcClientService) RPCHeartbeatResponse(request service.IServiceRequest, msg *rpc.HeartbeatResponse) error {
	logger.Info("RPCHeartbeatResponse")
	return nil
}

// RouteForRPCLoginConflictPush 登录冲突
func (*RpcClientService) RouteForRPCLoginConflictPush() (string, uint32) {
	return "/RPCLoginConflictPush", uint32(mid.RPCLoginConflictPush)
}

func (*RpcClientService) RPCLoginConflictPush(request service.IServiceRequest, msg *rpc.LoginConflictPush) error {
	logger.Info("RPCLoginConflictPush")
	chlCtx := request.ChannelContext()
	extends.Conflict(chlCtx)
	return nil
}

// RouteForRPCCallPush 处理服务器推送的RPC调用
func (*RpcClientService) RouteForRPCCallPush() (string, uint32) {
	return "/RPCCallPush", uint32(mid.RPCCallPush)
}

func (*RpcClientService) RPCCallPush(request service.IServiceRequest, msg *rpc.RPCCallPush) error {
	logger.Info(fmt.Sprintf("RPCCallPush PlayerID:%v FuncCode:%v", msg.PlayerID, msg.FuncCode))
	executor.FireEvent(newRpcCallPushEvent(msg))
	return nil
}

type rpcCallPushEvent struct {
	msg *rpc.RPCCallPush
}

func newRpcCallPushEvent(msg *rpc.RPCCallPush) (this *rpcCallPushEvent) {
	this = new(rpcCallPushEvent)
	this.msg = msg
	return this
}

func (recvEvent *rpcCallPushEvent) QueueId() int64 {
	return recvEvent.msg.PlayerID
}

func (recvEvent *rpcCallPushEvent) Wait() (interface{}, bool) {
	return nil, true
}

func (recvEvent *rpcCallPushEvent) Exec() {
	playerSession := session.GetSession(recvEvent.msg.PlayerID)
	if playerSession == nil {
		fmt.Printf("recvEvent to handle for playerID:%v, but player session is nil.\n", recvEvent.msg.PlayerID)
		return
	}
	Ctx := extends.GetChlCtx(playerSession)
	if Ctx == nil {
		fmt.Printf("recvEvent to handle for playerID:%v, but player ctx is nil.\n", recvEvent.msg.PlayerID)
		return
	}
	funcMsg := message.NewMessageDirect(1, 0, 1, 1, uint16(recvEvent.msg.FuncCode), recvEvent.msg.FuncData)
	callService(Ctx, funcMsg)
}
