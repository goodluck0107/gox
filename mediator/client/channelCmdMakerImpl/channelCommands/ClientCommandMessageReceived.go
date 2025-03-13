package channelCommands

import (
	"gitee.com/andyxt/gox/code/protocol"
	"gitee.com/andyxt/gox/extends"
	"gitee.com/andyxt/gox/handler"
	"gitee.com/andyxt/gox/mediator/client/clientkey"
	"gitee.com/andyxt/gox/service"

	"gitee.com/andyxt/gox/executor"
	"gitee.com/andyxt/gox/internal/logger"
)

type ClientCommandMessageReceived struct {
	mEventMaker handler.IRoutineInboundEventMaker
	ChannelCtx  service.IChannelContext
	e           interface{}
}

func NewClientCommandMessageReceived(mEventMaker handler.IRoutineInboundEventMaker, ChannelCtx service.IChannelContext, e interface{}) (this *ClientCommandMessageReceived) {
	this = new(ClientCommandMessageReceived)
	this.mEventMaker = mEventMaker
	this.ChannelCtx = ChannelCtx
	this.e = e
	return
}

func (event *ClientCommandMessageReceived) Exec() {
	logger.Debug("ClientCommandMessageReceived Exec", extends.ChannelContextToString(event.ChannelCtx))
	if event.ChannelCtx == nil || event.e == nil {
		return
	}
	buf, _ := event.e.(protocol.Protocol)
	uID := event.ChannelCtx.ContextAttr().GetInt64(clientkey.KeyFireUser)
	executor.FireEvent(event.mEventMaker.MakeMessageReceivedEvent(uID, buf, event.ChannelCtx))
}
