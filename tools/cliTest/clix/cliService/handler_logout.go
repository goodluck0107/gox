package cliService

import (
	"gitee.com/andyxt/gox/extends"
	"gitee.com/andyxt/gox/service"
	"gitee.com/andyxt/gox/tools/cliTest/generic/mid"
	"gitee.com/andyxt/gox/tools/cliTest/pb/cli"
)

// RouteForLogoutResponse is route for the handler LoginResp.
func (*Service) RouteForLogoutResponse() (string, uint32, service.MessageType) {
	return "/LogoutResponse", uint32(mid.LogoutResponse), service.ProtoTypePB
}

// LogoutResponse is the handler for AccountService.Login.
func (*Service) LogoutResponse(request service.IServiceRequest, msg *cli.LogoutResponse) error {
	chlContext := request.ChannelContext()
	extends.Logout(chlContext)
	return nil
}
