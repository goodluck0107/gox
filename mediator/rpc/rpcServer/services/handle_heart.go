package services

import (
	"gitee.com/andyxt/gona/logger"
	"gitee.com/andyxt/gox/mediator/rpc/mid"
	"gitee.com/andyxt/gox/mediator/rpc/pb/rpc"
	"gitee.com/andyxt/gox/service"
)

// RouteForRPCHeartbeatRequest 心跳
func (*RpcService) RouteForRPCHeartbeatRequest() (string, uint32) {
	return "/RPCHeartbeatRequest", uint32(mid.RPCHeartbeatRequest)
}

func (*RpcService) RPCHeartbeatRequest(request service.IServiceRequest, msg *rpc.HeartbeatRequest) error {
	logger.Info("RPCHeartbeatRequest")
	return nil
}
