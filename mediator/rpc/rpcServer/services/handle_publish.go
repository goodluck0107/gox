package services

import (
	"fmt"
	"strings"
	"time"

	"gitee.com/andyxt/gona/logger"
	"gitee.com/andyxt/gona/utils/cast"
	"gitee.com/andyxt/gox/executor"
	"gitee.com/andyxt/gox/mediator/rpc/center"
	"gitee.com/andyxt/gox/mediator/rpc/mid"
	"gitee.com/andyxt/gox/mediator/rpc/pb/rpc"
	"gitee.com/andyxt/gox/messageImpl"
	"gitee.com/andyxt/gox/service"
)

// RouteForPublishRequest 发布订阅
func (*RpcService) RouteForPublishRequest() (string, uint32, uint32) {
	return "/PublishRequest", uint32(mid.PublishRequest), service.ProtoTypePB
}

func (*RpcService) PublishRequest(request service.IServiceRequest, msg *rpc.PublishRequest) error {
	logger.Info(fmt.Sprintf("PublishRequest Topic:%v", msg.Topic))
	executor.FireEvent(newPublishEvent(msg))
	return nil
}

type publishEvent struct {
	msg *rpc.PublishRequest
}

func newPublishEvent(msg *rpc.PublishRequest) *publishEvent {
	return &publishEvent{msg: msg}
}

func (recvEvent *publishEvent) QueueId() int64 {
	return time.Now().UnixNano()
}

func (recvEvent *publishEvent) Wait() (interface{}, bool) {
	return nil, true
}

func (recvEvent *publishEvent) Exec() {
	// pasrse Topic to messageID  TODO
	parts := strings.Split(recvEvent.msg.Topic, "_") // playerID_topic_msgID
	if len(parts) != 3 {
		logger.Error("parse topic error:", recvEvent.msg.Topic)
		return
	}
	playerIDPart := parts[0]
	playerID, castE := cast.ToInt64E(playerIDPart)
	if castE != nil {
		logger.Error("parse topic error:", recvEvent.msg.Topic)
		return
	}
	msgIDPart := parts[2]
	msgID, castE := cast.ToInt64E(msgIDPart)
	if castE != nil {
		logger.Error("parse topic error:", recvEvent.msg.Topic)
		return
	}
	center.TraverseDo(recvEvent.msg.Topic, func(ctx service.IChannelContext) {
		// executor.FireEvent(newSubscribeEvent(i1.UID(), msg))
		messageImpl.Push(ctx, mid.MessagePush, &rpc.MessagePush{
			Topic:    recvEvent.msg.Topic,
			PlayerID: playerID,
			MsgCode:  msgID,
			MsgData:  recvEvent.msg.MsgData,
		})
	})
}
