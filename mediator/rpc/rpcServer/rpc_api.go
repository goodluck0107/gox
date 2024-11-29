package rpcServer

import (
	"fmt"

	"gitee.com/andyxt/gona/boot"
	"gitee.com/andyxt/gona/boot/boots"
	"gitee.com/andyxt/gona/logger"
	"gitee.com/andyxt/gona/utils"
	"gitee.com/andyxt/gox/mediator"
	"gitee.com/andyxt/gox/mediator/rpc/mid"
	"gitee.com/andyxt/gox/mediator/rpc/pb/rpc"
	"gitee.com/andyxt/gox/mediator/rpc/rpcServer/services"
	"gitee.com/andyxt/gox/mediator/server/channelCmdMakerImpl"
	"gitee.com/andyxt/gox/mediator/server/channelCmdMakerImpl/channelCommands"
	"gitee.com/andyxt/gox/mediator/server/routineCmdMakerImpl"
	"gitee.com/andyxt/gox/message"
	"gitee.com/andyxt/gox/service"
	"google.golang.org/protobuf/proto"
)

func Start(port int64) {
	service.Register(services.NewService())
	listenRPC(port)
}

// listenRPC 监听远端通知(下注与结算以及更新玩家信息)
func listenRPC(port int64) {
	params := make(map[string]interface{})
	params[boot.KeyPacketBytesCount] = 4
	params[boot.KeyChannelReadLimit] = 10240
	params[boot.KeyReadTimeOut] = -1
	params[boot.KeyWriteTimeOut] = -1
	bootStrap :=
		boots.NewServerBootStrap().
			Params(params).
			Port(fmt.Sprintf(":%v", port)).
			ChannelInitializer(
				mediator.NewChannelInitializer(
					channelCmdMakerImpl.NewChannelInboundCmdMaker(NewNofityMessage(), routineCmdMakerImpl.NewRoutineInboundCmdMaker()), message.NewMessageFactory(),
				))
	err := bootStrap.Listen()
	utils.CheckError(err)
}

func NewNofityMessage() channelCommands.ILoginMessage {
	instance := new(NofityMessage)
	return instance
}

type NofityMessage struct {
}

func (loginMessage *NofityMessage) IsLoginMessage(msg *message.Message) bool {
	return msg.MsgID == mid.RPCLoginRequest
}

func (loginMessage *NofityMessage) IsValid(msg *message.Message) bool {
	loginRequest := &rpc.LoginRequest{}
	loginRequestErr := proto.Unmarshal(msg.Data, loginRequest)
	if loginRequestErr != nil {
		logger.Error("loginRequestErr:", loginRequestErr)
		return false
	}
	if loginRequest.NodeID <= 0 {
		logger.Error("loginRequest.NodeID:", loginRequest.NodeID)
		return false
	}
	return true
}

func (loginMessage *NofityMessage) GetLoginUID(msg *message.Message) int64 {
	loginRequest := &rpc.LoginRequest{}
	loginRequestErr := proto.Unmarshal(msg.Data, loginRequest)
	if loginRequestErr != nil {
		return 0
	}
	return loginRequest.NodeID
}

func (loginMessage *NofityMessage) GetLngType(msg *message.Message) int8 {
	return 0
}
