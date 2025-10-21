package cliService

import (
	"github.com/goodluck0107/gox/code/message"
	"github.com/goodluck0107/gox/extends"
	"github.com/goodluck0107/gox/service"
	"github.com/goodluck0107/gox/tools/cliTest/generic/mid"
	"github.com/goodluck0107/gox/tools/cliTest/internal/pb/cli"
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
