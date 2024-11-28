package cliService

import (
	"gitee.com/andyxt/gox/extends"
	"gitee.com/andyxt/gox/mediator/client/clientkey"
	"gitee.com/andyxt/gox/service"
	"gitee.com/andyxt/gox/tools/cliTest/generic/mid"
	"gitee.com/andyxt/gox/tools/cliTest/pb/cli"
)

// RouteForLoginResp is route for the handler LoginResp.
func (*Service) RouteForLoginResp() (string, uint32) {
	return "/LoginResp", uint32(mid.LoginResponse)
}

// LoginResp is the handler for AccountService.Login.
func (*Service) LoginResp(request service.IServiceRequest, msg *cli.LoginResponse) error {
	chlContext := request.ChannelContext()
	uID := chlContext.ContextAttr().GetInt64(clientkey.KeyFireUser)
	extends.PutInUserInfo(chlContext, uID, 0)
	return nil
}
