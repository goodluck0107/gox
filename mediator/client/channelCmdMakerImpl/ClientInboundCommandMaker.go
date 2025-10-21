package channelCmdMakerImpl

import (
	"github.com/goodluck0107/gox/handler"
	"github.com/goodluck0107/gox/mediator/client/channelCmdMakerImpl/channelCommands"
	"github.com/goodluck0107/gox/service"
)

type ClientInboundCommandMaker struct {
	mEventMaker handler.IRoutineInboundEventMaker
}

func NewClientInboundCommandMaker(mEventMaker handler.IRoutineInboundEventMaker) (impl *ClientInboundCommandMaker) {
	impl = new(ClientInboundCommandMaker)
	impl.mEventMaker = mEventMaker
	return
}

// 触发异常
func (impl *ClientInboundCommandMaker) MakeExceptionCommand(ctx service.IChannelContext, err error) handler.ICommand {
	return channelCommands.NewClientCommandException(ctx, err)
}

// 新连接
func (impl *ClientInboundCommandMaker) MakeActiveCommand(Ctx service.IChannelContext) handler.ICommand {
	return channelCommands.NewClientCommandActive(impl.mEventMaker, Ctx)
}

// 连接中断
func (impl *ClientInboundCommandMaker) MakeInActiveCommand(Ctx service.IChannelContext) handler.ICommand {
	return channelCommands.NewClientCommandInActive(impl.mEventMaker, Ctx)
}

// 收到消息包
func (impl *ClientInboundCommandMaker) MakeMessageReceivedCommand(Ctx service.IChannelContext, Data interface{}) handler.ICommand {
	return channelCommands.NewClientCommandMessageReceived(impl.mEventMaker, Ctx, Data)
}
