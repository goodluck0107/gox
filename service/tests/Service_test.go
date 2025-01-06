package main

import (
	"reflect"
	"testing"

	"gitee.com/andyxt/gox/service"

	"gitee.com/andyxt/gona/boot/channel"
	"gitee.com/andyxt/gona/logger"
)

type TestParam1 struct{}
type TestParam2 struct{}
type DemoService struct{}

func (t *DemoService) RouteForTest1() (string, uint32) {
	return "/Test1", 1
}
func (t *DemoService) RouteForTest2() (string, uint32) {
	return "/Test2", 2
}
func (t *DemoService) Test1(chlCtx channel.ChannelContext, p *TestParam1, attr channel.IAttr) error {
	logger.Info("Test1")
	return nil
}
func (t *DemoService) Test2(chlCtx channel.ChannelContext, p *TestParam2, attr channel.IAttr) error {
	logger.Info("Test2")
	return nil
}

func TestService(t *testing.T) {
	service.Init(NewDefaultHandleChecker())
	service.Register(new(DemoService))
	service.DispatchByCode(1, new(channel.DefaultChannelHandlerContext), new(TestParam1), channel.NewAttr(nil))
	service.DispatchByPath("/Test2", new(channel.DefaultChannelHandlerContext), new(TestParam2), channel.NewAttr(nil))
}

var (
	typeOfError          = reflect.TypeOf((*error)(nil)).Elem()
	typeOfString         = reflect.TypeOf(string(""))
	typeOfUInt32         = reflect.TypeOf(uint32(0))
	typeOfChannelContext = reflect.TypeOf((*channel.ChannelContext)(nil)).Elem()
	typeOfIAttr          = reflect.TypeOf((*channel.IAttr)(nil)).Elem()
)

type DefaultHandleChecker struct{}

func NewDefaultHandleChecker() (checker *DefaultHandleChecker) {
	checker = new(DefaultHandleChecker)
	return
}

// isHandlerMethod decide a method is suitable handler method
func (checker *DefaultHandleChecker) IsHandlerMethod(method reflect.Method) bool {
	mt := method.Type
	// Method must be exported.
	if method.PkgPath != "" {
		return false
	}
	// Method needs three ins: receiver, channel.ChannelHandlerContext, pointer, channel.IAttr
	if mt.NumIn() != 4 {
		return false
	}
	// Method needs one outs: error
	if mt.NumOut() != 1 {
		return false
	}
	if t1 := mt.In(1); t1 != typeOfChannelContext {
		return false
	}
	if mt.In(2).Kind() != reflect.Ptr {
		return false
	}
	if t1 := mt.In(3); t1 != typeOfIAttr {
		return false
	}
	if mt.Out(0) != typeOfError {
		return false
	}
	return true
}

// AdaptArgs create the params a handler method need
func (checker *DefaultHandleChecker) AdaptArgs(types []reflect.Type, params []interface{}, protoType uint32) []reflect.Value {
	args := []reflect.Value{reflect.ValueOf(params[0]), reflect.ValueOf(params[1]), reflect.ValueOf(params[2])}
	return args
}
