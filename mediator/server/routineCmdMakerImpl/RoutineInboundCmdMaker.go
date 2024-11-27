package routineCmdMakerImpl

import (
	"gitee.com/andyxt/gox/executor"
	"gitee.com/andyxt/gox/handler/protocol"
	"gitee.com/andyxt/gox/handler/schedule"
	"gitee.com/andyxt/gox/mediator/server/routineCmdMakerImpl/routineCommands"
	"gitee.com/andyxt/gox/service"
)

type RoutineInboundCmdMaker struct {
}

func NewRoutineInboundCmdMaker() schedule.IRoutineInboundEventMaker {
	impl := new(RoutineInboundCmdMaker)
	return impl
}

// 新连接
func (impl *RoutineInboundCmdMaker) MakeActiveEvent(routineId int64, Ctx service.IChannelContext) executor.Event {
	return nil
}

// 连接中断
func (impl *RoutineInboundCmdMaker) MakeInActiveEvent(routineId int64, Ctx service.IChannelContext) executor.Event {
	return routineCommands.NewRoutineInboundCmdInactive(routineId, Ctx)
}

// 收到消息包
func (impl *RoutineInboundCmdMaker) MakeMessageReceivedEvent(routineId int64, Data protocol.IProtocol, Ctx service.IChannelContext) executor.Event {
	return routineCommands.NewRoutineInboundCmdMsgRecv(routineId, Data, Ctx)
}
