package messageImpl

import (
	"gitee.com/andyxt/gox/mediator/server"
	"gitee.com/andyxt/gox/service"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var routeTypeClient Type = 0
var messageTypeClose Type = 0
var messageTypeProto Type = 1

func Response(chlCtx service.IChannelContext, msgSeqID uint32, msgMsgID uint16, v interface{}) error {
	return server.Response(chlCtx, msgSeqID, msgMsgID, NewMessage(routeTypeClient, messageTypeProto, 1, msgSeqID,
		msgMsgID, v.(protoreflect.ProtoMessage)))
}

func ResponseClose(chlCtx service.IChannelContext, msgSeqID uint32, msgMsgID uint16, v interface{}, desc string) error {
	return server.ResponseClose(chlCtx, msgSeqID, msgMsgID, NewMessage(routeTypeClient, messageTypeClose, 1, msgSeqID,
		msgMsgID, v.(protoreflect.ProtoMessage)), desc)
}

func Push(chlCtx service.IChannelContext, msgMsgID uint16, v interface{}) error {
	return server.Push(chlCtx, msgMsgID, NewMessage(routeTypeClient, messageTypeProto, 1, 0,
		msgMsgID, v.(protoreflect.ProtoMessage)))
}

func PushClose(chlCtx service.IChannelContext, msgMsgID uint16, v interface{}, desc string) error {
	return server.PushClose(chlCtx, msgMsgID, NewMessage(routeTypeClient, messageTypeClose, 1, 0,
		msgMsgID, v.(protoreflect.ProtoMessage)), desc)
}
