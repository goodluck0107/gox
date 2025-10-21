package channelCommands

import (
	"github.com/goodluck0107/gox/extends"
	"github.com/goodluck0107/gox/handler"
	"github.com/goodluck0107/gox/mediator/client/clientkey"
	"github.com/goodluck0107/gox/service"

	"github.com/goodluck0107/gox/executor"
	"github.com/goodluck0107/gox/internal/logger"
)

type ClientCommandInActive struct {
	mEventMaker handler.IRoutineInboundEventMaker
	ChannelCtx  service.IChannelContext
}

func NewClientCommandInActive(mEventMaker handler.IRoutineInboundEventMaker, ChannelCtx service.IChannelContext) (this *ClientCommandInActive) {
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
