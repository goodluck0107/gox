package server

import (
	"github.com/goodluck0107/gox/code/protocol"
	"github.com/goodluck0107/gox/eventBus"
	"github.com/goodluck0107/gox/executor"
	"github.com/goodluck0107/gox/handler"

	"github.com/goodluck0107/gox/extends"
	"github.com/goodluck0107/gox/mediator/server/evts"
	"github.com/goodluck0107/gox/mediator/server/routineCmdMakerImpl"

	"github.com/goodluck0107/gox/service"
)

var mRoutineInboundCmdMaker handler.IRoutineInboundEventMaker = routineCmdMakerImpl.NewRoutineInboundCmdMaker()
var mRoutineOutboundCmdMaker handler.IRoutineOutboundEventMaker = routineCmdMakerImpl.NewRoutineOutboundCmdMaker()

func Response(ChlCtx service.IChannelContext, v protocol.Protocol) error {
	executor.FireEvent(mRoutineOutboundCmdMaker.MakeSendMessageEvent(extends.UID(ChlCtx), v, false, extends.UID(ChlCtx), ChlCtx, ""))
	return nil
}

func ResponseClose(ChlCtx service.IChannelContext, v protocol.Protocol, desc string) error {
	executor.FireEvent(mRoutineOutboundCmdMaker.MakeSendMessageEvent(extends.UID(ChlCtx), v, true, extends.UID(ChlCtx), ChlCtx, desc))
	return nil
}

func Push(ChlCtx service.IChannelContext, v protocol.Protocol) error {
	executor.FireEvent(mRoutineOutboundCmdMaker.MakeSendMessageEvent(extends.UID(ChlCtx), v, false, extends.UID(ChlCtx), ChlCtx, ""))
	return nil
}

func PushClose(ChlCtx service.IChannelContext, v protocol.Protocol, desc string) error {
	executor.FireEvent(mRoutineOutboundCmdMaker.MakeSendMessageEvent(extends.UID(ChlCtx), v, true, extends.UID(ChlCtx), ChlCtx, desc))
	return nil
}

// OnClose 监听连接中断
func OnClose(closeFunc func(playerID int64, chlCtx service.IChannelContext)) {
	eventBus.On(evts.EVT_Inactive, func(data ...interface{}) {
		playerID := data[0].(int64)
		channelContext := data[1].(service.IChannelContext)
		closeFunc(playerID, channelContext)
	})
}

func BeforeService(beforeFunc func(request service.IServiceRequest, playerID int64, msgProtoID uint16, msgSeqID uint32)) {
	eventBus.On(evts.EVT_ServiceBefore, func(data ...interface{}) {
		request := data[0].(service.IServiceRequest)
		playerID := extends.UID(request.ChannelContext())
		msgProtoID := extends.MsgID(request)
		msgSeqID := extends.SeqID(request)
		beforeFunc(request, playerID, msgProtoID, msgSeqID)
	})
}

func AfterService(afterFunc func(request service.IServiceRequest, playerID int64, msgProtoID uint16, msgSeqID uint32, serviceE error)) {
	eventBus.On(evts.EVT_ServiceAfter, func(data ...interface{}) {
		request := data[0].(service.IServiceRequest)
		playerID := extends.UID(request.ChannelContext())
		msgProtoID := extends.MsgID(request)
		msgSeqID := extends.SeqID(request)
		if data[1] != nil {
			afterFunc(request, playerID, msgProtoID, msgSeqID, data[1].(error))
			return
		}
		afterFunc(request, playerID, msgProtoID, msgSeqID, nil)
	})
}
