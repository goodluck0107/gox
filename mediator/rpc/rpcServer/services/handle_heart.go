package services

import (
	"gitee.com/andyxt/gona/logger"
	"gitee.com/andyxt/gox/mediator/rpc/mid"
	"gitee.com/andyxt/gox/mediator/rpc/pb/rpc"
	"gitee.com/andyxt/gox/messageImpl"
	"gitee.com/andyxt/gox/service"
)

// RouteForHeartbeatRequest 心跳
func (*RpcService) RouteForHeartbeatRequest() (string, uint32, uint32) {
	return "/HeartbeatRequest", uint32(mid.HeartbeatRequest), service.ProtoTypePB
}

func (*RpcService) HeartbeatRequest(request service.IServiceRequest, msg *rpc.HeartbeatRequest) error {
	logger.Info("HeartbeatRequest")
	messageImpl.Push(request.ChannelContext(), mid.HeartbeatResponse, &rpc.HeartbeatResponse{})
	return nil
}
