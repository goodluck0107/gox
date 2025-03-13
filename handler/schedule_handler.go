package handler

import (
	"gitee.com/andyxt/gox/code/protocol"
	"gitee.com/andyxt/gox/executor"
	"gitee.com/andyxt/gox/extends"
	"gitee.com/andyxt/gox/service"

	"gitee.com/andyxt/gona/boot/channel"
	"gitee.com/andyxt/gox/internal/logger"
)

// UpBase --->
type InBoundExecutionHandler struct {
	mInboundCommandMaker IChannelInboundCommandMaker
}

func NewInBoundExecutionHandler(mInboundCommandMaker IChannelInboundCommandMaker) (this *InBoundExecutionHandler) {
	this = new(InBoundExecutionHandler)
	this.mInboundCommandMaker = mInboundCommandMaker
	return
}

func (handler *InBoundExecutionHandler) ExceptionCaught(ctx channel.ChannelContext, err error) {
	logger.Debug("InBoundExecutionHandler ExceptionCaught", extends.ChannelContextToString(ctx))
	handler.mInboundCommandMaker.MakeExceptionCommand(ctx, err).Exec()
}

func (handler *InBoundExecutionHandler) ChannelActive(ctx channel.ChannelContext) (goonNext bool) {
	logger.Debug("InBoundExecutionHandler ChannelActive", extends.ChannelContextToString(ctx))
	handler.mInboundCommandMaker.MakeActiveCommand(ctx).Exec()
	return
}

func (handler *InBoundExecutionHandler) ChannelInactive(ctx channel.ChannelContext) (goonNext bool) {
	logger.Debug("InBoundExecutionHandler ChannelInactive", extends.ChannelContextToString(ctx))
	handler.mInboundCommandMaker.MakeInActiveCommand(ctx).Exec()
	return
}

func (handler *InBoundExecutionHandler) MessageReceived(ctx channel.ChannelContext, e interface{}) (ret interface{}, goonNext bool) {
	logger.Debug("InBoundExecutionHandler MessageReceived", extends.ChannelContextToString(ctx))
	handler.mInboundCommandMaker.MakeMessageReceivedCommand(ctx, e).Exec()
	return
}

// IChannelInboundCommandMaker 创建的所有Commands全部都在Tcp读协程中执行
// IChannelOutboundCommandMaker 创建的所有Commands全部都在Tcp写协程中执行

type ICommand interface {
	Exec()
}

type IChannelInboundCommandMaker interface {
	//触发异常
	MakeExceptionCommand(ctx service.IChannelContext, err error) ICommand

	//新连接
	MakeActiveCommand(Ctx service.IChannelContext) ICommand
	//连接中断
	MakeInActiveCommand(Ctx service.IChannelContext) ICommand
	//收到消息包
	MakeMessageReceivedCommand(Ctx service.IChannelContext, Data interface{}) ICommand
}

type IChannelOutboundCommandMaker interface {
	// 触发异常
	MakeExceptionCommand(ctx service.IChannelContext, err error) ICommand

	// 请求关闭连接
	MakeCloseCommand(Ctx service.IChannelContext) ICommand
	// 下发消息包
	MakeMessageSendCommand(Ctx service.IChannelContext, Data interface{}) ICommand
}

// Inbound
type IRoutineInboundEventMaker interface {
	//收到消息包
	MakeMessageReceivedEvent(routineId int64, Data protocol.Protocol, Ctx service.IChannelContext) executor.Event
	//新连接
	MakeActiveEvent(routineId int64, Ctx service.IChannelContext) executor.Event
	//连接中断
	MakeInActiveEvent(routineId int64, Ctx service.IChannelContext) executor.Event
}

// Outbound
type IRoutineOutboundEventMaker interface {
	//发起连接
	MakeConnectEvent(routineId int64, ip string, port int, uID int64, params map[string]interface{}) executor.Event
	//关闭连接
	MakeCloseEvent(routineId int64, uID int64, Desc string) executor.Event
	//下发消息包:OnClose是否在消息发送完毕后关闭连接
	MakeSendMessageEvent(routineId int64, Data protocol.Protocol, OnClose bool, PoolKey int64, ChlCtx service.IChannelContext, Desc string) executor.Event
}
