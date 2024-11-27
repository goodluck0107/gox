package service

import "reflect"

var localServiceCenter *serviceCenter = newServiceCenter()

func Init(checker IHandleChecker) {
	localServiceCenter.Init(checker)
}

// Register register service handlers for message
func Register(receiver interface{}) {
	localServiceCenter.Register(receiver)
}

// Dispatch dispatch message to a service handler
func DispatchByPath(servicePath string, data ...interface{}) error {
	return localServiceCenter.DispatchByPath(servicePath, data)
}

// Dispatch dispatch message to a service handler
func DispatchByCode(serviceCode int32, data ...interface{}) error {
	return localServiceCenter.DispatchByCode(serviceCode, data)
}

// HandlerType find the input-params type
func HandlerType(servicePath string) []reflect.Type {
	return localServiceCenter.HandlerType(servicePath)
}

// 通过服务路径获取服务码
func Code(servicePath string) uint16 {
	return localServiceCenter.Code(servicePath)
}

// 通过服务码获取服务路径
func Path(serviceCode uint16) string {
	return localServiceCenter.Path(serviceCode)
}
