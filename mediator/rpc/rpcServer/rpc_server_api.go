package rpcServer

import (
	"fmt"

	"github.com/goodluck0107/gona/boot/boots"
	"github.com/goodluck0107/gox/code/protocol"
	"github.com/goodluck0107/gox/mediator"
	"github.com/goodluck0107/gox/mediator/rpc/rpcServer/services"
	"github.com/goodluck0107/gox/mediator/server/channelCmdMakerImpl"
	"github.com/goodluck0107/gox/mediator/server/channelCmdMakerImpl/channelCommands"
	"github.com/goodluck0107/gox/mediator/server/routineCmdMakerImpl"
	"github.com/goodluck0107/gox/messageImpl"
	"github.com/goodluck0107/gox/service"
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
