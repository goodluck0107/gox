package executor

type Event interface {
	/**
	 * 为保证事件序列化执行，需要序列化执行的事件必须提供一致的queueId
	 * */
	QueueId() (queueId int64)
	Exec()
	Wait() (result interface{}, ok bool)
}
