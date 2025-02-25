package services

import (
	"fmt"

	"gitee.com/andyxt/gona/logger"
	"gitee.com/andyxt/gox/mediator/server"
	"gitee.com/andyxt/gox/service"
	"gitee.com/andyxt/gox/session"
)

type RpcService struct{}

func NewService() *RpcService {
	listenEvent()
	return &RpcService{}
}

func listenEvent() {
	server.OnClose(onInactive)
}

// onInactive 连接中断
func onInactive(playerID int64, chlCtx service.IChannelContext) {
	logger.Info(fmt.Sprintf("onInactive playerID:%v 掉线", playerID))
	s := session.GetSession(playerID)
	if s == nil {
		return
	}
	logger.Info(fmt.Sprintf("onInactive  playerID:%v 掉线成功", playerID))
	session.RemoveSession(playerID)
}
