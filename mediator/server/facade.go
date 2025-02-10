package server

import (
	"gitee.com/andyxt/gox/eventBus"
	"gitee.com/andyxt/gox/executor"
	"gitee.com/andyxt/gox/handler/schedule"
	"gitee.com/andyxt/gox/message"

	"gitee.com/andyxt/gox/extends"
	"gitee.com/andyxt/gox/mediator/server/evts"
	"gitee.com/andyxt/gox/mediator/server/routineCmdMakerImpl"

	"gitee.com/andyxt/gox/service"

	"google.golang.org/protobuf/reflect/protoreflect"
)

var routeTypeClient message.Type = 0
var messageTypeClose message.Type = 0
var messageTypeProto message.Type = 1
var mRoutineInboundCmdMaker schedule.IRoutineInboundEventMaker = routineCmdMakerImpl.NewRoutineInboundCmdMaker()
var mRoutineOutboundCmdMaker schedule.IRoutineOutboundEventMaker = routineCmdMakerImpl.NewRoutineOutboundCmdMaker()

func Response(ChlCtx service.IChannelContext, msgSeqID uint32, msgMsgID uint16, v interface{}) error {
	executor.FireEvent(mRoutineOutboundCmdMaker.MakeSendMessageEvent(extends.UID(ChlCtx), message.NewMessage(routeTypeClient, messageTypeProto, 1, msgSeqID,
		msgMsgID, v.(protoreflect.ProtoMessage)), false, extends.UID(ChlCtx), ChlCtx, ""))
	return nil
}

func ResponseClose(ChlCtx service.IChannelContext, msgSeqID uint32, msgMsgID uint16, v interface{}, desc string) error {
	executor.FireEvent(mRoutineOutboundCmdMaker.MakeSendMessageEvent(extends.UID(ChlCtx), message.NewMessage(routeTypeClient, messageTypeClose, 1, msgSeqID,
		msgMsgID, v.(protoreflect.ProtoMessage)), true, extends.UID(ChlCtx), ChlCtx, desc))
	return nil
}

func Push(ChlCtx service.IChannelContext, msgMsgID uint16, v interface{}) error {
	executor.FireEvent(mRoutineOutboundCmdMaker.MakeSendMessageEvent(extends.UID(ChlCtx), message.NewMessage(routeTypeClient, messageTypeProto, 1, 0,
		msgMsgID, v.(protoreflect.ProtoMessage)), false, extends.UID(ChlCtx), ChlCtx, ""))
	return nil
}

func PushClose(ChlCtx service.IChannelContext, msgMsgID uint16, v interface{}, desc string) error {
	executor.FireEvent(mRoutineOutboundCmdMaker.MakeSendMessageEvent(extends.UID(ChlCtx), message.NewMessage(routeTypeClient, messageTypeClose, 1, 0,
		msgMsgID, v.(protoreflect.ProtoMessage)), true, extends.UID(ChlCtx), ChlCtx, desc))
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
		serviceE := data[1].(error)
		playerID := extends.UID(request.ChannelContext())
		msgProtoID := extends.MsgID(request)
		msgSeqID := extends.SeqID(request)
		afterFunc(request, playerID, msgProtoID, msgSeqID, serviceE)
	})
}
