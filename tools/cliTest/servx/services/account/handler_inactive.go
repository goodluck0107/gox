package account

import (
	"fmt"

	"gitee.com/andyxt/gox/service"
	"gitee.com/andyxt/gox/session"

	"gitee.com/andyxt/gox/internal/logger"
)

// OnInactive 连接中断
func onInactive(playerID int64, chlCtx service.IChannelContext) {
	fmt.Println("onInactive", "playerID:", playerID)
	logger.Info(fmt.Sprintf("InactiveRequest playerID:%v", playerID))
	s := session.GetSession(playerID)
	if s == nil {
		return
	}
	logger.Info(fmt.Sprintf("service account onInactive 玩家%v掉线", playerID))
	session.RemoveSession(playerID)
}
