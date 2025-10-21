package services

import (
	"github.com/goodluck0107/gox/code/message"
	"github.com/goodluck0107/gox/mediator/rpc/mid"
	"github.com/goodluck0107/gox/mediator/rpc/pb/rpc"
	"github.com/goodluck0107/gox/messageImpl"
	"github.com/goodluck0107/gox/service"
)

const (
	heartbeatPath = "/HeartbeatRequest"
)

// RouteForHeartbeatRequest 心跳
func (*RpcService) RouteForHeartbeatRequest() (string, uint32, message.MessageType) {
	return heartbeatPath, uint32(mid.HeartbeatRequest), message.MessageTypePB
}

func (*RpcService) HeartbeatRequest(request service.IServiceRequest, msg *rpc.HeartbeatRequest) error {
	messageImpl.Push(request.ChannelContext(), mid.HeartbeatResponse, &rpc.HeartbeatResponse{})
	return nil
}
