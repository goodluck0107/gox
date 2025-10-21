package account

import (
	"github.com/goodluck0107/gox/code/message"
	"github.com/goodluck0107/gox/extends"
	"github.com/goodluck0107/gox/messageImpl"
	"github.com/goodluck0107/gox/service"
	"github.com/goodluck0107/gox/tools/cliTest/generic/mid"
	"github.com/goodluck0107/gox/tools/cliTest/internal/pb/cli"
)

func (*AccountService) RouteForHeartbeatRequest() (string, uint32, message.MessageType) {
	return "/HeartbeatRequest", uint32(mid.HeartbeatRequest), message.MessageTypePB
}

func (*AccountService) HeartbeatRequest(request service.IServiceRequest, msg *cli.HeartbeatRequest) error {
	return messageImpl.Response(request.ChannelContext(), extends.SeqID(request), mid.HeartbeatResponse, &cli.HeartbeatResponse{})
}
