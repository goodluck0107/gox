package account

import (
	"gitee.com/andyxt/gox/extends"
	"gitee.com/andyxt/gox/messageImpl"
	"gitee.com/andyxt/gox/service"
	"gitee.com/andyxt/gox/tools/cliTest/generic/mid"
	"gitee.com/andyxt/gox/tools/cliTest/pb/cli"
)

func (*AccountService) RouteForHeartbeatRequest() (string, uint32, uint32) {
	return "/HeartbeatRequest", uint32(mid.HeartbeatRequest), service.ProtoTypePB
}

func (*AccountService) HeartbeatRequest(request service.IServiceRequest, msg *cli.HeartbeatRequest) error {
	return messageImpl.Response(request.ChannelContext(), extends.SeqID(request), mid.HeartbeatResponse, &cli.HeartbeatResponse{})
}
