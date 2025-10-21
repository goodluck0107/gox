package services

import (
	"fmt"

	"github.com/goodluck0107/gox/code/message"
	"github.com/goodluck0107/gox/mediator/rpc/mid"
	"github.com/goodluck0107/gox/mediator/rpc/pb/rpc"
	"github.com/goodluck0107/gox/service"
)

// RouteForEchoRequest Echo
func (*RpcService) RouteForEchoRequest() (string, uint32, message.MessageType) {
	return "/EchoRequest", uint32(mid.EchoRequest), message.MessageTypePB
}

func (*RpcService) EchoRequest(request service.IServiceRequest, msg *rpc.EchoRequest) error {
	fmt.Printf("EchoRequest param1:%v param2:%v param3:%v \n", msg.Param1, msg.Param2, msg.Param3)
	return nil
}
