package cliService

import (
	"gitee.com/andyxt/gox/code/message"
	"gitee.com/andyxt/gox/service"
	"gitee.com/andyxt/gox/tools/cliTest/generic/mid"
	"gitee.com/andyxt/gox/tools/cliTest/pb/cli"
)

// RouteForHeartbeatResp is route for the handler LoginResp.
func (*Service) RouteForHeartbeatResp() (string, uint32, message.MessageType) {
	return "/HeartbeatResp", uint32(mid.HeartbeatResponse), message.MessageTypePB
}

// HeartbeatResp is the handler for AccountService.Login.
func (*Service) HeartbeatResp(request service.IServiceRequest, msg *cli.HeartbeatResponse) error {
	return nil
}
