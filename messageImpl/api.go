package messageImpl

import (
	"github.com/goodluck0107/gox/mediator/server"
	"github.com/goodluck0107/gox/service"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var routeTypeClient Type = 0
var messageTypeClose Type = 0
var messageTypeProto Type = 1
var messageVersion Type = 1

func ResponseMessage(msgSeqID uint32, msgMsgID uint16, v interface{}) *Message {
	return NewMessage(routeTypeClient, messageTypeProto, messageVersion, msgSeqID,
		msgMsgID, v.(protoreflect.ProtoMessage))
}

func Response(chlCtx service.IChannelContext, msgSeqID uint32, msgMsgID uint16, v interface{}) error {
	return server.Response(chlCtx, NewMessage(routeTypeClient, messageTypeProto, messageVersion, msgSeqID,
		msgMsgID, v.(protoreflect.ProtoMessage)))
}

func ResponseClose(chlCtx service.IChannelContext, msgSeqID uint32, msgMsgID uint16, v interface{}, desc string) error {
	return server.ResponseClose(chlCtx, NewMessage(routeTypeClient, messageTypeClose, messageVersion, msgSeqID,
		msgMsgID, v.(protoreflect.ProtoMessage)), desc)
}

func Push(chlCtx service.IChannelContext, msgMsgID uint16, v interface{}) error {
	return server.Push(chlCtx, NewMessage(routeTypeClient, messageTypeProto, messageVersion, 0,
		msgMsgID, v.(protoreflect.ProtoMessage)))
}

func PushClose(chlCtx service.IChannelContext, msgMsgID uint16, v interface{}, desc string) error {
	return server.PushClose(chlCtx, NewMessage(routeTypeClient, messageTypeClose, messageVersion, 0,
		msgMsgID, v.(protoreflect.ProtoMessage)), desc)
}
