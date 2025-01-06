package services

import (
	"fmt"

	"gitee.com/andyxt/gona/logger"
	"gitee.com/andyxt/gona/utils"
	"gitee.com/andyxt/gox/extends"
	"gitee.com/andyxt/gox/mediator/rpc/mid"
	"gitee.com/andyxt/gox/mediator/rpc/pb/rpc"
	"gitee.com/andyxt/gox/mediator/server"
	"gitee.com/andyxt/gox/service"
	"gitee.com/andyxt/gox/session"
)

// RouteForRPCLoginRequest 登录
func (*RpcService) RouteForRPCLoginRequest() (string, uint32, uint32) {
	return "/RPCLoginRequest", uint32(mid.RPCLoginRequest), service.ProtoTypePB
}

func (*RpcService) RPCLoginRequest(request service.IServiceRequest, msg *rpc.LoginRequest) error {
	logger.Info(fmt.Sprintf("RPCLoginRequest NodeID:%v", msg.NodeID))
	channelCtx := request.ChannelContext()
	uID := msg.NodeID
	playerSession := session.GetSession(uID)
	if playerSession != nil {
		if checkConflict(request) { // 登录冲突(异地登录或断线重连)
			oldChlCtx := extends.GetChlCtx(playerSession)
			extends.Conflict(oldChlCtx)
			server.ResponseClose(oldChlCtx, 0, mid.RPCLoginConflictPush, &rpc.LoginConflictPush{}, "Conflict")
			extends.ChangeChlCtx(playerSession, channelCtx)
		}
		return server.Response(channelCtx, extends.MsgID(request), mid.RPCLoginResponse, &rpc.LoginResponse{})
	}
	playerSession = session.NewSession(utils.UUID(), uID) // 构建玩家Session
	extends.ChangeChlCtx(playerSession, channelCtx)
	session.AddSession(playerSession) // 将玩家Session放入Session池中
	return server.Response(channelCtx, extends.MsgID(request), mid.RPCLoginResponse, &rpc.LoginResponse{})
}

// checkConflict 检查异地登录冲突-断线重连或异地登陆
func checkConflict(request service.IServiceRequest) bool {
	newChlCtx := request.ChannelContext() // 本次请求的网络连接上下文
	uID := extends.UID(newChlCtx)
	playerSession := session.GetSession(uID)
	oldChlCtx := extends.GetChlCtx(playerSession) // session当前正在使用的网络连接上下文
	logger.Error(fmt.Sprintf("service account handler_login checkConflict 玩家%v,旧连接%v,新连接%v", uID, oldChlCtx.ID(), newChlCtx.ID()))
	return !extends.ChannelContextEquals(oldChlCtx, newChlCtx)
}
