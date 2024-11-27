package channelCommands

import (
	"gitee.com/andyxt/gox/extends"
	"gitee.com/andyxt/gox/service"

	"gitee.com/andyxt/gona/logger"
)

type ClientCommandException struct {
	ChannelCtx service.IChannelContext
	err        error
}

func NewClientCommandException(ChannelCtx service.IChannelContext, err error) (this *ClientCommandException) {
	this = new(ClientCommandException)
	this.ChannelCtx = ChannelCtx
	this.err = err
	return
}

func (event *ClientCommandException) Exec() {
	logger.Debug("ClientCommandException Exec", extends.ChannelContextToString(event.ChannelCtx))
	if event.ChannelCtx == nil || event.err == nil {
		return
	}
	logger.Error("关闭连接：", " 关闭原因：ClientCommandException ExceptionCaught:", event.err, extends.ChannelContextToString(event.ChannelCtx))
	extends.Close(event.ChannelCtx)
}
