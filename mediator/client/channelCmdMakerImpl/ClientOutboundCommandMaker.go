package channelCmdMakerImpl

import (
	"gitee.com/andyxt/gona/boot/channel"
	"gitee.com/andyxt/gox/handler"
	"gitee.com/andyxt/gox/mediator/client/channelCmdMakerImpl/channelCommands"
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
