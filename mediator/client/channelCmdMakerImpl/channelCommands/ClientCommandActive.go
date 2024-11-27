package channelCommands

import (
	"gitee.com/andyxt/gox/extends"
	"gitee.com/andyxt/gox/handler/schedule"
	"gitee.com/andyxt/gox/service"

	"gitee.com/andyxt/gona/logger"
	"gitee.com/andyxt/gox/executor"
)

type ClientCommandActive struct {
	mEventMaker schedule.IRoutineInboundEventMaker
	ChannelCtx  service.IChannelContext
}

func NewClientCommandActive(mEventMaker schedule.IRoutineInboundEventMaker, ChannelCtx service.IChannelContext) (this *ClientCommandActive) {
	this = new(ClientCommandActive)
	this.mEventMaker = mEventMaker
	this.ChannelCtx = ChannelCtx
	return
}

func (event *ClientCommandActive) Exec() {
	logger.Debug("ClientCommandActive Exec", extends.ChannelContextToString(event.ChannelCtx))
	poolKey, ok := event.getFireUserID()
	if !ok {
		logger.Debug("连接激活失败：", " 失败原因：ctx.Get(ChannelId) not ok for type int64", extends.ChannelContextToString(event.ChannelCtx))
		extends.Close(event.ChannelCtx)
		return
	}
	extends.PutInUserInfo(event.ChannelCtx, poolKey, 0)
	executor.FireEvent(event.mEventMaker.MakeActiveEvent(poolKey, event.ChannelCtx))
}

func (event *ClientCommandActive) getFireUserID() (int64, bool) {
	channelId := extends.GetFireUser(event.ChannelCtx)
	ok := channelId != -1
	return channelId, ok
}
