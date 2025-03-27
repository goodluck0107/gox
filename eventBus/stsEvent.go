package eventBus

import "gitee.com/andyxt/gox/internal/logger"

type stsEvent struct {
	routineId int64
	waitChan  chan int8
	evt       string
	params    []interface{}
}

func newStsEvent(routineId int64, syncWait bool,
	evt string, params []interface{}) (this *stsEvent) {
	this = new(stsEvent)
	this.routineId = routineId
	if syncWait {
		this.waitChan = make(chan int8, 1)
	}
	this.evt = evt
	this.params = params
	return
}

func (stsEvent *stsEvent) QueueId() (queueId int64) {
	return stsEvent.routineId
}

func (stsEvent *stsEvent) Wait() (interface{}, bool) {
	if stsEvent.waitChan != nil {
		<-stsEvent.waitChan
	}
	return nil, true
}

func (stsEvent *stsEvent) Exec() {
	for _, f := range onEvents[stsEvent.evt] {
		stsEvent.safeCall(f)
	}
	if stsEvent.waitChan != nil {
		stsEvent.waitChan <- 1
	}
}

func (stsEvent *stsEvent) safeCall(f EventCallback) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("stsEvent.safeCall err:", err)
		}
	}()
	f(stsEvent.params...)
}
