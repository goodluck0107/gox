package routineCommands

import (
	"github.com/goodluck0107/gox/eventBus"
	"github.com/goodluck0107/gox/extends"
	"github.com/goodluck0107/gox/internal/logger"
	"github.com/goodluck0107/gox/mediator/server/evts"
	"github.com/goodluck0107/gox/service"
)

type RoutineInboundCmdInactive struct {
	routineId int64
	ChlCtx    service.IChannelContext
}

func NewRoutineInboundCmdInactive(routineId int64, ChlCtx service.IChannelContext) (this *RoutineInboundCmdInactive) {
	this = new(RoutineInboundCmdInactive)
	this.routineId = routineId
	this.ChlCtx = ChlCtx
	return
}

func (event *RoutineInboundCmdInactive) QueueId() int64 {
	return event.routineId
}

func (event *RoutineInboundCmdInactive) Wait() (interface{}, bool) {
	return nil, true
}

func (event *RoutineInboundCmdInactive) Exec() {
	logger.Info("RoutineInboundCmdInactive Exec-Start", extends.ChannelContextToString(event.ChlCtx))
	extends.Offlie(event.ChlCtx)          // 标记此连接已经离线
	if extends.IsConflict(event.ChlCtx) { // 此连接已经被标记为冲突
		logger.Info("RoutineInboundCmdInactive Exec-End-ChlCtx IsConflict !!!", extends.ChannelContextToString(event.ChlCtx))
		return
	}
	eventBus.Trigger(evts.EVT_Inactive, event.routineId, event.ChlCtx)
}
