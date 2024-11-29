package services

import (
	"fmt"

	"gitee.com/andyxt/gox/mediator/rpc/mid"
	"gitee.com/andyxt/gox/mediator/rpc/pb/rpc"
	"gitee.com/andyxt/gox/service"
)

// RouteForEchoRequest Echo
func (*RpcService) RouteForEchoRequest() (string, uint32) {
	return "/EchoRequest", uint32(mid.EchoRequest)
}

func (*RpcService) EchoRequest(request service.IServiceRequest, msg *rpc.EchoRequest) error {
	fmt.Printf("EchoRequest param1:%v param2:%v param3:%v \n", msg.Param1, msg.Param2, msg.Param3)
	return nil
}
