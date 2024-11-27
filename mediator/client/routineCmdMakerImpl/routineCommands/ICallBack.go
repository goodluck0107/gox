package routineCommands

import (
	"gitee.com/andyxt/gox/handler/protocol"
	"gitee.com/andyxt/gox/service"
)

type ICallBack interface {
	ConnectSuccess(uID int64, currentChlCtx service.IChannelContext)
	MessageReceived(Ctx service.IChannelContext, Data protocol.IProtocol)
}
