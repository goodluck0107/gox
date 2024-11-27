package channelCmdMakerImpl

import (
	"gitee.com/andyxt/gox/handler/schedule"
	"gitee.com/andyxt/gox/mediator/server/channelCmdMakerImpl/channelCommands"
	"gitee.com/andyxt/gox/service"
)

type ChannelInboundCmdMaker struct {
	mLoginMessage channelCommands.ILoginMessage
	mEventMaker   schedule.IRoutineInboundEventMaker
}

func NewChannelInboundCmdMaker(mLoginMessage channelCommands.ILoginMessage, mEventMaker schedule.IRoutineInboundEventMaker) (impl *ChannelInboundCmdMaker) {
	impl = new(ChannelInboundCmdMaker)
	impl.mLoginMessage = mLoginMessage
	impl.mEventMaker = mEventMaker
	return
}

// 触发异常
func (impl *ChannelInboundCmdMaker) MakeExceptionCommand(ctx service.IChannelContext, err error) schedule.ICommand {
	return channelCommands.NewServerCommandException(ctx, err)
}

// 新连接
func (impl *ChannelInboundCmdMaker) MakeActiveCommand(Ctx service.IChannelContext) schedule.ICommand {
	return channelCommands.NewServerCommandActive(Ctx)
}

// 连接中断
func (impl *ChannelInboundCmdMaker) MakeInActiveCommand(Ctx service.IChannelContext) schedule.ICommand {
	return channelCommands.NewServerCommandInActive(impl.mEventMaker, Ctx)
}

// 收到消息包
func (impl *ChannelInboundCmdMaker) MakeMessageReceivedCommand(Ctx service.IChannelContext, Data interface{}) schedule.ICommand {
	return channelCommands.NewChannelInboundCmdMsgRecv(impl.mEventMaker, Ctx, Data, impl.mLoginMessage)
}
