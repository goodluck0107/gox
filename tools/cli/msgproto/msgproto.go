package msgproto

import (
	"encoding/json"
	"reflect"

	"google.golang.org/protobuf/proto"
)

// AdaptArgs create the params a handler method need
func AdaptArgsFromProto(types []reflect.Type, param interface{}) reflect.Value {
	data := reflect.New(types[1].Elem()).Interface()
	pb, ok := data.(proto.Message)
	if !ok {
		return reflect.Zero(nil)
	}
	err := proto.Unmarshal(param.([]byte), pb)
	if err != nil {
		return reflect.Zero(nil)
	}
	return reflect.ValueOf(data)
}

// AdaptArgs create the params a handler method need
func AdaptArgsFromJson(types []reflect.Type, param interface{}) reflect.Value {
	data := reflect.New(types[1].Elem()).Interface()
	err := json.Unmarshal(param.([]byte), data)
	if err != nil {
		return reflect.Zero(nil)
	}
	return reflect.ValueOf(data)
}
