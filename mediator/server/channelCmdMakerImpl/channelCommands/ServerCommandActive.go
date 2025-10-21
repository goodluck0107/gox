package channelCommands

import (
	"github.com/goodluck0107/gox/extends"
	"github.com/goodluck0107/gox/service"

	"github.com/goodluck0107/gox/internal/logger"
)

type ServerCommandActive struct {
	ChannelCtx service.IChannelContext
}

func NewServerCommandActive(ChannelCtx service.IChannelContext) (this *ServerCommandActive) {
	this = new(ServerCommandActive)
	this.ChannelCtx = ChannelCtx
	return
}

func (event *ServerCommandActive) Exec() {
	logger.Debug("ServerCommandActive Exec", extends.ChannelContextToString(event.ChannelCtx))
}
