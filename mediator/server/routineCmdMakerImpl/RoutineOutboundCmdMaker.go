package routineCmdMakerImpl

import (
	"gitee.com/andyxt/gox/executor"
	"gitee.com/andyxt/gox/handler/protocol"
	"gitee.com/andyxt/gox/mediator/server/routineCmdMakerImpl/routineCommands"
	"gitee.com/andyxt/gox/service"
)

type RoutineOutboundCmdMaker struct {
}

func NewRoutineOutboundCmdMaker() (impl *RoutineOutboundCmdMaker) {
	impl = new(RoutineOutboundCmdMaker)
	return
}

func (impl *RoutineOutboundCmdMaker) MakeConnectEvent(routineId int64, ip string, port int, uID int64, params map[string]interface{}) executor.Event {
	return nil
}

func (impl *RoutineOutboundCmdMaker) MakeCloseEvent(routineId int64, uID int64, Desc string) executor.Event {
	return nil
}

func (impl *RoutineOutboundCmdMaker) MakeSendMessageEvent(routineId int64, Data protocol.IProtocol, OnClose bool, PoolKey int64, ChlCtx service.IChannelContext, Desc string) executor.Event {
	return routineCommands.NewServerRoutineOutboundCmdMsgSend(routineId, Data, OnClose, PoolKey, ChlCtx, Desc)
}
