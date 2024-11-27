package executor

type IExecutor interface {
	FireEvent(e Event)
	FireEventWait(e Event) (result interface{}, ok bool)
}
