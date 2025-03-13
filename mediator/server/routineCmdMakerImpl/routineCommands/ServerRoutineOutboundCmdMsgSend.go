package routineCommands

import (
	"gitee.com/andyxt/gox/code/protocol"
	"gitee.com/andyxt/gox/extends"
	"gitee.com/andyxt/gox/service"

	"gitee.com/andyxt/gox/internal/logger"
)

type ServerRoutineOutboundCmdMsgSend struct {
	routineId int64
	Data      protocol.Protocol
	OnClose   bool // 是否在消息发送完毕后关闭连接
	PoolKey   int64
	ChlCtx    service.IChannelContext
	Desc      string
}

func NewServerRoutineOutboundCmdMsgSend(routineId int64, Data protocol.Protocol, OnClose bool, PoolKey int64, ChlCtx service.IChannelContext, Desc string) (this *ServerRoutineOutboundCmdMsgSend) {
	this = new(ServerRoutineOutboundCmdMsgSend)
	this.routineId = routineId
	this.Data = Data
	this.OnClose = OnClose
	this.PoolKey = PoolKey
	this.ChlCtx = ChlCtx
	this.Desc = Desc
	return
}

func (this *ServerRoutineOutboundCmdMsgSend) QueueId() int64 {
	return this.routineId
}
func (this *ServerRoutineOutboundCmdMsgSend) Wait() (interface{}, bool) {
	return nil, true
}
func (this *ServerRoutineOutboundCmdMsgSend) Exec() {
	//	logger.Debug(ctx.GetPoolKey(), " OnChannelMsgSend:", data)
	//fmt.Println("want OnChannelMsgSend msgId",  fmt.Sprintf("0x%04x", data.GetMessageId()))
	if this.ChlCtx == nil {
		logger.Error("ServerRoutineOutboundCmdMsgSend Exec", "Fail:", " ChlCtx == nil")
		return
	}
	if this.Data == nil {
		logger.Debug("ServerRoutineOutboundCmdMsgSend Exec", extends.ChannelContextToString(this.ChlCtx), "Fail:", "Data == nil")
		return
	}
	if extends.IsClose(this.ChlCtx) {
		logger.Debug("ServerRoutineOutboundCmdMsgSend Exec", extends.ChannelContextToString(this.ChlCtx), "Fail:", "extends.IsClose(ChlCtx)")
		return
	}
	this.ChlCtx.Write(this.Data)
	//fmt.Println("over OnChannelMsgSend msgId",  fmt.Sprintf("0x%04x", data.GetMessageId()))
	if this.OnClose {
		logger.Debug("ServerRoutineOutboundCmdMsgSend Exec", extends.ChannelContextToString(this.ChlCtx), " Success & CloseChannel")
		extends.Close(this.ChlCtx)
		return
	}
	logger.Debug("ServerRoutineOutboundCmdMsgSend Exec", extends.ChannelContextToString(this.ChlCtx), " Success")
}
