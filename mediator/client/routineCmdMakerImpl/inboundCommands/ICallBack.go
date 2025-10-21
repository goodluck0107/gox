package inboundCommands

import (
	"github.com/goodluck0107/gox/code/protocol"
	"github.com/goodluck0107/gox/service"
)

type ICallBack interface {
	ConnectSuccess(uID int64, currentChlCtx service.IChannelContext)
	ConnectFail(err error, params map[string]interface{})
	ConnectInactive(uID int64, currentChlCtx service.IChannelContext)
	MessageReceived(Ctx service.IChannelContext, Data protocol.Protocol)
}
