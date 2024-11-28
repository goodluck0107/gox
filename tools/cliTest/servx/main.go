package main

import (
	"fmt"
	"runtime/debug"

	"gitee.com/andyxt/gona/boot"
	"gitee.com/andyxt/gona/boot/boots"
	"gitee.com/andyxt/gona/logger"
	"gitee.com/andyxt/gona/utils"
	"gitee.com/andyxt/gox/mediator"
	"gitee.com/andyxt/gox/mediator/server/channelCmdMakerImpl"
	"gitee.com/andyxt/gox/mediator/server/channelCmdMakerImpl/channelCommands"
	"gitee.com/andyxt/gox/mediator/server/routineCmdMakerImpl"
	"gitee.com/andyxt/gox/message"
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
	params := make(map[string]interface{})
	params[boot.KeyChannelReadLimit] = 512
	bootStrap :=
		boots.NewServerBootStrap().
			Params(params).
			Port(fmt.Sprintf(":%v", 20000)).
			ChannelInitializer(
				mediator.NewChannelInitializer(
					channelCmdMakerImpl.NewChannelInboundCmdMaker(NewEnterMessage(), routineCmdMakerImpl.NewRoutineInboundCmdMaker()), message.NewMessageFactory(),
				))
	err := bootStrap.Listen()
	logger.Error("game exit error:", err)
	utils.CheckError(err)
}

func NewEnterMessage() channelCommands.ILoginMessage {
	instance := new(EnterMessage)
	return instance
}

type EnterMessage struct {
}

func (loginMessage *EnterMessage) IsLoginMessage(msg *message.Message) bool {
	return msg.MsgID == mid.LoginRequest
}

func (loginMessage *EnterMessage) IsValid(msg *message.Message) bool {
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

func (loginMessage *EnterMessage) GetLoginUID(msg *message.Message) int64 {
	loginRequest := &cli.LoginRequest{}
	loginRequestErr := proto.Unmarshal(msg.Data, loginRequest)
	if loginRequestErr != nil {
		return 0
	}
	return loginRequest.UID
}

func (loginMessage *EnterMessage) GetLngType(msg *message.Message) int8 {
	return 0
}
