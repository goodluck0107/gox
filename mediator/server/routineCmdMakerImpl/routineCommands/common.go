package routineCommands

import (
	"fmt"

	"gitee.com/andyxt/gox/message"

	"gitee.com/andyxt/gona/logger"
	"gitee.com/andyxt/gox/extends"

	"gitee.com/andyxt/gox/service"
)

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
