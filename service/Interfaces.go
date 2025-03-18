package service

import (
	"reflect"

	"gitee.com/andyxt/gona/boot/channel"
)

type IHandleChecker interface {
	// IsHandlerMethod decide a method is suitable handler method
	IsHandlerMethod(method reflect.Method) bool
	// AdaptArgs create the params a handler method need
	AdaptArgs(types []reflect.Type, data []interface{}, protoType MessageType) ([]reflect.Value, error)
}

type IRouteMapper interface {
	GetCodeForPath(reqPath string) int32
	GetPathForCode(code int32) string
}
type IReqContext IAttr
type IChannelContext interface {
	ContextAttr() channel.IAttr

	ID() string

	RemoteAddr() string

	/*发起写事件，消息将被送往管道处理*/
	Write(data interface{})

	/*发起关闭事件，消息将被送往管道处理*/
	Close()
}

type IServiceRequest interface {
	ChannelContext() IChannelContext
	ReqContext() IReqContext
}
type IServiceResponse interface {
	Write(interface{}, ISerializer)
}

type ISerializer interface {
	Serialize(message interface{}) ([]byte, error)
}
type IAttr interface {
	Get(key string) (value interface{})
	GetBool(key string) bool
	GetInt8(key string) (value int8)
	GetInt16(key string) (value int16)
	GetInt32(key string) (value int32)
	GetInt64(key string) (value int64)
	GetInt(key string) (value int)
	GetString(key string) (value string)
	Set(key string, value interface{})
	CopyToMap() map[string]interface{}
	CopyFromMap(newAttr map[string]interface{})
	Copy(newAttr IAttr)
}
