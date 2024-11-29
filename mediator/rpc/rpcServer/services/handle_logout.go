package services

import (
	"fmt"

	"gitee.com/andyxt/gona/logger"
	"gitee.com/andyxt/gox/extends"
	"gitee.com/andyxt/gox/mediator/rpc/mid"
	"gitee.com/andyxt/gox/mediator/rpc/pb/rpc"
	"gitee.com/andyxt/gox/service"
	"gitee.com/andyxt/gox/session"
)

// RouteForRPCLogoutRequest 登出
func (*RpcService) RouteForRPCLogoutRequest() (string, uint32) {
	return "/RPCLogoutRequest", uint32(mid.RPCLogoutRequest)
}

func (*RpcService) RPCLogoutRequest(request service.IServiceRequest, msg *rpc.LogoutRequest) error {
	playerID := extends.UID(request.ChannelContext())
	logger.Info(fmt.Sprintf("RPCLogoutRequest playerID:%v 登出", playerID))
	s := session.GetSession(playerID)
	if s == nil {
		return nil
	}
	logger.Info(fmt.Sprintf("RPCLogoutRequest playerID:%v 登出成功", playerID))
	session.RemoveSession(playerID)
	return nil
}

// onInactive 连接中断
func onInactive(data ...interface{}) {
	playerID := data[0].(int64)
	logger.Info(fmt.Sprintf("onInactive playerID:%v 掉线", playerID))
	s := session.GetSession(playerID)
	if s == nil {
		return
	}
	logger.Info(fmt.Sprintf("onInactive  playerID:%v 掉线成功", playerID))
	session.RemoveSession(playerID)
}
