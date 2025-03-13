package outboundCommands

import (
	"gitee.com/andyxt/gona/boot/bootc/connector"
	"gitee.com/andyxt/gona/boot/bootc/listener"
	"gitee.com/andyxt/gox/executor"
	"gitee.com/andyxt/gox/internal/logger"
	"gitee.com/andyxt/gox/mediator/client/routineCmdMakerImpl/inboundCommands"
)

type ChannelParams struct {
	Uid   int64
	Token string
}

type ClientRoutineInboundCmdConnect struct {
	connType  connector.SocketType
	uID       int64
	ip        string
	port      int
	params    map[string]interface{}
	connector listener.IConnector
	callback  inboundCommands.ICallBack
}

func NewClientRoutineInboundCmdConnect(uID int64, ip string, port int, connType connector.SocketType,
	params map[string]interface{},
	connector listener.IConnector, callback inboundCommands.ICallBack) (this *ClientRoutineInboundCmdConnect) {
	this = new(ClientRoutineInboundCmdConnect)
	this.uID = uID
	this.ip = ip
	this.port = port
	this.connType = connType
	this.params = params
	this.connector = connector
	this.callback = callback
	return
}

func (upConnectEvent *ClientRoutineInboundCmdConnect) QueueId() int64 {
	return upConnectEvent.uID
}

func (upConnectEvent *ClientRoutineInboundCmdConnect) Wait() (result interface{}, ok bool) {
	return nil, true
}

func (upConnectEvent *ClientRoutineInboundCmdConnect) Exec() {
	logger.Debug("ClientRoutineInboundCmdConnect Exec:", upConnectEvent.params)
	upConnectEvent.connector.Connect(upConnectEvent.connType, upConnectEvent.ip, upConnectEvent.port,
		upConnectEvent.params, upConnectEvent.failFunc)
}

func (upConnectEvent *ClientRoutineInboundCmdConnect) failFunc(err error, params map[string]interface{}) {
	logger.Debug("ClientRoutineInboundCmdConnect failFunc")
	executor.FireEvent(inboundCommands.NewClientChannelUpActiveFailEvent(upConnectEvent.uID, upConnectEvent.callback, err, params))
}
