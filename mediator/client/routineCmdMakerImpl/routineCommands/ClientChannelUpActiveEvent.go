package routineCommands

import (
	"gitee.com/andyxt/gona/logger"
	"gitee.com/andyxt/gox/extends"
	"gitee.com/andyxt/gox/handler/schedule"
	"gitee.com/andyxt/gox/key"
	"gitee.com/andyxt/gox/service"
	"gitee.com/andyxt/gox/session"
)

type ClientChannelUpActiveEvent struct {
	routineId   int64
	mEventMaker schedule.IRoutineInboundEventMaker
	ChlCtx      service.IChannelContext
	callback    ICallBack
}

func NewClientChannelUpActiveEvent(routineId int64,
	mEventMaker schedule.IRoutineInboundEventMaker, ChlCtx service.IChannelContext, callback ICallBack) (this *ClientChannelUpActiveEvent) {
	this = new(ClientChannelUpActiveEvent)
	this.routineId = routineId
	this.mEventMaker = mEventMaker
	this.ChlCtx = ChlCtx
	this.callback = callback
	return
}

func (activeEvent *ClientChannelUpActiveEvent) QueueId() int64 {
	return activeEvent.routineId
}

func (activeEvent *ClientChannelUpActiveEvent) Wait() (result interface{}, ok bool) {
	return nil, true
}

func (activeEvent *ClientChannelUpActiveEvent) Exec() {
	uID := activeEvent.ChlCtx.ContextAttr().GetInt64(key.ChannelFireUser)
	iSession := session.GetSession(0, uID)
	if iSession == nil {
		logger.Debug("ClientChannelUpActiveEvent Exec", "连接已经被主动关闭，新连接直接关闭")
		extends.Close(activeEvent.ChlCtx)
		return
	}
	oldChlCtx := extends.GetChlCtx(iSession)
	if oldChlCtx == nil {
		extends.ChangeChlCtx(iSession, activeEvent.ChlCtx)
		activeEvent.callback.ConnectSuccess(uID, activeEvent.ChlCtx)
		return
	}
	if !extends.ChannelContextEquals(activeEvent.ChlCtx, oldChlCtx) {
		logger.Debug("ClientChannelUpActiveEvent Exec", "已经存在旧连接，直接关闭旧连接")
		extends.Conflict(oldChlCtx)
		extends.Close(oldChlCtx)
	}
	extends.ChangeChlCtx(iSession, activeEvent.ChlCtx)
	logger.Debug("ClientChannelUpActiveEvent Exec", "新连接建立成功")
	activeEvent.callback.ConnectSuccess(uID, activeEvent.ChlCtx)
}
