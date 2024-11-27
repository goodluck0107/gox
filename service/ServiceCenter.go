package service

import (
	"errors"
	"fmt"
	"reflect"

	"gitee.com/andyxt/gona/logger"
)

type serviceCenter struct {
	handlerChecker      IHandleChecker
	localHandlers       map[uint16]*handler
	localHandlersByPath map[string]*handler
}

func newServiceCenter() *serviceCenter {
	sc := new(serviceCenter)
	sc.localHandlers = make(map[uint16]*handler)
	sc.localHandlersByPath = make(map[string]*handler)
	return sc
}

func (sc *serviceCenter) Init(checker IHandleChecker) {
	sc.handlerChecker = checker
}

// Register register service handlers for message
func (sc *serviceCenter) Register(receiver interface{}) {
	if sc.handlerChecker == nil {
		sc.handlerChecker = newDefaultHandleChecker()
	}
	if err := sc.checkExtract(receiver); err != nil {
		logger.Error(fmt.Sprintf("service %v checkExtract error: %v", reflect.Indirect(reflect.ValueOf(receiver)).Type().Name(), err))
		return
	}
	sc.register(newService(receiver))
}

// 通过服务路径派发消息到服务方法
func (sc *serviceCenter) DispatchByPath(servicePath string, data []interface{}) error {
	if h, found := sc.localHandlersByPath[servicePath]; found {
		return sc.callServiceHandle(h, data)
	}
	return fmt.Errorf("no handler for route-path: %s", servicePath)
}

// 通过服务码派发消息到服务方法
func (sc *serviceCenter) DispatchByCode(serviceCode int32, data []interface{}) error {
	if h, found := sc.localHandlers[uint16(serviceCode)]; found {
		return sc.callServiceHandle(h, data)
	}
	return fmt.Errorf("no handler for route-code: %d", serviceCode)
}

// 通过服务路径获取服务码
func (sc *serviceCenter) Code(servicePath string) uint16 {
	if h, found := sc.localHandlersByPath[servicePath]; found {
		return uint16(h.Code)
	}
	return 0
}

// 通过服务码获取服务路径
func (sc *serviceCenter) Path(serviceCode uint16) string {
	if h, found := sc.localHandlers[serviceCode]; found {
		return h.Path
	}
	return ""
}

// HandlerType find the input-params type
func (sc *serviceCenter) HandlerType(servicePath string) []reflect.Type {
	if h, found := sc.localHandlersByPath[servicePath]; found {
		return h.Types
	}
	return nil
}

// 调用服务方法
func (sc *serviceCenter) callServiceHandle(h *handler, data []interface{}) error {
	handlerArgs := sc.handlerChecker.AdaptArgs(h.Types, data)
	if handlerArgs == nil {
		return errors.New("param for handler is invalid")
	}
	args := append([]reflect.Value{h.Receiver}, handlerArgs...)
	result := h.Method.Func.Call(args)
	if len(result) <= 0 || result[0].Interface() == nil {
		return nil
	}
	return result[0].Interface().(error)
}

// 检查服务是否满足注册条件
func (sc *serviceCenter) checkExtract(receiver interface{}) error {
	typeName := reflect.Indirect(reflect.ValueOf(receiver)).Type().Name()
	if typeName == "" {
		return errors.New("no service name for type " + reflect.TypeOf(receiver).String())
	}
	if !isExported(typeName) {
		return errors.New("type " + typeName + " is not exported")
	}
	return nil
}

// 往服务中心注册一个服务
func (sc *serviceCenter) register(s *service) {
	handlers := s.extractHandler(sc.handlerChecker)
	if len(handlers) > 0 {
		sc.registerHandlersWithCode(s.Name, handlers)
		sc.registerHandlersWithPath(s.Name, handlers)
		return
	}
	logger.Error(fmt.Sprintf("service %v has no handler", s.Name))
}

func (sc *serviceCenter) registerHandlersWithCode(serviceName string, handlers map[string]*handler) {
	for handlerName, handler := range handlers {
		if handler.Code <= 0 {
			logger.Error(fmt.Sprintf("Service %v has handler %s with no route-code", serviceName, handlerName))
			continue
		}
		// logger.Info(fmt.Sprintf("Service %v has handler %s for route-code %v", serviceName, handlerName, handler.Code))
		if _, ok := sc.localHandlers[uint16(handler.Code)]; ok {
			logger.Error(fmt.Sprintf("Service %v has duplicate handler %s for route-code: %v", serviceName, handlerName, handler.Code))
			panic(fmt.Errorf("service %v has duplicate handler %s for route-code: %v", serviceName, handlerName, handler.Code))
		}
		sc.localHandlers[uint16(handler.Code)] = handler
	}
}

func (sc *serviceCenter) registerHandlersWithPath(serviceName string, handlers map[string]*handler) {
	for handlerName, handler := range handlers {
		if handler.Path == "" {
			logger.Error(fmt.Sprintf("Service %v has handler %s with no route-path", serviceName, handlerName))
			continue
		}
		// logger.Info(fmt.Sprintf("Service %v has handler %s for route-path %v", serviceName, handlerName, handler.Path))
		if _, ok := sc.localHandlersByPath[handler.Path]; ok {
			logger.Error(fmt.Sprintf("Service %v has duplicate handler %s for route-path: %v", serviceName, handlerName, handler.Path))
			panic(fmt.Errorf("service %v has duplicate handler %s for route-path: %v", serviceName, handlerName, handler.Path))
		}
		sc.localHandlersByPath[handler.Path] = handler
	}
}
