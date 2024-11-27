package routineCommands

import (
	"gitee.com/andyxt/gox/key"
	"gitee.com/andyxt/gox/session"

	"gitee.com/andyxt/gona/boot/bootc/listener"
	"gitee.com/andyxt/gona/logger"
	"gitee.com/andyxt/gona/utils"
)

type ChannelParams struct {
	Uid   int64
	Token string
}

type ClientRoutineInboundCmdConnect struct {
	routineId int64
	uID       int64
	ip        string
	port      int
	params    map[string]interface{}
	connector listener.IConnector
}

func NewClientRoutineInboundCmdConnect(routineId int64, uID int64, ip string, port int, params map[string]interface{}, connector listener.IConnector) (this *ClientRoutineInboundCmdConnect) {
	this = new(ClientRoutineInboundCmdConnect)
	this.routineId = routineId
	this.uID = uID
	this.ip = ip
	this.port = port
	this.params = params
	this.connector = connector
	return
}

func (upConnectEvent *ClientRoutineInboundCmdConnect) QueueId() int64 {
	return upConnectEvent.routineId
}

func (upConnectEvent *ClientRoutineInboundCmdConnect) Wait() (result interface{}, ok bool) {
	return nil, true
}

func (upConnectEvent *ClientRoutineInboundCmdConnect) Exec() {
	logger.Debug("ClientRoutineInboundCmdConnect Exec")
	iSession := session.GetSession(0, upConnectEvent.uID) // find in online players
	if iSession == nil {
		logger.Debug("ClientRoutineInboundCmdConnect iSession == nil")
		iSession = session.NewSession(utils.UUID(), upConnectEvent.uID)
		session.AddSession(0, iSession)
	}
	upConnectEvent.params[key.ChannelIp] = upConnectEvent.ip
	upConnectEvent.params[key.ChannelPort] = upConnectEvent.port
	upConnectEvent.params[key.ChannelFireUser] = upConnectEvent.uID
	upConnectEvent.params[key.ChannelTag] = ""
	upConnectEvent.params[key.ChannelParams] = ""
	upConnectEvent.connector.Connect(upConnectEvent.ip, upConnectEvent.port, upConnectEvent.params)
}
