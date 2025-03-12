package channelCommands

import (
	"gitee.com/andyxt/gox/extends"
	"gitee.com/andyxt/gox/handler"
	"gitee.com/andyxt/gox/service"

	"gitee.com/andyxt/gona/logger"
	"gitee.com/andyxt/gox/executor"
)

type ServerCommandInActive struct {
	mEventMaker handler.IRoutineInboundEventMaker
	ChannelCtx  service.IChannelContext
}

func NewServerCommandInActive(
	mEventMaker handler.IRoutineInboundEventMaker, ChannelCtx service.IChannelContext) (this *ServerCommandInActive) {
	this = new(ServerCommandInActive)
	this.mEventMaker = mEventMaker
	this.ChannelCtx = ChannelCtx
	return
}

func (event *ServerCommandInActive) Exec() {
	logger.Debug("ServerCommandInActive Exec", extends.ChannelContextToString(event.ChannelCtx))
	if !extends.HasUserInfo(event.ChannelCtx) {
		logger.Debug("ServerCommandInActive Exec", extends.ChannelContextToString(event.ChannelCtx), "ChannelCtx is not IsInPool")
		extends.Close(event.ChannelCtx)
		return
	}
	executor.FireEvent(event.mEventMaker.MakeInActiveEvent(extends.UID(event.ChannelCtx), event.ChannelCtx))
}
