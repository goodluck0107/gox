package service

import (
	"reflect"

	"gitee.com/andyxt/gox/code/message"
)

type handler struct {
	Receiver    reflect.Value       // receiver of method
	Method      reflect.Method      // method stub
	Types       []reflect.Type      // low-level type of method
	Code        uint32              // Route compressed code
	Path        string              // Route service path
	messageType message.MessageType // Route compressed code
}
