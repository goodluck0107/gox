package service

import (
	"reflect"

	"google.golang.org/protobuf/proto"
)

var (
	typeOfError    = reflect.TypeOf((*error)(nil)).Elem()
	typeOfUInt32   = reflect.TypeOf(uint32(0))
	typeOfString   = reflect.TypeOf(string(""))
	typeOfIRequest = reflect.TypeOf((*IServiceRequest)(nil)).Elem()
)

const (
	PreStringForRouteMethod string = "RouteFor"
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
func (checker *defaultHandleChecker) AdaptArgs(types []reflect.Type, params []interface{}) []reflect.Value {
	data := reflect.New(types[1].Elem()).Interface()
	pb, ok := data.(proto.Message)
	if !ok {
		return nil
	}
	err := proto.Unmarshal(params[1].([]byte), pb)
	if err != nil {
		return nil
	}
	args := []reflect.Value{reflect.ValueOf(params[0]), reflect.ValueOf(data)}
	return args
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
