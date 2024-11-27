package executor

import (
	"fmt"
	"sync"

	"gitee.com/andyxt/gona/logger"
)

type TailHandler struct {
}

func (this *TailHandler) OnExceptionCaught(ctx EventHandlerContext, err error) {
	logger.Error(fmt.Sprintf("exceptionCaught(e = [%#v])\n", err))
}
func (this *TailHandler) OnEventUp(ctx EventHandlerContext, e interface{}) (ret interface{}) {
	return e
}

type HeadHandler struct {
}

func (this *HeadHandler) OnExceptionCaught(ctx EventHandlerContext, err error) {
}
func (this *HeadHandler) OnEventUp(ctx EventHandlerContext, e interface{}) (ret interface{}) {
	return e
}

type DefaultEventPipeline struct {
	Lock *sync.Mutex
	Head *DefaultEventHandlerContext
	Tail *DefaultEventHandlerContext
}

func NewDefaultEventPipeline() (pipeline *DefaultEventPipeline) {
	pipeline = new(DefaultEventPipeline)
	pipeline.Lock = new(sync.Mutex)

	tailHandler := new(TailHandler)
	pipeline.Tail = NewDefaultEventHandlerContext("Tail", pipeline, EventHandler(tailHandler))

	headHandler := new(HeadHandler)
	pipeline.Head = NewDefaultEventHandlerContext("Head", pipeline, EventHandler(headHandler))

	pipeline.Head.Next = pipeline.Tail
	pipeline.Tail.Prev = pipeline.Head
	return
}

func (this *DefaultEventPipeline) AddFirst(name string, handler EventHandler) (pipeline EventPipeline) {
	defer this.Lock.Unlock()
	this.Lock.Lock()
	newCtx := NewDefaultEventHandlerContext(name, this, handler)
	nextCtx := this.Head.Next
	newCtx.Prev = this.Head
	newCtx.Next = nextCtx
	this.Head.Next = newCtx
	nextCtx.Prev = newCtx

	return EventPipeline(this)
}

func (this *DefaultEventPipeline) AddLast(name string, handler EventHandler) (pipeline EventPipeline) {
	defer this.Lock.Unlock()
	this.Lock.Lock()
	newCtx := NewDefaultEventHandlerContext(name, this, handler)
	prev := this.Tail.Prev
	newCtx.Prev = prev
	newCtx.Next = this.Tail
	prev.Next = newCtx
	this.Tail.Prev = newCtx

	return EventPipeline(this)
}

func (this *DefaultEventPipeline) FireUpEvent(event interface{}) (invoker EventInboundInvoker) {
	this.Head.FireUpEvent(event)
	return this
}

func (this *DefaultEventPipeline) FireExceptionCaught(err error) (invoker EventInboundInvoker) {
	this.Head.FireExceptionCaught(err)
	return this
}
