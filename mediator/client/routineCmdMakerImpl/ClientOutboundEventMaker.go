package routineCmdMakerImpl

import (
	"gitee.com/andyxt/gox/handler/protocol"
	"gitee.com/andyxt/gox/mediator/client/routineCmdMakerImpl/routineCommands"
	"gitee.com/andyxt/gox/service"

	"gitee.com/andyxt/gona/boot/bootc/listener"
	"gitee.com/andyxt/gona/boot/channel"
	"gitee.com/andyxt/gox/executor"
)

type ClientOutboundEventMaker struct {
	connector listener.IConnector
}

func NewClientOutboundEventMaker(connector listener.IConnector) (impl *ClientOutboundEventMaker) {
	impl = new(ClientOutboundEventMaker)
	impl.connector = connector
	return
}

func (impl *ClientOutboundEventMaker) MakeConnectEvent(routineId int64, ip string, port int, uID int64, params map[string]interface{}) executor.Event {
	return routineCommands.NewClientRoutineInboundCmdConnect(routineId, uID, ip, port, params, impl.connector)
}

func (impl *ClientOutboundEventMaker) MakeCloseEvent(routineId int64, uID int64, Desc string) executor.Event {
	return routineCommands.NewClientRoutineInboundCmdClose(routineId, uID, Desc)
}

func (impl *ClientOutboundEventMaker) MakeSendMessageEvent(routineId int64, Data protocol.IProtocol, OnClose bool, PoolKey int64, ChlCtx service.IChannelContext, Desc string) executor.Event {
	return routineCommands.NewClientRoutineOutboundCmdMsgSend(routineId, Data, OnClose, PoolKey, ChlCtx.(channel.ChannelContext), Desc)
}
