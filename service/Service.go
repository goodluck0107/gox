package service

import (
	"reflect"
	"strings"
)

type service struct {
	Name     string        // name of service
	Type     reflect.Type  // type of the receiver
	Receiver reflect.Value // receiver of methods for the service
}

func newService(receiver interface{}) (s *service) {
	s = &service{
		Name:     reflect.Indirect(reflect.ValueOf(receiver)).Type().Name(),
		Type:     reflect.TypeOf(receiver),
		Receiver: reflect.ValueOf(receiver),
	}
	return
}

// extractHandler extract the set of methods from the
// receiver value which satisfy the following conditions:
// - exported method of exported type
// - three arguments, both of exported type
// - the first argument is channel.ChannelHandlerContext
// - the second argument is a pointer
// - the third argument is channel.IAttr
func (s *service) extractHandler(handlerChecker IHandleChecker) map[string]*handler {
	handlers := s.suitableHandlerMethods(s.Type, handlerChecker)
	for _, handler := range handlers {
		handler.Receiver = s.Receiver
		s.matchHandlerRoute(s.Type, handler)
	}
	return handlers
}

// suitableMethods returns suitable methods of typ
func (s *service) suitableHandlerMethods(typ reflect.Type, handlerChecker IHandleChecker) map[string]*handler {
	methods := make(map[string]*handler)
	for m := 0; m < typ.NumMethod(); m++ {
		method := typ.Method(m)
		mt := method.Type
		mn := method.Name
		if handlerChecker.IsHandlerMethod(method) {
			numInCount := mt.NumIn() - 1
			types := make([]reflect.Type, numInCount)
			for i := 0; i < numInCount; i++ {
				types[i] = mt.In(i + 1)
			}
			methods[mn] = &handler{
				Method: method,
				Types:  types,
			}
		}
	}
	return methods
}

// matchHandlerRoute find the route message code for handler
func (s *service) matchHandlerRoute(typ reflect.Type, h *handler) {
	for m := 0; m < typ.NumMethod(); m++ {
		method := typ.Method(m)
		if !s.IsRouteMethod(method) || !s.IsRouteMethodForHandlerMethod(method, h.Method) {
			continue
		}
		args := []reflect.Value{h.Receiver} //, reflect.New(h.Types[1].Elem())}
		result := method.Func.Call(args)
		h.Path = string(result[0].String())
		h.Code = uint32(result[1].Uint())
		return
	}
}

// IsRouteMethod decide a method is suitable handler method
func (s *service) IsRouteMethod(method reflect.Method) bool {
	if strings.Index(method.Name, PreStringForRouteMethod) != 0 {
		return false
	}
	mt := method.Type
	// Method must be exported.
	if method.PkgPath != "" {
		return false
	}
	// Method needs three ins: receiver, *Session, []byte or pointer.
	if mt.NumIn() != 1 {
		return false
	}
	// Method needs one outs: error
	if mt.NumOut() != 2 {
		return false
	}
	if mt.Out(0) != typeOfString {
		return false
	}
	if mt.Out(1) != typeOfUInt32 {
		return false
	}
	return true
}

// IsRouteMethodForHandlerMethod decide a routeMethod is suitable for handlerMethod
func (s *service) IsRouteMethodForHandlerMethod(routeMethod reflect.Method, handlerMethod reflect.Method) bool {
	return routeMethod.Name == PreStringForRouteMethod+handlerMethod.Name
}
