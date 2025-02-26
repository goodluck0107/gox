package services

import (
	"fmt"
	"strings"

	"gitee.com/andyxt/gona/logger"
	"gitee.com/andyxt/gona/utils/cast"
	"gitee.com/andyxt/gox/executor"
	"gitee.com/andyxt/gox/mediator/rpc/mid"
	"gitee.com/andyxt/gox/mediator/rpc/pb/rpc"
	"gitee.com/andyxt/gox/mediator/rpc/rpcServer/center"
	"gitee.com/andyxt/gox/messageImpl"
	"gitee.com/andyxt/gox/service"
)

const (
	publishPath = "/PublishRequest"
)

// RouteForPublishRequest 发布订阅
func (*RpcService) RouteForPublishRequest() (string, uint32, uint32) {
	return publishPath, uint32(mid.PublishRequest), service.ProtoTypePB
}

func (*RpcService) PublishRequest(request service.IServiceRequest, msg *rpc.PublishRequest) error {
	logger.Info(fmt.Sprintf("PublishRequest Topic:%v", msg.Topic))
	executor.FireEvent(newPublishEvent(request.ChannelContext(), msg))
	return nil
}

type publishEvent struct {
	ctx service.IChannelContext
	msg *rpc.PublishRequest
}

func newPublishEvent(ctx service.IChannelContext, msg *rpc.PublishRequest) *publishEvent {
	return &publishEvent{ctx: ctx, msg: msg}
}

func (recvEvent *publishEvent) QueueId() int64 {
	return stringToInt64(recvEvent.ctx.ID())
}

func (recvEvent *publishEvent) Wait() (interface{}, bool) {
	return nil, true
}

func (recvEvent *publishEvent) Exec() {
	parts := strings.SplitN(recvEvent.msg.Topic, "_", 3) // playerID_topic_msgID
	if len(parts) != 3 {
		logger.Error("parse topic error:", recvEvent.msg.Topic)
		return
	}
	playerIDPart, msgIDPart := parts[0], parts[2]
	playerID, castE := cast.ToInt64E(playerIDPart)
	if castE != nil {
		logger.Error(fmt.Printf("failed to parse player ID from topic: %v", recvEvent.msg.Topic))
		return
	}
	msgID, castE := cast.ToInt64E(msgIDPart)
	if castE != nil {
		logger.Error(fmt.Printf("failed to parse message ID from topic: %v", recvEvent.msg.Topic))
		return
	}
	center.TraverseDo(recvEvent.msg.Topic, func(ctx service.IChannelContext) {
		err := messageImpl.Push(ctx, mid.MessagePush, &rpc.MessagePush{
			Topic:    recvEvent.msg.Topic,
			PlayerID: playerID,
			MsgCode:  msgID,
			MsgData:  recvEvent.msg.MsgData,
		})
		if err != nil {
			logger.Error(fmt.Printf("failed to push MessagePush: %v", err))
		}
	})
}
