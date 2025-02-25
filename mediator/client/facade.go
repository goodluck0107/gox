package client

import (
	"gitee.com/andyxt/gona/boot"
	"gitee.com/andyxt/gona/boot/bootc"
	"gitee.com/andyxt/gona/boot/bootc/connector"
	"gitee.com/andyxt/gona/boot/bootc/listener"
	"gitee.com/andyxt/gox/executor"
	"gitee.com/andyxt/gox/handler/protocol"
	"gitee.com/andyxt/gox/mediator"
	"gitee.com/andyxt/gox/mediator/client/channelCmdMakerImpl"
	"gitee.com/andyxt/gox/mediator/client/clientkey"
	"gitee.com/andyxt/gox/mediator/client/routineCmdMakerImpl"
	"gitee.com/andyxt/gox/mediator/client/routineCmdMakerImpl/inboundCommands"
	"gitee.com/andyxt/gox/mediator/client/routineCmdMakerImpl/outboundCommands"
	"gitee.com/andyxt/gox/messageImpl"
	"gitee.com/andyxt/gox/service"
)

func BootClient(callback inboundCommands.ICallBack) *ClientFacade {
	connector := bootc.Serv(
		bootc.WithInitializer(mediator.NewChannelInitializer(
			channelCmdMakerImpl.NewClientInboundCommandMaker(routineCmdMakerImpl.NewClientInboundEventMakerImpl(callback)),
			messageImpl.NewMessageFactory())),
	)
	return newClientFacade(connector, callback)
}

func (facade *ClientFacade) Connect(ip string, port int, uID int64) {
	connType := connector.NormalSocket
	params := make(map[string]interface{})
	// params[boot.KeyPacketBytesCount] = 4
	params[boot.KeyConnType] = connType
	params[boot.KeyIP] = ip
	params[boot.KeyPort] = port
	params[clientkey.KeyFireUser] = uID
	executor.FireEvent(outboundCommands.NewClientRoutineInboundCmdConnect(uID, ip, port, connType, params,
		facade.mConnector, facade.mCallback))
}

func (facade *ClientFacade) Close(uID int64, ChlCtx service.IChannelContext, Desc string) {
	executor.FireEvent(outboundCommands.NewClientRoutineInboundCmdClose(uID, ChlCtx, Desc))
}

// SendMessage OnClose:是否在消息发送完毕后关闭连接
func (facade *ClientFacade) SendMessage(uID int64, ChlCtx service.IChannelContext, Data protocol.Protocol, OnClose bool, Desc string) {
	executor.FireEvent(outboundCommands.NewClientRoutineOutboundCmdMsgSend(Data, OnClose, uID, ChlCtx, Desc))
}

func newClientFacade(mConnector listener.IConnector, mCallback inboundCommands.ICallBack) (facade *ClientFacade) {
	facade = new(ClientFacade)
	facade.mConnector = mConnector
	facade.mCallback = mCallback
	return
}

type ClientFacade struct {
	mConnector listener.IConnector
	mCallback  inboundCommands.ICallBack
}
