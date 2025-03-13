package outboundCommands

import (
	"gitee.com/andyxt/gox/extends"
	"gitee.com/andyxt/gox/service"

	"gitee.com/andyxt/gox/internal/logger"
)

type ClientRoutineInboundCmdClose struct {
	uID    int64
	ChlCtx service.IChannelContext
	Desc   string
}

func NewClientRoutineInboundCmdClose(uID int64, ChlCtx service.IChannelContext, Desc string) (this *ClientRoutineInboundCmdClose) {
	this = new(ClientRoutineInboundCmdClose)
	this.uID = uID
	this.ChlCtx = ChlCtx
	this.Desc = Desc
	return
}

func (closeEvent *ClientRoutineInboundCmdClose) QueueId() int64 {
	return closeEvent.uID
}

func (closeEvent *ClientRoutineInboundCmdClose) Wait() (result interface{}, ok bool) {
	return nil, true
}
func (closeEvent *ClientRoutineInboundCmdClose) Exec() {
	logger.Debug("ClientRoutineInboundCmdClose Exec")
	extends.Close(closeEvent.ChlCtx)
}
