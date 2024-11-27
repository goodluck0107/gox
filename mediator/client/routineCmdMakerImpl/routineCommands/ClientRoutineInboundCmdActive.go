package routineCommands

import (
	"gitee.com/andyxt/gox/extends"
	"gitee.com/andyxt/gox/key"
	"gitee.com/andyxt/gox/service"
	"gitee.com/andyxt/gox/session"

	"gitee.com/andyxt/gona/logger"
)

type ClientRoutineInboundCmdActive struct {
	routineId int64
	ChlCtx    service.IChannelContext
}

func NewClientRoutineInboundCmdActive(routineId int64, ChlCtx service.IChannelContext) (this *ClientRoutineInboundCmdActive) {
	this = new(ClientRoutineInboundCmdActive)
	this.routineId = routineId
	this.ChlCtx = ChlCtx
	return
}

func (activeEvent *ClientRoutineInboundCmdActive) QueueId() int64 {
	return activeEvent.routineId
}

func (activeEvent *ClientRoutineInboundCmdActive) Wait() (result interface{}, ok bool) {
	return nil, true
}

func (activeEvent *ClientRoutineInboundCmdActive) Exec() {
	logger.Debug("ClientRoutineInboundCmdActive Exec")
	uID := activeEvent.ChlCtx.ContextAttr().GetInt64(key.ChannelFireUser)
	iSession := session.GetSession(0, uID)
	if iSession == nil {
		logger.Debug("连接已经被主动关闭，新连接直接关闭")
		extends.Close(activeEvent.ChlCtx)
		return
	}
	oldChlCtx := extends.GetChlCtx(iSession)
	if oldChlCtx == nil {
		extends.ChangeChlCtx(iSession, activeEvent.ChlCtx)
		return
	}
	if !extends.ChannelContextEquals(activeEvent.ChlCtx, oldChlCtx) {
		logger.Debug("已经存在旧连接，直接关闭旧连接")
		extends.Conflict(oldChlCtx)
		extends.Close(oldChlCtx)
	}
	extends.ChangeChlCtx(iSession, activeEvent.ChlCtx)
}
