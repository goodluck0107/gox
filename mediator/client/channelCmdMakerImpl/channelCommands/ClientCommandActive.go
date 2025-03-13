package channelCommands

import (
	"gitee.com/andyxt/gox/extends"
	"gitee.com/andyxt/gox/handler"
	"gitee.com/andyxt/gox/mediator/client/clientkey"
	"gitee.com/andyxt/gox/service"

	"gitee.com/andyxt/gox/executor"
	"gitee.com/andyxt/gox/internal/logger"
)

type ClientCommandActive struct {
	mEventMaker handler.IRoutineInboundEventMaker
	ChannelCtx  service.IChannelContext
}

func NewClientCommandActive(mEventMaker handler.IRoutineInboundEventMaker, ChannelCtx service.IChannelContext) (this *ClientCommandActive) {
	this = new(ClientCommandActive)
	this.mEventMaker = mEventMaker
	this.ChannelCtx = ChannelCtx
	return
}

func (event *ClientCommandActive) Exec() {
	logger.Debug("ClientCommandActive Exec", extends.ChannelContextToString(event.ChannelCtx))
	uID := event.ChannelCtx.ContextAttr().GetInt64(clientkey.KeyFireUser)
	executor.FireEvent(event.mEventMaker.MakeActiveEvent(uID, event.ChannelCtx))
}
