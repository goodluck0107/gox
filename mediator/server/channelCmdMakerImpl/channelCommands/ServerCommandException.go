package channelCommands

import (
	"github.com/goodluck0107/gox/extends"
	"github.com/goodluck0107/gox/service"

	"github.com/goodluck0107/gox/internal/logger"
)

type ServerCommandException struct {
	ChannelCtx service.IChannelContext
	err        error
}

func NewServerCommandException(ChannelCtx service.IChannelContext, err error) (this *ServerCommandException) {
	this = new(ServerCommandException)
	this.ChannelCtx = ChannelCtx
	this.err = err
	return
}

func (event *ServerCommandException) Exec() {
	if event.ChannelCtx == nil {
		logger.Debug("ServerCommandException Exec 1!")
		return
	}
	if event.err == nil {
		logger.Debug("ServerCommandException Exec 2!", extends.ChannelContextToString(event.ChannelCtx))
		return
	}
	logger.Error("ServerCommandException Exec 3!", extends.ChannelContextToString(event.ChannelCtx), "关闭连接,关闭原因：ServerCommandException ExceptionCaught:", event.err)
	extends.Close(event.ChannelCtx)
}
