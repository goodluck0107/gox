package schedule

import (
	"gitee.com/andyxt/gox/extends"

	"gitee.com/andyxt/gona/boot/channel"
	"gitee.com/andyxt/gona/logger"
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
