package commands

import (
	"fmt"

	"gitee.com/andyxt/gona/logger"
	"gitee.com/andyxt/gox/extends"
	"gitee.com/andyxt/gox/mediator/client/routineCmdMakerImpl/routineCommands"
	"gitee.com/andyxt/gox/message"
	"gitee.com/andyxt/gox/tools/cli/msgproto"

	"gitee.com/andyxt/gox/handler/protocol"

	"gitee.com/andyxt/gox/service"
)

var UID int64
var CurrentChlCtx service.IChannelContext

type CallBack struct{}

func NewCallBack() routineCommands.ICallBack {
	instance := new(CallBack)
	return instance
}

func (cb *CallBack) ConnectSuccess(uID int64, currentChlCtx service.IChannelContext) {
	fmt.Println("ConnectSuccess")
	UID = uID
	CurrentChlCtx = currentChlCtx
}

func (cb *CallBack) MessageReceived(Ctx service.IChannelContext, Data protocol.IProtocol) {
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
