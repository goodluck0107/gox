package inboundCommands

import (
	"gitee.com/andyxt/gox/code/protocol"
	"gitee.com/andyxt/gox/handler"
	"gitee.com/andyxt/gox/service"
)

type ClientChannelUpMsgRecvEvent struct {
	routineId   int64
	Data        protocol.Protocol
	Ctx         service.IChannelContext
	mEventMaker handler.IRoutineInboundEventMaker
	callback    ICallBack
}

func NewClientChannelUpMsgRecvEvent(routineId int64, Data protocol.Protocol,
	Ctx service.IChannelContext, mEventMaker handler.IRoutineInboundEventMaker, callback ICallBack) (this *ClientChannelUpMsgRecvEvent) {
	this = new(ClientChannelUpMsgRecvEvent)
	this.routineId = routineId
	this.Data = Data
	this.Ctx = Ctx
	this.mEventMaker = mEventMaker
	this.callback = callback
	return this
}

func (msgRecvEvent *ClientChannelUpMsgRecvEvent) QueueId() int64 {
	return msgRecvEvent.routineId
}
func (msgRecvEvent *ClientChannelUpMsgRecvEvent) Wait() (result interface{}, ok bool) {
	return nil, true
}
func (msgRecvEvent *ClientChannelUpMsgRecvEvent) Exec() {
	// logger.Info("ClientChannelUpMsgRecvEvent Exec")
	msgRecvEvent.callback.MessageReceived(msgRecvEvent.Ctx, msgRecvEvent.Data)
}
