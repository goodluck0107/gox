package executor

import (
	"errors"
	"fmt"

	"gitee.com/andyxt/gona/utils"
)

type DefaultEventHandlerContext struct {
	Name     string
	Next     *DefaultEventHandlerContext
	Prev     *DefaultEventHandlerContext
	pipeline EventPipeline
	handler  EventHandler
	attr     map[string]interface{}
}

func NewDefaultEventHandlerContext(name string, pipeline EventPipeline, handler EventHandler) (context *DefaultEventHandlerContext) {
	if handler == nil {
		panic(errors.New("NullPointerException on NewDefaultEventHandlerContext handler is nil"))
	}
	context = new(DefaultEventHandlerContext)
	context.Name = name
	context.pipeline = pipeline
	context.handler = handler
	context.attr = make(map[string]interface{})
	return
}

func (this *DefaultEventHandlerContext) Handler() (handler EventHandler) {
	return this.handler
}

func (this *DefaultEventHandlerContext) Pipeline() (pipeline EventPipeline) {
	return this.pipeline
}

func (this *DefaultEventHandlerContext) Get(key string) (value interface{}) {
	if v, ok := this.attr[key]; ok {
		return v
	}
	return nil
}

func (this *DefaultEventHandlerContext) Set(key string, value interface{}) {
	this.attr[key] = value
}

func (this *DefaultEventHandlerContext) FireUpEvent(event interface{}) (invoker EventInboundInvoker) {
	if event == nil {
		panic(errors.New("NullPointerException on executor.FireUpEvent event is nil"))
	}
	ret, err := this.invokeUpEvent(event)
	if err != nil {
		this.FireExceptionCaught(err)
	} else {
		if next := this.findContextInbound(); next != nil {
			next.FireUpEvent(ret)
		}
	}
	return EventInboundInvoker(this)
}

func (this *DefaultEventHandlerContext) FireExceptionCaught(err error) (invoker EventInboundInvoker) {
	if err == nil {
		panic(errors.New("NullPointerException on FireExceptionCaught event is nil"))
	}
	this.invokeExceptionCaught(err)
	if next := this.findContextInbound(); next != nil {
		next.FireExceptionCaught(err)
	}
	return EventInboundInvoker(this)
}

func (this *DefaultEventHandlerContext) invokeUpEvent(event interface{}) (ret interface{}, err error) {
	handler, _ := this.Handler().(EventInboundHandler)
	defer func() {
		if recoverErr := recover(); recoverErr != nil {
			err = errors.New(fmt.Sprint(recoverErr, string(utils.Stack(3))))
		}
	}()
	ret = handler.OnEventUp(this, event)
	return
}

func (this *DefaultEventHandlerContext) invokeExceptionCaught(err error) {
	handler, _ := this.Handler().(EventInboundHandler)
	defer func() {
		if recoverErr := recover(); recoverErr != nil {
			err = errors.New(fmt.Sprint(recoverErr, string(utils.Stack(3))))
		}
	}()
	handler.OnExceptionCaught(this, err)
}

func (this *DefaultEventHandlerContext) findContextInbound() *DefaultEventHandlerContext {
	var ctx *DefaultEventHandlerContext = this.Next
	for {
		isNil := (ctx == nil)
		if isNil {
			break
		}
		_, isType := ctx.Handler().(EventInboundHandler)
		if isType {
			break
		}
		ctx = ctx.Next
	}
	return ctx
}
