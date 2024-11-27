package executor

type EventHandler interface {
	OnExceptionCaught(ctx EventHandlerContext, err error)
}

type EventInboundHandler interface {
	EventHandler
	OnEventUp(ctx EventHandlerContext, e interface{}) (ret interface{})
}
