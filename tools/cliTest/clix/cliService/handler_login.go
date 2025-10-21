package cliService

import (
	"github.com/goodluck0107/gox/code/message"
	"github.com/goodluck0107/gox/extends"
	"github.com/goodluck0107/gox/mediator/client/clientkey"
	"github.com/goodluck0107/gox/service"
	"github.com/goodluck0107/gox/tools/cliTest/generic/mid"
	"github.com/goodluck0107/gox/tools/cliTest/internal/pb/cli"
)

// RouteForLoginResp is route for the handler LoginResp.
func (*Service) RouteForLoginResp() (string, uint32, message.MessageType) {
	return "/LoginResp", uint32(mid.LoginResponse), message.MessageTypePB
}

// LoginResp is the handler for AccountService.Login.
func (*Service) LoginResp(request service.IServiceRequest, msg *cli.LoginResponse) error {
	chlContext := request.ChannelContext()
	uID := chlContext.ContextAttr().GetInt64(clientkey.KeyFireUser)
	extends.PutInUserInfo(chlContext, uID, 0)
	return nil
}
