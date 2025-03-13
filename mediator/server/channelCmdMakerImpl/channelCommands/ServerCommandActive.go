package channelCommands

import (
	"gitee.com/andyxt/gox/extends"
	"gitee.com/andyxt/gox/service"

	"gitee.com/andyxt/gox/internal/logger"
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
