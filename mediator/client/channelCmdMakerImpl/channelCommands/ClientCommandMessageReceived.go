package channelCommands

import (
	"github.com/goodluck0107/gox/code/protocol"
	"github.com/goodluck0107/gox/extends"
	"github.com/goodluck0107/gox/handler"
	"github.com/goodluck0107/gox/mediator/client/clientkey"
	"github.com/goodluck0107/gox/service"

	"github.com/goodluck0107/gox/executor"
	"github.com/goodluck0107/gox/internal/logger"
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
