package commands

import (
	"fmt"

	"gitee.com/andyxt/gona/boot"
	"gitee.com/andyxt/gona/logger"
	"gitee.com/andyxt/gox/extends"
	"gitee.com/andyxt/gox/handler/protocol"
	"gitee.com/andyxt/gox/mediator/client"
	"gitee.com/andyxt/gox/mediator/client/clientkey"
	"gitee.com/andyxt/gox/mediator/client/routineCmdMakerImpl/inboundCommands"
	"gitee.com/andyxt/gox/message"
	"gitee.com/andyxt/gox/service"
	"gitee.com/andyxt/gox/tools/cli/msgproto"
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
	ip := params[boot.KeyIP].(string)
	port := params[boot.KeyPort].(int)
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
	ip := currentChlCtx.ContextAttr().GetString(boot.KeyIP)
	port := currentChlCtx.ContextAttr().GetInt(boot.KeyPort)
	clientFacade.Connect(ip, port, uID)
}

// MessageReceived called on executor routine
func (cb *CallBack) MessageReceived(Ctx service.IChannelContext, Data protocol.Protocol) {
	fmt.Println("CallBack.MessageReceived")
	msg := Data.(*message.Message)
	msgPath := service.Path(msg.MsgID)
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

func callService(chlContext service.IChannelContext, msg *message.Message) error {
	request := service.NewSessionRequest(chlContext, service.NewAttr(nil))
	extends.SetMsgID(request, msg.SeqID)
	serviceCode := int32(msg.MsgID)
	serviceErr := service.DispatchByCode(serviceCode, request, msg.Data)
	if serviceErr != nil {
		logger.Error(fmt.Sprintf("chlCtx %v callService %v error %v ", extends.ChannelContextToString(chlContext), serviceCode, serviceErr))
	}
	return serviceErr
}
