package commands

import (
	"fmt"

	"github.com/goodluck0107/gona/boot/bootc"
	"github.com/goodluck0107/gox/code/protocol"
	"github.com/goodluck0107/gox/extends"
	"github.com/goodluck0107/gox/internal/logger"
	"github.com/goodluck0107/gox/mediator/client"
	"github.com/goodluck0107/gox/mediator/client/clientkey"
	"github.com/goodluck0107/gox/mediator/client/routineCmdMakerImpl/inboundCommands"
	"github.com/goodluck0107/gox/messageImpl"
	"github.com/goodluck0107/gox/service"
	"github.com/goodluck0107/gox/tools/cli/msgproto"
)

var clientFacade *client.ClientFacade

func Init() {
	clientFacade = client.BootClient(NewCallBack())
}

var UID int64
var CurrentChlCtx service.IChannelContext

type CallBack struct{}

func NewCallBack() inboundCommands.ICallBack {
	instance := new(CallBack)
	return instance
}

// ConnectSuccess called on executor routine
func (cb *CallBack) ConnectSuccess(uID int64, currentChlCtx service.IChannelContext) {
	fmt.Println("CallBack.ConnectSuccess")
	UID = uID
	CurrentChlCtx = currentChlCtx
}

// ConnectFail called on executor routine
func (cb *CallBack) ConnectFail(err error, params map[string]interface{}) {
	fmt.Println("CallBack.ConnectFail")
	ip := params[bootc.KeyIP].(string)
	port := params[bootc.KeyPort].(int)
	uID := params[clientkey.KeyFireUser].(int64)
	clientFacade.Connect(ip, port, uID)
}

// ConnectInactive  called on executor routine
func (cb *CallBack) ConnectInactive(uID int64, currentChlCtx service.IChannelContext) {
	fmt.Println("CallBack.ConnectInactive")
	if extends.IsConflict(currentChlCtx) {
		logger.Debug("CallBack.ConnectInactive 连接已经被其他连接挤下线，不需要重连", uID)
		return
	}
	if extends.IsClose(currentChlCtx) {
		logger.Debug("CallBack.ConnectInactive 连接已经被主动关闭，不需要重连", uID)
		return
	}
	if extends.IsLogout(currentChlCtx) {
		logger.Debug("CallBack.ConnectInactive 连接已经被主动登出，不需要重连", uID)
		return
	}
	if extends.IsSystemKick(currentChlCtx) {
		logger.Debug("CallBack.ConnectInactive 连接已经被踢出，不需要重连", uID)
		return
	}
	// broken: default reconnect
	logger.Debug("CallBack.ConnectInactive 连接中断，需要重连", uID)
	ip := currentChlCtx.ContextAttr().GetString(bootc.KeyIP)
	port := currentChlCtx.ContextAttr().GetInt(bootc.KeyPort)
	clientFacade.Connect(ip, port, uID)
}

// MessageReceived called on executor routine
func (cb *CallBack) MessageReceived(Ctx service.IChannelContext, Data protocol.Protocol) {
	fmt.Println("CallBack.MessageReceived")
	msg := Data.(*messageImpl.Message)
	msgPath := service.Path(uint32(msg.MsgID))
	if msgPath == "" {
		fmt.Println(`no msgPath for response`, msg.MsgID)
		return
	}
	handlerTypes := service.HandlerType(msgPath)
	if handlerTypes == nil {
		fmt.Println(`no handler for msgPath`, msgPath)
		return
	}
	adapterResult := msgproto.AdaptArgsFromProto(handlerTypes, msg.Data)
	fmt.Println("server->", "SeqID:", msg.SeqID, " MsgID:", msg.MsgID, " MsgPath:", msgPath, ":", adapterResult)
	callService(Ctx, msg)
}

func callService(chlContext service.IChannelContext, msg *messageImpl.Message) error {
	request := service.NewSessionRequest(chlContext, service.NewAttr(nil))
	extends.SetSeqID(request, msg.SeqID)
	serviceCode := uint32(msg.MsgID)
	serviceErr := service.DispatchByCode(serviceCode, request, msg.Data)
	if serviceErr != nil {
		logger.Error(fmt.Sprintf("chlCtx %v callService %v error %v ", extends.ChannelContextToString(chlContext), serviceCode, serviceErr))
	}
	return serviceErr
}
