package channelCmdMakerImpl

import (
	"github.com/goodluck0107/gona/boot/channel"
	"github.com/goodluck0107/gox/handler"
	"github.com/goodluck0107/gox/mediator/client/channelCmdMakerImpl/channelCommands"
)

type ClientOutboundCommandMaker struct {
}

func NewClientOutboundCommandMaker() (impl *ClientOutboundCommandMaker) {
	impl = new(ClientOutboundCommandMaker)
	return
}

// 触发异常
func (impl *ClientOutboundCommandMaker) MakeExceptionCommand(ctx channel.ChannelContext, err error) handler.ICommand {
	return channelCommands.NewClientCommandException(ctx, err)
}

// 请求关闭连接
func (impl *ClientOutboundCommandMaker) MakeCloseCommand(Ctx channel.ChannelContext) handler.ICommand {
	return nil
}

// 下发消息包
func (impl *ClientOutboundCommandMaker) MakeMessageSendCommand(Ctx channel.ChannelContext, Data interface{}) handler.ICommand {
	return nil
}
