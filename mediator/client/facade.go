package client

import (
	"errors"

	"gitee.com/andyxt/gona/boot"
	"gitee.com/andyxt/gona/boot/bootc"
	"gitee.com/andyxt/gona/boot/bootc/connector"
	"gitee.com/andyxt/gox/executor"
	"gitee.com/andyxt/gox/handler/protocol"
	"gitee.com/andyxt/gox/handler/schedule"
	"gitee.com/andyxt/gox/key"
	"gitee.com/andyxt/gox/mediator"
	"gitee.com/andyxt/gox/mediator/client/channelCmdMakerImpl"
	"gitee.com/andyxt/gox/mediator/client/routineCmdMakerImpl"
	"gitee.com/andyxt/gox/mediator/client/routineCmdMakerImpl/routineCommands"
	"gitee.com/andyxt/gox/message"
	"gitee.com/andyxt/gox/service"
)

func BootClient(callback routineCommands.ICallBack) *ClientFacade {
	bootStrap :=
		bootc.NewClientBootStrap(connector.NormalSocket)
	connector := bootStrap.GetConnector()
	bootStrap.ChannelInitializer(
		mediator.NewTcpChannelInitializer(
			channelCmdMakerImpl.NewClientInboundCommandMaker(routineCmdMakerImpl.NewClientInboundEventMakerImpl(connector, callback)),
			message.NewMessageFactory()))
	bootStrap.Listen()
	return NewClientFacade(routineCmdMakerImpl.NewClientOutboundEventMaker(connector))
}

type ClientFacade struct {
	mEventMaker schedule.IRoutineOutboundEventMaker
}

func NewClientFacade(mEventMaker schedule.IRoutineOutboundEventMaker) (facade *ClientFacade) {
	facade = new(ClientFacade)
	facade.mEventMaker = mEventMaker
	return
}

func (facade *ClientFacade) Connect(ip string, port int, uID int64) {
	if facade == nil {
		panic(errors.New("clientFacade == nil while Connect"))
	}
	params := make(map[string]interface{})
	params[boot.KeyPacketBytesCount] = 4
	params[boot.KeyChannelReadLimit] = 10240
	params[boot.KeyReadTimeOut] = -1
	params[boot.KeyWriteTimeOut] = -1
	params[key.ChannelFireUser] = uID
	executor.FireEvent(facade.mEventMaker.MakeConnectEvent(uID, ip, port, uID, params))
}

func (facade *ClientFacade) Close(uID int64, Desc string) {
	if facade == nil {
		panic(errors.New("clientFacade == nil while Close"))
	}
	executor.FireEvent(facade.mEventMaker.MakeCloseEvent(uID, uID, Desc))
}

// SendMessage OnClose:是否在消息发送完毕后关闭连接
func (facade *ClientFacade) SendMessage(Data protocol.IProtocol, OnClose bool, PoolKey int64, ChlCtx service.IChannelContext, Desc string) {
	if facade == nil {
		panic(errors.New("clientFacade == nil while SendMessage"))
	}
	executor.FireEvent(facade.mEventMaker.MakeSendMessageEvent(PoolKey, Data, OnClose, PoolKey, ChlCtx, Desc))
}
