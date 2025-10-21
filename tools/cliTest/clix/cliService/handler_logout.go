package cliService

import (
	"github.com/goodluck0107/gox/code/message"
	"github.com/goodluck0107/gox/extends"
	"github.com/goodluck0107/gox/service"
	"github.com/goodluck0107/gox/tools/cliTest/generic/mid"
	"github.com/goodluck0107/gox/tools/cliTest/internal/pb/cli"
)

// RouteForLogoutResponse is route for the handler LoginResp.
func (*Service) RouteForLogoutResponse() (string, uint32, message.MessageType) {
	return "/LogoutResponse", uint32(mid.LogoutResponse), message.MessageTypePB
}

// LogoutResponse is the handler for AccountService.Login.
func (*Service) LogoutResponse(request service.IServiceRequest, msg *cli.LogoutResponse) error {
	chlContext := request.ChannelContext()
	extends.Logout(chlContext)
	return nil
}
