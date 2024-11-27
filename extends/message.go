package extends

import "gitee.com/andyxt/gox/service"

func MsgID(request service.IServiceRequest) uint32 {
	reqContext := request.ReqContext()
	return uint32(reqContext.GetInt32("msgSeqId"))
}

func SetMsgID(request service.IServiceRequest, msgSeqID uint32) {
	request.ReqContext().Set("msgSeqId", int32(msgSeqID))
}
