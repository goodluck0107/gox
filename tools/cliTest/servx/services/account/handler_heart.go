package account

import (
	"gitee.com/andyxt/gox/extends"
	"gitee.com/andyxt/gox/mediator/server"
	"gitee.com/andyxt/gox/service"
	"gitee.com/andyxt/gox/tools/cliTest/generic/mid"
	"gitee.com/andyxt/gox/tools/cliTest/pb/cli"
)

func (*AccountService) RouteForHeartbeatRequest() (string, uint32) {
	return "/HeartbeatRequest", uint32(mid.HeartbeatRequest)
}

func (*AccountService) HeartbeatRequest(request service.IServiceRequest, msg *cli.HeartbeatRequest) error {
	return server.Response(request.ChannelContext(), extends.MsgID(request), mid.HeartbeatResponse, &cli.HeartbeatResponse{})
}
