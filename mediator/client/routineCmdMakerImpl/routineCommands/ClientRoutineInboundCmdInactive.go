package routineCommands

import (
	"gitee.com/andyxt/gona/boot"
	"gitee.com/andyxt/gox/extends"
	"gitee.com/andyxt/gox/key"
	"gitee.com/andyxt/gox/service"
	"gitee.com/andyxt/gox/session"

	"gitee.com/andyxt/gona/boot/bootc/listener"
	"gitee.com/andyxt/gona/logger"
)

type ClientRoutineInboundCmdInactive struct {
	routineId int64
	connector listener.IConnector
	ChlCtx    service.IChannelContext
}

func NewClientRoutineInboundCmdInactive(routineId int64,
	ChlCtx service.IChannelContext, connector listener.IConnector) (this *ClientRoutineInboundCmdInactive) {
	this = new(ClientRoutineInboundCmdInactive)
	this.routineId = routineId
	this.ChlCtx = ChlCtx
	this.connector = connector
	return
}

func (inactiveEvent *ClientRoutineInboundCmdInactive) QueueId() int64 {
	return inactiveEvent.routineId
}

func (inactiveEvent *ClientRoutineInboundCmdInactive) Wait() (result interface{}, ok bool) {
	return nil, true
}

func (inactiveEvent *ClientRoutineInboundCmdInactive) Exec() {
	logger.Debug("ClientRoutineInboundCmdInactive Exec")
	if extends.IsConflict(inactiveEvent.ChlCtx) {
		logger.Debug("ClientRoutineInboundCmdInactive 连接已经被其他连接挤下线，不需要重连", extends.UID(inactiveEvent.ChlCtx))
		return
	}
	if extends.IsClose(inactiveEvent.ChlCtx) {
		logger.Debug("ClientRoutineInboundCmdInactive 连接已经被主动关闭，不需要重连", extends.UID(inactiveEvent.ChlCtx))
		return
	}
	if extends.IsLogout(inactiveEvent.ChlCtx) {
		logger.Debug("ClientRoutineInboundCmdInactive 连接已经被主动登出，不需要重连", extends.UID(inactiveEvent.ChlCtx))
		return
	}
	if extends.IsSystemKick(inactiveEvent.ChlCtx) {
		logger.Debug("ClientRoutineInboundCmdInactive 连接已经被踢出，不需要重连", extends.UID(inactiveEvent.ChlCtx))
		return
	}
	uId := extends.UID(inactiveEvent.ChlCtx)
	iSession := session.GetSession(uId)
	if iSession == nil {
		logger.Debug("ClientRoutineInboundCmdInactive 连接已经被主动关闭，不需要重连", extends.UID(inactiveEvent.ChlCtx))
		return
	}
	//default: broken reconnect
	logger.Debug("ClientRoutineInboundCmdInactive 连接中断，需要重连", extends.UID(inactiveEvent.ChlCtx))
	ip := inactiveEvent.ChlCtx.ContextAttr().GetString(key.ChannelIp)
	port := inactiveEvent.ChlCtx.ContextAttr().GetInt(key.ChannelPort)
	channelTag := inactiveEvent.ChlCtx.ContextAttr().GetString(key.ChannelTag)
	jsonData := inactiveEvent.ChlCtx.ContextAttr().GetString(key.ChannelParams)
	channelReadLimit := inactiveEvent.ChlCtx.ContextAttr().GetString(boot.KeyChannelReadLimit)
	params := make(map[string]interface{})
	params[key.ChannelIp] = ip
	params[key.ChannelPort] = port
	params[key.ChannelFireUser] = uId
	params[key.ChannelTag] = channelTag
	params[key.ChannelParams] = jsonData
	params[boot.KeyChannelReadLimit] = channelReadLimit
	inactiveEvent.connector.Connect(ip, port, params)
}
