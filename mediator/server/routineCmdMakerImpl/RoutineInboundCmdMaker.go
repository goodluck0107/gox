package routineCmdMakerImpl

import (
	"github.com/goodluck0107/gox/code/protocol"
	"github.com/goodluck0107/gox/executor"
	"github.com/goodluck0107/gox/handler"
	"github.com/goodluck0107/gox/mediator/server/routineCmdMakerImpl/routineCommands"
	"github.com/goodluck0107/gox/service"
)

type RoutineInboundCmdMaker struct {
}

func NewRoutineInboundCmdMaker() handler.IRoutineInboundEventMaker {
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
func (impl *RoutineInboundCmdMaker) MakeMessageReceivedEvent(routineId int64, Data protocol.Protocol, Ctx service.IChannelContext) executor.Event {
	return routineCommands.NewRoutineInboundCmdMsgRecv(routineId, Data, Ctx)
}
