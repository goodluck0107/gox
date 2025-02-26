package rpcServer

import (
	"fmt"

	"gitee.com/andyxt/gona/boot/boots"
	"gitee.com/andyxt/gox/handler/protocol"
	"gitee.com/andyxt/gox/mediator"
	"gitee.com/andyxt/gox/mediator/rpc/rpcServer/services"
	"gitee.com/andyxt/gox/mediator/server/channelCmdMakerImpl"
	"gitee.com/andyxt/gox/mediator/server/channelCmdMakerImpl/channelCommands"
	"gitee.com/andyxt/gox/mediator/server/routineCmdMakerImpl"
	"gitee.com/andyxt/gox/messageImpl"
	"gitee.com/andyxt/gox/service"
)

func Start(port int64) {
	service.Register(services.NewService())
	listenRPC(port)
}

// listenRPC 监听远端通知(下注与结算以及更新玩家信息)
func listenRPC(port int64) {
	boots.Serve(
		boots.WithTCPAddr(fmt.Sprintf(":%v", port)),
		boots.WithInitializer(mediator.NewChannelInitializer(channelCmdMakerImpl.NewChannelInboundCmdMaker(newLoginMessage(), routineCmdMakerImpl.NewRoutineInboundCmdMaker()), messageImpl.NewMessageFactory())),
		boots.WithReadTimeOut(-1),
		boots.WithWriteTimeOut(-1),
		boots.WithReadLimit(10240),
	)
}

func newLoginMessage() channelCommands.ILoginMessage {
	instance := new(LoginMessage)
	return instance
}

type LoginMessage struct {
}

func (loginMessage *LoginMessage) IsLoginMessage(protocol protocol.Protocol) bool {
	return false
}

func (loginMessage *LoginMessage) IsWhiteMessage(protocol protocol.Protocol) bool {
	return true
}

func (loginMessage *LoginMessage) IsValid(protocol protocol.Protocol) bool {
	return true
}

func (loginMessage *LoginMessage) GetLoginUID(protocol protocol.Protocol) int64 {
	return 0
}

func (loginMessage *LoginMessage) GetLngType(protocol protocol.Protocol) int8 {
	return 0
}
