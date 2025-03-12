package main

import (
	"fmt"
	"runtime/debug"

	"gitee.com/andyxt/gona/boot/boots"
	"gitee.com/andyxt/gona/logger"
	"gitee.com/andyxt/gox/code/protocol"
	"gitee.com/andyxt/gox/mediator"
	"gitee.com/andyxt/gox/mediator/server/channelCmdMakerImpl"
	"gitee.com/andyxt/gox/mediator/server/channelCmdMakerImpl/channelCommands"
	"gitee.com/andyxt/gox/mediator/server/routineCmdMakerImpl"
	"gitee.com/andyxt/gox/messageImpl"
	"gitee.com/andyxt/gox/tools/cliTest/generic/mid"
	"gitee.com/andyxt/gox/tools/cliTest/pb/cli"
	"gitee.com/andyxt/gox/tools/cliTest/servx/services"
	"google.golang.org/protobuf/proto"
)

func main() {
	defer func() {
		if recoverErr := recover(); recoverErr != nil {
			logger.Error("game exit ", fmt.Sprint(recoverErr, string(debug.Stack())))
		}
	}()
	startServer()
}

func startServer() {
	services.Register()
	listenSocket()
}

func serviceName() string {
	return "hall"
}

// listenSocket 监听TcpSocket
func listenSocket() {
	boots.Serve(
		boots.WithTCPAddr(fmt.Sprintf(":%v", 20000)),
		boots.WithInitializer(mediator.NewChannelInitializer(channelCmdMakerImpl.NewChannelInboundCmdMaker(NewEnterMessage(), routineCmdMakerImpl.NewRoutineInboundCmdMaker()), messageImpl.NewMessageFactory())),
		boots.WithReadLimit(512),
	)
}

func NewEnterMessage() channelCommands.ILoginMessage {
	instance := new(EnterMessage)
	return instance
}

type EnterMessage struct {
}

func (loginMessage *EnterMessage) IsLoginMessage(protocol protocol.Protocol) bool {
	msg, ok := protocol.(*messageImpl.Message)
	if !ok {
		return false
	}
	return msg.MsgID == mid.LoginRequest
}
func (loginMessage *EnterMessage) IsWhiteMessage(protocol protocol.Protocol) bool {
	return false
}
func (loginMessage *EnterMessage) IsValid(protocol protocol.Protocol) bool {
	msg, ok := protocol.(*messageImpl.Message)
	if !ok {
		return false
	}
	loginRequest := &cli.LoginRequest{}
	loginRequestErr := proto.Unmarshal(msg.Data, loginRequest)
	if loginRequestErr != nil {
		logger.Error("loginRequestErr:", loginRequestErr)
		return false
	}
	if loginRequest.UID <= 0 {
		logger.Error("loginRequest.UID:", loginRequest.UID)
		return false
	}
	if loginRequest.Token == "" {
		logger.Error("loginRequest.Token:", loginRequest.Token)
		return false
	}
	return true
}

func (loginMessage *EnterMessage) GetLoginUID(protocol protocol.Protocol) int64 {
	msg, ok := protocol.(*messageImpl.Message)
	if !ok {
		return 0
	}
	loginRequest := &cli.LoginRequest{}
	loginRequestErr := proto.Unmarshal(msg.Data, loginRequest)
	if loginRequestErr != nil {
		return 0
	}
	return loginRequest.UID
}

func (loginMessage *EnterMessage) GetLngType(protocol protocol.Protocol) int8 {
	return 0
}
