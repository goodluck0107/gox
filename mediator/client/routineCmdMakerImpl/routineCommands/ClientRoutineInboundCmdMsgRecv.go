package routineCommands

import (
	"gitee.com/andyxt/gox/handler/protocol"
	"gitee.com/andyxt/gox/service"
)

type ClientRoutineInboundCmdMsgRecv struct {
	routineId int64
	Data      protocol.IProtocol
	Ctx       service.IChannelContext
}

func NewClientRoutineInboundCmdMsgRecv(routineId int64, Data protocol.IProtocol,
	Ctx service.IChannelContext) (this *ClientRoutineInboundCmdMsgRecv) {
	this = new(ClientRoutineInboundCmdMsgRecv)
	this.routineId = routineId
	this.Data = Data
	this.Ctx = Ctx
	return this
}

func (msgRecvEvent *ClientRoutineInboundCmdMsgRecv) QueueId() int64 {
	return msgRecvEvent.routineId
}
func (msgRecvEvent *ClientRoutineInboundCmdMsgRecv) Wait() (result interface{}, ok bool) {
	return nil, true
}
func (msgRecvEvent *ClientRoutineInboundCmdMsgRecv) Exec() {

}
