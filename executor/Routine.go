package executor

import (
	"fmt"

	"gitee.com/andyxt/gona/logger"
	"gitee.com/andyxt/gona/utils"
)

type routine struct {
	queue chan Event
}

func NewRoutine(ChanSize int64) (ret *routine) {
	ret = new(routine)
	ret.queue = make(chan Event, ChanSize)
	return
}

func (this *routine) Put(event Event) {
	this.queue <- event
}

func (this *routine) Start(startChan chan int) {
	defer func() {
		if recoverErr := recover(); recoverErr != nil {
			logger.Error(fmt.Sprint(recoverErr, string(utils.Stack(3))))
		}
	}()
	go func() {
		startChan <- 1
		for {
			event, ok := <-this.queue
			if event != nil {
				this.ExecEvent(event)
			}
			if !ok {
				break
			}
		}
	}()
}

func (this *routine) Close() {
	close(this.queue)
}

func (this *routine) ExecEvent(event Event) {
	defer func() {
		if recoverErr := recover(); recoverErr != nil {
			logger.Error(fmt.Sprint(recoverErr, string(utils.Stack(3))))
		}
	}()
	event.Exec()
}
