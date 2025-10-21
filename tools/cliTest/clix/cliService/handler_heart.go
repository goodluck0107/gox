package cliService

import (
	"github.com/goodluck0107/gox/code/message"
	"github.com/goodluck0107/gox/service"
	"github.com/goodluck0107/gox/tools/cliTest/generic/mid"
	"github.com/goodluck0107/gox/tools/cliTest/internal/pb/cli"
)

// RouteForHeartbeatResp is route for the handler LoginResp.
func (*Service) RouteForHeartbeatResp() (string, uint32, message.MessageType) {
	return "/HeartbeatResp", uint32(mid.HeartbeatResponse), message.MessageTypePB
}

// HeartbeatResp is the handler for AccountService.Login.
func (*Service) HeartbeatResp(request service.IServiceRequest, msg *cli.HeartbeatResponse) error {
	return nil
}
