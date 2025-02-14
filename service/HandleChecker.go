package service

import (
	"encoding/json"
	"errors"
	"reflect"

	"gitee.com/andyxt/gox/handler/message"
	"google.golang.org/protobuf/proto"
)

const (
	ProtoTypePB   uint32 = 0 // ProtoBuffer
	ProtoTypeBN   uint32 = 1 // Binary
	ProtoTypeJson uint32 = 2 // Json
)

var (
	typeOfError    = reflect.TypeOf((*error)(nil)).Elem()
	typeOfUInt32   = reflect.TypeOf(uint32(0))
	typeOfString   = reflect.TypeOf(string(""))
	typeOfIRequest = reflect.TypeOf((*IServiceRequest)(nil)).Elem()
)

const (
	routeMethodPre string = "RouteFor"
)

type defaultHandleChecker struct{}

func newDefaultHandleChecker() (checker *defaultHandleChecker) {
	checker = new(defaultHandleChecker)
	return
}

// isHandlerMethod decide a method is suitable handler method
func (checker *defaultHandleChecker) IsHandlerMethod(method reflect.Method) bool {
	mt := method.Type
	// Method must be exported.
	if method.PkgPath != "" {
		return false
	}
	// Method needs three ins: receiver, channel.ChannelHandlerContext, pointer, channel.IAttr
	if mt.NumIn() != 3 {
		return false
	}
	// Method needs one outs: error
	if mt.NumOut() != 1 {
		return false
	}
	if t1 := mt.In(1); t1 != typeOfIRequest {
		return false
	}
	if mt.In(2).Kind() != reflect.Ptr {
		return false
	}
	if mt.Out(0) != typeOfError {
		return false
	}
	return true
}

// AdaptArgs create the params a handler method need
func (checker *defaultHandleChecker) AdaptArgs(types []reflect.Type, params []interface{}, protoType uint32) ([]reflect.Value, error) {
	data := reflect.New(types[1].Elem()).Interface()
	b := params[1].([]byte)
	if protoType == ProtoTypePB { // ProtoBuffer消息
		pm, ok := data.(proto.Message)
		if !ok {
			return nil, errors.New("param is not proto.Message while protoType is ProtoTypePB")
		}
		err := proto.Unmarshal(b, pm)
		if err != nil {
			return nil, errors.New("param can not unmarshal to proto.Message while protoType is ProtoTypePB")
		}
	} else if protoType == ProtoTypeBN { // 二进制消息
		bm, ok := data.(message.IMessage)
		if !ok {
			return nil, errors.New("param is not message.IMessage while protoType is ProtoTypeBN")
		}
		err := message.Unmarshal(b, bm)
		if err != nil {
			return nil, errors.New("param can not unmarshal to message.IMessage while protoType is ProtoTypeBN")
		}
	} else if protoType == ProtoTypeJson { // json消息
		err := json.Unmarshal(b, data)
		if err != nil {
			return nil, errors.New("param can not unmarshal to json while protoType is ProtoTypeJson")
		}
	}
	args := []reflect.Value{reflect.ValueOf(params[0]), reflect.ValueOf(data)}
	return args, nil
}

// fmt.Println("err", err)
// fmt.Println("result", result[0].Interface().(*cli.LoginRequest))
// fmt.Println("result1", result[0].Interface().(protoiface.MessageV1))
// ProtoData := result[0].Interface().(protoiface.MessageV1)
// byteArr, err1 := proto.Marshal(ProtoData)
// if err1 != nil {
// 	panic(err)
// }
// reflect.New(h.Types[1].Elem())}
