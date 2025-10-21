package channelCmdMakerImpl

import (
	"github.com/goodluck0107/gox/handler"
	"github.com/goodluck0107/gox/mediator/server/channelCmdMakerImpl/channelCommands"
	"github.com/goodluck0107/gox/service"
)

type ChannelInboundCmdMaker struct {
	mLoginMessage channelCommands.ILoginMessage
	mEventMaker   handler.IRoutineInboundEventMaker
}

func NewChannelInboundCmdMaker(mLoginMessage channelCommands.ILoginMessage, mEventMaker handler.IRoutineInboundEventMaker) (impl *ChannelInboundCmdMaker) {
	impl = new(ChannelInboundCmdMaker)
	impl.mLoginMessage = mLoginMessage
	impl.mEventMaker = mEventMaker
	return
}

// 触发异常
func (impl *ChannelInboundCmdMaker) MakeExceptionCommand(ctx service.IChannelContext, err error) handler.ICommand {
	return channelCommands.NewServerCommandException(ctx, err)
}

// 新连接
func (impl *ChannelInboundCmdMaker) MakeActiveCommand(Ctx service.IChannelContext) handler.ICommand {
	return channelCommands.NewServerCommandActive(Ctx)
}

// 连接中断
func (impl *ChannelInboundCmdMaker) MakeInActiveCommand(Ctx service.IChannelContext) handler.ICommand {
	return channelCommands.NewServerCommandInActive(impl.mEventMaker, Ctx)
}

// 收到消息包
func (impl *ChannelInboundCmdMaker) MakeMessageReceivedCommand(Ctx service.IChannelContext, Data interface{}) handler.ICommand {
	return channelCommands.NewChannelInboundCmdMsgRecv(impl.mEventMaker, Ctx, Data, impl.mLoginMessage)
}
