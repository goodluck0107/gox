package notify

import (
	"sync"

	"gitee.com/andyxt/gona/logger"
	"gitee.com/andyxt/gox/executor"
	"gitee.com/andyxt/gox/session"
)

// isOnline,session,playerID,dataJson,
type NotifyFunc func(bool, session.ISession, int64, string)

type NotifyMap struct {
	sync.Map
}

var settleMap = &NotifyMap{}

// RegisterNotify 注册异步通知
func RegisterNotify(typ string, settleFunc NotifyFunc) {
	settleMap.Store(typ, settleFunc)
}

// RegisterNotify 取消注册通知
func DeregisterNotify(typ string) {
	settleMap.Delete(typ)
}

func notify(playerID int64, notifyType string, notifyData string) {
	executor.FireEvent(newRoutineInboundCmdNotify(playerID, playerID, notifyType, notifyData))
}

type routineInboundCmdNotify struct {
	routineId  int64
	playerID   int64
	notifyType string
	notifyData string
}

func newRoutineInboundCmdNotify(routineId int64, playerID int64, notifyType string, notifyData string) (this *routineInboundCmdNotify) {
	this = new(routineInboundCmdNotify)
	this.routineId = routineId
	this.playerID = playerID
	this.notifyType = notifyType
	this.notifyData = notifyData
	return this
}

func (recvEvent *routineInboundCmdNotify) QueueId() int64 {
	return recvEvent.routineId
}

func (recvEvent *routineInboundCmdNotify) Wait() (interface{}, bool) {
	return nil, true
}

func (recvEvent *routineInboundCmdNotify) Exec() {
	uID := recvEvent.routineId
	playerSession := session.GetSession(0, uID)
	settleFunc, ok := settleMap.Load(recvEvent.notifyType)
	if !ok {
		logger.Error("Notify type not registerd: ", recvEvent.notifyType)
		return
	}
	if playerSession == nil { // 玩家不在线：封禁玩家时需要特殊处理；踢出离线玩家时与离桌玩家时不需要存储，其他需要存储；
		settleFunc.(NotifyFunc)(false, playerSession, uID, recvEvent.notifyData)
	} else { // 玩家在线
		settleFunc.(NotifyFunc)(true, playerSession, uID, recvEvent.notifyData)
	}
}
