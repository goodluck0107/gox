package inboundCommands

import (
	"gitee.com/andyxt/gona/logger"
	"gitee.com/andyxt/gox/handler/schedule"
	"gitee.com/andyxt/gox/mediator/client/clientkey"
	"gitee.com/andyxt/gox/service"
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
	logger.Debug("ClientChannelUpActiveEvent Exec", "新连接建立成功")
	uID := activeEvent.ChlCtx.ContextAttr().GetInt64(clientkey.KeyFireUser)
	activeEvent.callback.ConnectSuccess(uID, activeEvent.ChlCtx)
}
