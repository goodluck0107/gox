package executor

import "gitee.com/andyxt/gona/logger"

const (
	DefaultPoolSize int64 = 10000
	DefaultChanSize int64 = 1000
)

type RoutinePool struct {
	poolSize    int64
	chanSize    int64
	routinePool map[int64]*routine
}

func NewRoutinePool(PoolSize int64, ChanSize int64) (pool *RoutinePool) {
	pool = new(RoutinePool)
	pool.poolSize = PoolSize
	pool.chanSize = ChanSize
	pool.routinePool = make(map[int64]*routine)
	var routineId int64 = 0
	for routineId = 0; routineId < PoolSize; routineId = routineId + 1 {
		routine := NewRoutine(ChanSize)
		startChan := make(chan int, 1)
		routine.Start(startChan)
		<-startChan
		pool.routinePool[routineId] = routine
	}
	return
}
func (this *RoutinePool) ShutDown() {
	for _, value := range this.routinePool {
		value.Close()
	}
}
func (this *RoutinePool) FireEvent(e Event) {
	this.Exec(e)
}

/*发送内部STS消息(玩家不一定在线)*/
func (this *RoutinePool) FireEventWait(e Event) (result interface{}, ok bool) {
	this.Exec(e)
	return e.Wait()
}
func (this *RoutinePool) Exec(event Event) {
	routineId := event.QueueId() % this.poolSize
	if routine, ok := this.routinePool[routineId]; ok {
		routine.Put(event)
		return
	}
	logger.Error("event.QueueId()=", event.QueueId(), " % this.poolSize=", this.poolSize, " routineId==", routineId, " , but has no routine", " event:", event)
	if routine, ok := this.routinePool[0]; ok {
		routine.Put(event)
	}
}
