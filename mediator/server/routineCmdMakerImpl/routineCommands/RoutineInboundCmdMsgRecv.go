package routineCommands

import (
	"fmt"

	"gitee.com/andyxt/gox/eventBus"
	"gitee.com/andyxt/gox/service"

	"gitee.com/andyxt/gona/logger"
	"gitee.com/andyxt/gox/code/protocol"
	"gitee.com/andyxt/gox/extends"
	"gitee.com/andyxt/gox/mediator/server/evts"
)

type RoutineInboundCmdMsgRecv struct {
	routineId int64
	Data      protocol.Protocol
	Ctx       service.IChannelContext
}

func NewRoutineInboundCmdMsgRecv(routineId int64, Data protocol.Protocol,
	Ctx service.IChannelContext) (this *RoutineInboundCmdMsgRecv) {
	this = new(RoutineInboundCmdMsgRecv)
	this.routineId = routineId
	this.Data = Data
	this.Ctx = Ctx
	return this
}

func (event *RoutineInboundCmdMsgRecv) QueueId() int64 {
	return event.routineId
}
func (event *RoutineInboundCmdMsgRecv) Wait() (interface{}, bool) {
	return nil, true
}
func (event *RoutineInboundCmdMsgRecv) Exec() {
	msgCtx := event.Ctx
	logger.Debug("RoutineInboundCmdMsgRecv Exec-Start", extends.ChannelContextToString(msgCtx))
	if extends.IsConflict(msgCtx) { // 此连接已经被标记为冲突
		logger.Debug("RoutineInboundCmdMsgRecv Exec-End-ChlCtx IsConflict !!!", extends.ChannelContextToString(msgCtx))
		return
	}
	if extends.IsClose(msgCtx) { // 此连接已经被标记为关闭
		logger.Debug("RoutineInboundCmdMsgRecv Exec-End-ChlCtx IsClose !!!", extends.ChannelContextToString(msgCtx))
		return
	}
	if extends.IsLogout(msgCtx) { // 此连接已经被标记为登出
		logger.Debug("RoutineInboundCmdMsgRecv Exec-End-ChlCtx IsLogout !!!", extends.ChannelContextToString(msgCtx))
		return
	}
	if extends.IsSystemKick(msgCtx) { // 此连接已经被标记为踢出
		logger.Debug("RoutineInboundCmdMsgRecv Exec-End-ChlCtx IsSystemKick !!!", extends.ChannelContextToString(msgCtx))
		return
	}
	logger.Debug("RoutineInboundCmdMsgRecv Exec-CallService", extends.ChannelContextToString(msgCtx))
	serviceErr := callService(msgCtx, event.Data)
	if serviceErr != nil {
		return
	}
	logger.Debug("RoutineInboundCmdMsgRecv Exec-End-Success", extends.ChannelContextToString(msgCtx))
}

func callService(chlContext service.IChannelContext, protocol protocol.Protocol) error {
	seqID := protocol.GetSeqID()     // uint32
	msgID := protocol.GetMsgID()     // uint16
	msgData := protocol.GetMsgData() // []byte
	reqContext := service.NewAttr(nil)
	request := service.NewSessionRequest(chlContext, reqContext)
	extends.SetSeqID(request, seqID)
	extends.SetMsgID(request, msgID)
	eventBus.Trigger(evts.EVT_ServiceBefore, request)
	serviceCode := uint32(msgID)
	serviceErr := service.DispatchByCode(serviceCode, request, msgData)
	if serviceErr != nil {
		logger.Error(fmt.Sprintf("chlCtx %v callService %v error %v ", extends.ChannelContextToString(chlContext), serviceCode, serviceErr))
	}
	eventBus.Trigger(evts.EVT_ServiceAfter, request, serviceErr)
	return serviceErr
}
