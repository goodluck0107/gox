package executor

type EventPipeline interface {
	EventInboundInvoker
	AddFirst(name string, handler EventHandler) (pipeline EventPipeline)
	AddLast(name string, handler EventHandler) (pipeline EventPipeline)
}
