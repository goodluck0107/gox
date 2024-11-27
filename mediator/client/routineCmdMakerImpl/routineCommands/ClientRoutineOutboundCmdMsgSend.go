package routineCommands

import (
	"gitee.com/andyxt/gox/extends"
	"gitee.com/andyxt/gox/handler/protocol"
	"gitee.com/andyxt/gox/service"

	"gitee.com/andyxt/gona/logger"
)

type ClientRoutineOutboundCmdMsgSend struct {
	routineId int64
	Data      protocol.IProtocol
	OnClose   bool // 是否在消息发送完毕后关闭连接
	PoolKey   int64
	ChlCtx    service.IChannelContext
	Desc      string
}

func NewClientRoutineOutboundCmdMsgSend(routineId int64, Data protocol.IProtocol, OnClose bool, PoolKey int64, ChlCtx service.IChannelContext, Desc string) (this *ClientRoutineOutboundCmdMsgSend) {
	this = new(ClientRoutineOutboundCmdMsgSend)
	this.routineId = routineId
	this.Data = Data
	this.OnClose = OnClose
	this.PoolKey = PoolKey
	this.ChlCtx = ChlCtx
	this.Desc = Desc
	return
}

func (msgSendEvent *ClientRoutineOutboundCmdMsgSend) QueueId() int64 {
	return msgSendEvent.routineId
}
func (msgSendEvent *ClientRoutineOutboundCmdMsgSend) Wait() (result interface{}, ok bool) {
	return nil, true
}
func (msgSendEvent *ClientRoutineOutboundCmdMsgSend) Exec() {
	if msgSendEvent.ChlCtx == nil {
		logger.Error("ClientRoutineOutboundCmdMsgSend Exec", "Fail:", " ChlCtx == nil")
		return
	}
	if msgSendEvent.Data == nil {
		logger.Debug("ClientRoutineOutboundCmdMsgSend Exec", extends.ChannelContextToString(msgSendEvent.ChlCtx), "Fail:", "Data == nil")
		return
	}
	if extends.IsClose(msgSendEvent.ChlCtx) {
		logger.Debug("ClientRoutineOutboundCmdMsgSend Exec", extends.ChannelContextToString(msgSendEvent.ChlCtx), "Fail:", "extends.IsClose(ChlCtx)")
		return
	}
	logger.Debug("ClientRoutineOutboundCmdMsgSend Exec", extends.ChannelContextToString(msgSendEvent.ChlCtx), " Success")
	msgSendEvent.ChlCtx.Write(msgSendEvent.Data)
	if msgSendEvent.OnClose {
		logger.Debug("ClientRoutineOutboundCmdMsgSend Exec", extends.ChannelContextToString(msgSendEvent.ChlCtx), " Success & CloseChannel")
		extends.Close(msgSendEvent.ChlCtx)
	}
}
