package inboundCommands

import (
	"gitee.com/andyxt/gox/internal/logger"
)

type ClientChannelUpActiveFailEvent struct {
	routineId int64
	callback  ICallBack
	err       error
	params    map[string]interface{}
}

func NewClientChannelUpActiveFailEvent(routineId int64, callback ICallBack, err error, params map[string]interface{}) (this *ClientChannelUpActiveFailEvent) {
	this = new(ClientChannelUpActiveFailEvent)
	this.routineId = routineId
	this.callback = callback
	this.err = err
	this.params = params
	return
}

func (activeEvent *ClientChannelUpActiveFailEvent) QueueId() int64 {
	return activeEvent.routineId
}

func (activeEvent *ClientChannelUpActiveFailEvent) Wait() (result interface{}, ok bool) {
	return nil, true
}

func (activeEvent *ClientChannelUpActiveFailEvent) Exec() {
	logger.Debug("ClientChannelUpActiveFailEvent Exec", "新连接建立失败")
	activeEvent.callback.ConnectFail(activeEvent.err, activeEvent.params)
}
