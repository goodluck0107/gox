package channelCommands

import (
	"github.com/goodluck0107/gox/extends"
	"github.com/goodluck0107/gox/handler"
	"github.com/goodluck0107/gox/mediator/client/clientkey"
	"github.com/goodluck0107/gox/service"

	"github.com/goodluck0107/gox/executor"
	"github.com/goodluck0107/gox/internal/logger"
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
