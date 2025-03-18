package cliService

import (
	"gitee.com/andyxt/gox/code/message"
	"gitee.com/andyxt/gox/extends"
	"gitee.com/andyxt/gox/service"
	"gitee.com/andyxt/gox/tools/cliTest/generic/mid"
	"gitee.com/andyxt/gox/tools/cliTest/pb/cli"
)

// RouteForLoginConflict is route for the handler LoginResp.
func (*Service) RouteForLoginConflict() (string, uint32, message.MessageType) {
	return "/LoginConflict", uint32(mid.LoginConflictPush), message.MessageTypePB
}

// LoginConflict is the handler for AccountService.Login.
func (*Service) LoginConflict(request service.IServiceRequest, msg *cli.LoginConflictPush) error {
	chlContext := request.ChannelContext()
	extends.Conflict(chlContext)
	return nil
}
