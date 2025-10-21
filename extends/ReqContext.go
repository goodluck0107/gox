package extends

import "github.com/goodluck0107/gox/service"

func SeqID(request service.IServiceRequest) uint32 {
	reqContext := request.ReqContext()
	return uint32(reqContext.GetInt32("msgSeqId"))
}

func SetSeqID(request service.IServiceRequest, msgSeqID uint32) {
	request.ReqContext().Set("msgSeqId", int32(msgSeqID))
}

func MsgID(request service.IServiceRequest) uint16 {
	reqContext := request.ReqContext()
	return uint16(reqContext.GetInt16("msgMsgId"))
}

func SetMsgID(request service.IServiceRequest, msgMsgID uint16) {
	request.ReqContext().Set("msgMsgId", int16(msgMsgID))
}
