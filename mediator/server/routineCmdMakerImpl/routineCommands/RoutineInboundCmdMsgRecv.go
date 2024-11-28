package routineCommands

import (
	"gitee.com/andyxt/gox/message"
	"gitee.com/andyxt/gox/service"

	"gitee.com/andyxt/gona/logger"
	"gitee.com/andyxt/gox/extends"
	"gitee.com/andyxt/gox/handler/protocol"
	"gitee.com/andyxt/gox/session"
)

type RoutineInboundCmdMsgRecv struct {
	routineId int64
	Data      protocol.IProtocol
	Ctx       service.IChannelContext
}

func NewRoutineInboundCmdMsgRecv(routineId int64, Data protocol.IProtocol,
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
	logger.Debug("RoutineInboundCmdMsgRecv Exec-Start", extends.ChannelContextToString(msgCtx), "sessionCount", session.GetCount())
	if extends.IsConflict(msgCtx) { // 此连接已经被标记为冲突
		logger.Debug("RoutineInboundCmdMsgRecv Exec-End-ChlCtx IsConflict !!!", extends.ChannelContextToString(msgCtx), "sessionCount", session.GetCount())
		return
	}
	if extends.IsClose(msgCtx) { // 此连接已经被标记为关闭
		logger.Debug("RoutineInboundCmdMsgRecv Exec-End-ChlCtx IsClose !!!", extends.ChannelContextToString(msgCtx), "sessionCount", session.GetCount())
		return
	}
	if extends.IsLogout(msgCtx) { // 此连接已经被标记为登出
		logger.Debug("RoutineInboundCmdMsgRecv Exec-End-ChlCtx IsLogout !!!", extends.ChannelContextToString(msgCtx), "sessionCount", session.GetCount())
		return
	}
	if extends.IsSystemKick(msgCtx) { // 此连接已经被标记为踢出
		logger.Debug("RoutineInboundCmdMsgRecv Exec-End-ChlCtx IsSystemKick !!!", extends.ChannelContextToString(msgCtx), "sessionCount", session.GetCount())
		return
	}
	logger.Debug("RoutineInboundCmdMsgRecv Exec-CallService", extends.ChannelContextToString(msgCtx), "sessionCount", session.GetCount())
	serviceErr := callService(msgCtx, event.Data.(*message.Message))
	if serviceErr != nil {
		logger.Debug("RoutineInboundCmdMsgRecv Exec-End-CallServiceError !!!", extends.ChannelContextToString(msgCtx), "serviceError:", serviceErr)
		return
	}
	logger.Debug("RoutineInboundCmdMsgRecv Exec-End-Success", extends.ChannelContextToString(msgCtx), "sessionCount", session.GetCount())
}
