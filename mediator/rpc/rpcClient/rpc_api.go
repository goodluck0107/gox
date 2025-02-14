package rpcClient

import (
	"crypto/md5"
	"fmt"
	"time"

	"gitee.com/andyxt/gox/mediator/client/clientkey"
	"gitee.com/andyxt/gox/mediator/client/routineCmdMakerImpl/inboundCommands"
	"gitee.com/andyxt/gox/messageImpl"
	"gitee.com/andyxt/gox/service"

	"gitee.com/andyxt/gox/mediator/rpc/mid"

	"gitee.com/andyxt/gox/mediator/rpc/pb/rpc"

	"gitee.com/andyxt/gox/mediator/client"

	"gitee.com/andyxt/gona/boot"
	"gitee.com/andyxt/gona/logger"
	"gitee.com/andyxt/gox/extends"
	"gitee.com/andyxt/gox/handler/protocol"

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

// RPCCall 调用master
func RPCCall(nodeName string, playerID int64, msgCode uint16, protoData protoreflect.ProtoMessage) {
	marshalV, marshalE := proto.Marshal(protoData)
	if marshalE != nil {
		logger.Error(fmt.Sprintf("RPC.RPCCall err:%v", marshalE))
		return
	}
	sendMessage(mid.RPCCallRequest, &rpc.RPCCallRequest{
		NodeID:   Mod(nodeName),
		PlayerID: playerID,
		FuncCode: int64(msgCode),
		FuncData: marshalV,
	})
}

// RPCBroadcastRequest 调用master
func RPCBroadcastRequest(playerID int64, msgCode uint16, protoData protoreflect.ProtoMessage) {
	marshalV, marshalE := proto.Marshal(protoData)
	if marshalE != nil {
		logger.Error(fmt.Sprintf("RPC.RPCBroadcastRequest err:%v", marshalE))
		return
	}
	sendMessage(mid.RPCBroadcastRequest, &rpc.RPCBroadcastRequest{
		PlayerID: playerID,
		FuncCode: int64(msgCode),
		FuncData: marshalV,
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
	login()
	heart()
}

// ConnectFail called on executor routine
func (cb *callBack) ConnectFail(err error, params map[string]interface{}) {
	ip := params[boot.KeyIP].(string)
	port := params[boot.KeyPort].(int)
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
	ip := currentChlCtx.ContextAttr().GetString(boot.KeyIP)
	port := currentChlCtx.ContextAttr().GetInt(boot.KeyPort)
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

// login 登录master
func login() {
	sendMessage(mid.RPCLoginRequest, &rpc.LoginRequest{NodeID: clientNodeID})
}

// heart 定时心跳
func heart() {
	go func() {
		for {
			time.Sleep(20 * time.Second)
			sendMessage(mid.RPCHeartbeatRequest, &rpc.HeartbeatRequest{})
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
	serviceCode := int32(msg.MsgID)
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
