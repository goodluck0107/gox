package inboundCommands

import (
	"github.com/goodluck0107/gox/mediator/client/clientkey"
	"github.com/goodluck0107/gox/service"

	"github.com/goodluck0107/gox/internal/logger"
)

type ClientRoutineInboundCmdInactive struct {
	routineId int64
	ChlCtx    service.IChannelContext
	callback  ICallBack
}

func NewClientRoutineInboundCmdInactive(routineId int64,
	ChlCtx service.IChannelContext, callback ICallBack) (this *ClientRoutineInboundCmdInactive) {
	this = new(ClientRoutineInboundCmdInactive)
	this.routineId = routineId
	this.ChlCtx = ChlCtx
	this.callback = callback
	return
}

func (inactiveEvent *ClientRoutineInboundCmdInactive) QueueId() int64 {
	return inactiveEvent.routineId
}

func (inactiveEvent *ClientRoutineInboundCmdInactive) Wait() (result interface{}, ok bool) {
	return nil, true
}

func (inactiveEvent *ClientRoutineInboundCmdInactive) Exec() {
	logger.Debug("ClientRoutineInboundCmdInactive Exec")
	logger.Debug("ClientRoutineInboundCmdInactive Exec", "连接中断")
	uID := inactiveEvent.ChlCtx.ContextAttr().GetInt64(clientkey.KeyFireUser)
	inactiveEvent.callback.ConnectInactive(uID, inactiveEvent.ChlCtx)
}
