package routineCmdMakerImpl

import (
	"gitee.com/andyxt/gona/boot/bootc/listener"
	"gitee.com/andyxt/gox/executor"
	"gitee.com/andyxt/gox/handler/protocol"
	"gitee.com/andyxt/gox/mediator/client/routineCmdMakerImpl/routineCommands"
	"gitee.com/andyxt/gox/service"
)

type ClientInboundEventMakerImpl struct {
	connector listener.IConnector
	callback  routineCommands.ICallBack
}

func NewClientInboundEventMakerImpl(connector listener.IConnector, callback routineCommands.ICallBack) (impl *ClientInboundEventMakerImpl) {
	impl = new(ClientInboundEventMakerImpl)
	impl.connector = connector
	impl.callback = callback
	return
}

// 新连接
func (impl *ClientInboundEventMakerImpl) MakeActiveEvent(routineId int64, Ctx service.IChannelContext) executor.Event {
	return routineCommands.NewClientChannelUpActiveEvent(routineId, impl, Ctx, impl.callback)
}

// 连接中断
func (impl *ClientInboundEventMakerImpl) MakeInActiveEvent(routineId int64, Ctx service.IChannelContext) executor.Event {
	return routineCommands.NewClientRoutineInboundCmdInactive(routineId, Ctx, impl.connector)
}

// 收到消息包
func (impl *ClientInboundEventMakerImpl) MakeMessageReceivedEvent(routineId int64, Data protocol.IProtocol, Ctx service.IChannelContext) executor.Event {
	return routineCommands.NewClientChannelUpMsgRecvEvent(routineId, Data, Ctx, impl, impl.callback)
}
