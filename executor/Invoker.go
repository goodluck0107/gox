package executor

type EventInboundInvoker interface {
	FireExceptionCaught(err error) (invoker EventInboundInvoker)
	FireUpEvent(event interface{}) (invoker EventInboundInvoker)
}
