package channelCommands

import (
	"gitee.com/andyxt/gox/extends"
	"gitee.com/andyxt/gox/handler/schedule"
	"gitee.com/andyxt/gox/mediator/client/clientkey"
	"gitee.com/andyxt/gox/service"

	"gitee.com/andyxt/gona/logger"
	"gitee.com/andyxt/gox/executor"
)

type ClientCommandInActive struct {
	mEventMaker schedule.IRoutineInboundEventMaker
	ChannelCtx  service.IChannelContext
}

func NewClientCommandInActive(mEventMaker schedule.IRoutineInboundEventMaker, ChannelCtx service.IChannelContext) (this *ClientCommandInActive) {
	this = new(ClientCommandInActive)
	this.mEventMaker = mEventMaker
	this.ChannelCtx = ChannelCtx
	return
}

func (event *ClientCommandInActive) Exec() {
	logger.Debug("ClientCommandInActive Exec", extends.ChannelContextToString(event.ChannelCtx))
	uID := event.ChannelCtx.ContextAttr().GetInt64(clientkey.KeyFireUser)
	executor.FireEvent(event.mEventMaker.MakeInActiveEvent(uID, event.ChannelCtx))
}
