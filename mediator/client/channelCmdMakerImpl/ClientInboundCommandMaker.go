package channelCmdMakerImpl

import (
	"gitee.com/andyxt/gox/handler/schedule"
	"gitee.com/andyxt/gox/mediator/client/channelCmdMakerImpl/channelCommands"
	"gitee.com/andyxt/gox/service"
)

type ClientInboundCommandMaker struct {
	mEventMaker schedule.IRoutineInboundEventMaker
}

func NewClientInboundCommandMaker(mEventMaker schedule.IRoutineInboundEventMaker) (impl *ClientInboundCommandMaker) {
	impl = new(ClientInboundCommandMaker)
	impl.mEventMaker = mEventMaker
	return
}

// 触发异常
func (impl *ClientInboundCommandMaker) MakeExceptionCommand(ctx service.IChannelContext, err error) schedule.ICommand {
	return channelCommands.NewClientCommandException(ctx, err)
}

// 新连接
func (impl *ClientInboundCommandMaker) MakeActiveCommand(Ctx service.IChannelContext) schedule.ICommand {
	return channelCommands.NewClientCommandActive(impl.mEventMaker, Ctx)
}

// 连接中断
func (impl *ClientInboundCommandMaker) MakeInActiveCommand(Ctx service.IChannelContext) schedule.ICommand {
	return channelCommands.NewClientCommandInActive(impl.mEventMaker, Ctx)
}

// 收到消息包
func (impl *ClientInboundCommandMaker) MakeMessageReceivedCommand(Ctx service.IChannelContext, Data interface{}) schedule.ICommand {
	return channelCommands.NewClientCommandMessageReceived(impl.mEventMaker, Ctx, Data)
}
