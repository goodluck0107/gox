package routineCmdMakerImpl

import (
	"github.com/goodluck0107/gox/code/protocol"
	"github.com/goodluck0107/gox/executor"
	"github.com/goodluck0107/gox/mediator/client/routineCmdMakerImpl/inboundCommands"
	"github.com/goodluck0107/gox/service"
)

type ClientInboundEventMakerImpl struct {
	callback inboundCommands.ICallBack
}

func NewClientInboundEventMakerImpl(callback inboundCommands.ICallBack) (impl *ClientInboundEventMakerImpl) {
	impl = new(ClientInboundEventMakerImpl)
	impl.callback = callback
	return
}

// 新连接
func (impl *ClientInboundEventMakerImpl) MakeActiveEvent(routineId int64, Ctx service.IChannelContext) executor.Event {
	return inboundCommands.NewClientChannelUpActiveEvent(routineId, impl, Ctx, impl.callback)
}

// 连接中断
func (impl *ClientInboundEventMakerImpl) MakeInActiveEvent(routineId int64, Ctx service.IChannelContext) executor.Event {
	return inboundCommands.NewClientRoutineInboundCmdInactive(routineId, Ctx, impl.callback)
}

// 收到消息包
func (impl *ClientInboundEventMakerImpl) MakeMessageReceivedEvent(routineId int64, Data protocol.Protocol, Ctx service.IChannelContext) executor.Event {
	return inboundCommands.NewClientChannelUpMsgRecvEvent(routineId, Data, Ctx, impl, impl.callback)
}
