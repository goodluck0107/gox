package rpcClient

import (
	"crypto/md5"
	"fmt"
	"time"

	"github.com/goodluck0107/gox/mediator/client/clientkey"
	"github.com/goodluck0107/gox/mediator/client/routineCmdMakerImpl/inboundCommands"
	"github.com/goodluck0107/gox/messageImpl"
	"github.com/goodluck0107/gox/service"

	"github.com/goodluck0107/gox/mediator/rpc/mid"

	"github.com/goodluck0107/gox/mediator/rpc/pb/rpc"

	"github.com/goodluck0107/gox/mediator/client"

	"github.com/goodluck0107/gona/boot/bootc"
	"github.com/goodluck0107/gox/code/protocol"
	"github.com/goodluck0107/gox/extends"
	"github.com/goodluck0107/gox/internal/logger"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var clientFacade *client.ClientFacade
var clientNodeID int64 // client节点ID

// Start 连接master
func Start(nodeName string, ip string, port int64) {
	service.Register(NewService())
	clientNodeID = Mod(nodeName)
	clientFacade = client.BootClient(newCallBack())
	clientFacade.Connect(ip, int(port), clientNodeID)
}

// Subscribe 调用master发起订阅
func Subscribe(topic string) {
	sendMessage(mid.SubscribeRequest, &rpc.SubscribeRequest{
		Topic: topic,
	})
}

// Subscribe 调用master发起订阅
func Unsubscribe(topic string) {
	sendMessage(mid.UnsubscribeRequest, &rpc.UnsubscribeRequest{
		Topic: topic,
	})
}

// Publish 调用master发布订阅
func Publish(topic string, protoData protoreflect.ProtoMessage) {
	marshalV, marshalE := proto.Marshal(protoData)
	if marshalE != nil {
		logger.Error(fmt.Sprintf("RPC.Publish err:%v", marshalE))
		return
	}
	sendMessage(mid.PublishRequest, &rpc.PublishRequest{
		Topic:   topic,
		MsgData: marshalV,
	})
}

type callBack struct{}

func newCallBack() inboundCommands.ICallBack {
	instance := new(callBack)
	return instance
}

func (cb *callBack) ConnectSuccess(uID int64, currentChlCtx service.IChannelContext) {
	logger.Info(fmt.Sprintf("RPC.CallBack.ConnectSuccess uID:%v", uID))
	chlCtxReference.Store(uID, currentChlCtx)
	heart()
}

// ConnectFail called on executor routine
func (cb *callBack) ConnectFail(err error, params map[string]interface{}) {
	ip := params[bootc.KeyIP].(string)
	port := params[bootc.KeyPort].(int)
	uID := params[clientkey.KeyFireUser].(int64)
	logger.Info(fmt.Sprintf("RPC.CallBack.ConnectFail uID:%v", uID))
	clientFacade.Connect(ip, port, uID)
}

// ConnectInactive  called on executor routine
func (cb *callBack) ConnectInactive(uID int64, currentChlCtx service.IChannelContext) {
	logger.Info(fmt.Sprintf("RPC.CallBack.ConnectInactive uID:%v", uID))
	if extends.IsConflict(currentChlCtx) {
		logger.Debug("RPC.CallBack.ConnectInactive 连接已经被其他连接挤下线，不需要重连", uID)
		return
	}
	if extends.IsClose(currentChlCtx) {
		logger.Debug("RPC.CallBack.ConnectInactive 连接已经被主动关闭，不需要重连", uID)
		return
	}
	if extends.IsLogout(currentChlCtx) {
		logger.Debug("RPC.CallBack.ConnectInactive 连接已经被主动登出，不需要重连", uID)
		return
	}
	if extends.IsSystemKick(currentChlCtx) {
		logger.Debug("RPC.CallBack.ConnectInactive 连接已经被踢出，不需要重连", uID)
		return
	}
	// broken: default reconnect
	logger.Debug("RPC.CallBack.ConnectInactive 连接中断，需要重连", uID)
	ip := currentChlCtx.ContextAttr().GetString(bootc.KeyIP)
	port := currentChlCtx.ContextAttr().GetInt(bootc.KeyPort)
	clientFacade.Connect(ip, port, uID)
}

// MessageReceived called on executor routine
func (cb *callBack) MessageReceived(Ctx service.IChannelContext, Data protocol.Protocol) {
	uID := Ctx.ContextAttr().GetInt64(clientkey.KeyFireUser)
	logger.Info(fmt.Sprintf("RPC.CallBack.MessageReceived uID:%v", uID))
	// executor.FireEvent(routineCommands.NewRoutineInboundCmdMsgRecv(uID, Data, Ctx))
	msg := Data.(*messageImpl.Message)
	callService(Ctx, msg)
}

// heart 定时心跳
func heart() {
	go func() {
		for {
			time.Sleep(20 * time.Second)
			sendMessage(mid.HeartbeatRequest, &rpc.HeartbeatRequest{})
		}
	}()
}

// sendMessage
func sendMessage(msgCode uint16, protoData protoreflect.ProtoMessage) {
	chlCtx := chlCtxReference.Search(clientNodeID)
	if chlCtx == nil {
		logger.Warn(fmt.Sprintf("rpc for nodeID %v is not connect", clientNodeID))
		return
	}
	clientFacade.SendMessage(clientNodeID, chlCtx, messageImpl.NewMessage(1, 0, 1, 1, msgCode, protoData), false, "")
}

// callService
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

// Mod returns the remainder of s divided by n
func Mod(s string) int64 {
	var sum int64
	for _, b := range md5.Sum([]byte(s)) {
		sum += int64(b)
	}
	return sum
}
