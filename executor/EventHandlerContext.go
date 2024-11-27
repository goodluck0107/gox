package executor

type EventHandlerContext interface {
	EventInboundInvoker
	Handler() (handler EventHandler)
	Pipeline() (pipeline EventPipeline)
}
