package service

import "reflect"

type handler struct {
	Receiver reflect.Value  // receiver of method
	Method   reflect.Method // method stub
	Types    []reflect.Type // low-level type of method
	Code     uint32         // Route compressed code
	Path     string         // Route service path
}
