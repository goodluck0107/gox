package client

import (
	"github.com/goodluck0107/gona/boot/bootc"
	"github.com/goodluck0107/gona/boot/bootc/connector"
	"github.com/goodluck0107/gona/boot/bootc/listener"
	"github.com/goodluck0107/gox/code/protocol"
	"github.com/goodluck0107/gox/executor"
	"github.com/goodluck0107/gox/mediator"
	"github.com/goodluck0107/gox/mediator/client/channelCmdMakerImpl"
	"github.com/goodluck0107/gox/mediator/client/clientkey"
	"github.com/goodluck0107/gox/mediator/client/routineCmdMakerImpl"
	"github.com/goodluck0107/gox/mediator/client/routineCmdMakerImpl/inboundCommands"
	"github.com/goodluck0107/gox/mediator/client/routineCmdMakerImpl/outboundCommands"
	"github.com/goodluck0107/gox/messageImpl"
	"github.com/goodluck0107/gox/service"
)

func BootClient(callback inboundCommands.ICallBack) *ClientFacade {
	connector := bootc.Serv(
		bootc.WithInitializer(mediator.NewChannelInitializer(
			channelCmdMakerImpl.NewClientInboundCommandMaker(routineCmdMakerImpl.NewClientInboundEventMakerImpl(callback)),
			messageImpl.NewMessageFactory())),
		bootc.WithReadLimit(10240),
	)
	return newClientFacade(connector, callback)
}

func (facade *ClientFacade) Connect(ip string, port int, uID int64) {
	connType := connector.NormalSocket
	params := make(map[string]interface{})
	// params[boot.KeyPacketBytesCount] = 4
	params[bootc.KeyConnType] = connType
	params[bootc.KeyIP] = ip
	params[bootc.KeyPort] = port
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
