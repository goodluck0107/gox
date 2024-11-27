package notify

import (
	"encoding/json"
	"fmt"

	"gitee.com/andyxt/gona/logger"
	"gitee.com/andyxt/gox/executor"
	"gitee.com/andyxt/gox/session"
)

// 通知单个玩家
func NotifyPlayer(playerID int64, notifyType string, v interface{}) {
	jsonV, jsonE := json.Marshal(v)
	if jsonE != nil {
		logger.Error(fmt.Sprintf("NotifyPlayer playerID:%v,notifyType:%v,v:%v,error:%v", playerID, notifyType, v, jsonE))
		return
	}
	notifyData := string(jsonV)
	notify(playerID, notifyType, notifyData)
}

// 通知全服在线玩家
func NotifyAllOnlinePlayer(notifyType string, v interface{}) {
	jsonV, jsonE := json.Marshal(v)
	if jsonE != nil {
		logger.Error(fmt.Sprintf("NotifyAllOnlinePlayer notifyType:%v,v:%v,error:%v", notifyType, v, jsonE))
		return
	}
	notifyData := string(jsonV)
	onlinePlayers := make([]int64, 0)
	session.TraverseDo(0, func(s session.ISession, param interface{}) {
		onlinePlayers = append(onlinePlayers, s.UID())
	}, nil)
	for _, playerID := range onlinePlayers {
		executor.FireEvent(newRoutineInboundCmdNotify(playerID, playerID, notifyType, notifyData))
	}
}

// // 通知数据
// type NotifyData struct {
// 	PlayerID int64                  // 玩家ID
// 	Context  service.IChannelContext // 上下文
// }

// // Inactive 通知玩家掉线
// func Inactive(playerID int64, chlCtx service.IChannelContext) {
// 	executor.FireEvent(NewRoutineInboundCmdInactive(playerID, chlCtx))
// }

// // 封禁玩家
// type ForbidReq struct {
// 	PlayerID int64 // 玩家ID
// }

// // ForbidPlayer 通知玩家被封禁
// func ForbidPlayer(playerID int64) {
// 	executor.FireEvent(NewRoutineInboundCmdForbid(playerID))
// }

// // 剔除离桌玩家
// type KickLeave struct {
// 	PlayerID      int64 // 玩家ID
// 	EnterGameTime int64 // 进入游戏的时间
// }

// // KickOfflinePlayer 通知掉线玩家被踢出
// func KickOfflinePlayer(playerID int64, chlCtx service.IChannelContext) {
// 	extends.SystemKick(chlCtx)
// 	executor.FireEvent(NewRoutineInboundCmdKickOffline(playerID, chlCtx, message.KickOfflineRequest, &cli.KickOfflineReq{}))
// }

// // UpdatePlayer 通知玩家更新信息
// func UpdatePlayer(playerID int64) {
// 	notify(playerID, notifyType, notifyData)
// 	executor.FireEvent(NewRoutineInboundCmdUpdate(playerID))
// }

// const (
// 	NotifyTypeKickLeave    = "KickLeave"    // 踢出离桌玩家(掉线或非法退桌)
// 	NotifyTypeForbidPlayer = "ForbidPlayer" // 通知玩家被封禁
// 	NotifyTypeUpdatePlayer = "UpdatePlayer" // 通知玩家更新信息
// )
