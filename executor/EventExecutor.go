package executor

var executorInstance *RoutinePool = NewRoutinePool(3000, 10000)

func FireEvent(e Event) {
	executorInstance.FireEvent(e)
}

func FireEventWait(e Event) (result interface{}, ok bool) {
	return executorInstance.FireEventWait(e)
}

//
//import (
//	"bean"
//	"event/events"
//	"executor"
//	"logger"
//	"protocol"
//)
//
//var pool *executor.RoutinePool
//
//func EventExecutorInit(pPoolSize int32, pChanSize int32) {
//	pool = executor.NewRoutinePool(pPoolSize, pChanSize)
//}
//
//func FireEvent(e executor.Event) {
//	pool.Exec(e)
//}
//
///*发送Stc消息*/
//func FireStcMsg(data interface{}, isClose bool, chlCtx *bean.ChannelContext, uid int32, infoStr string) {
//	downEvent := new(events.ChannelDownMsgSendEvent)
//	downEvent.Data = data
//	downEvent.OnClose = isClose
//	if isClose {
//		logger.Info("关闭连接：", uid, " 关闭原因：", infoStr)
//	}
//	downEvent.PoolKey = chlCtx.GetPoolKey()
//	downEvent.ChlCtx = chlCtx
//	FireEvent(downEvent)
//}
//
///*发送内部STS消息(玩家不一定在线)*/
//func FireEventStsMsg(data interface{}, eventQueueId int32) {
//	upEvent := new(events.ChannelUpStsEvent)
//	upEvent.Data = data
//	upEvent.EventQueueId = eventQueueId
//	FireEvent(upEvent)
//}
//
///*发送内部STS消息(玩家不一定在线)*/
//func FireEventStsWaitMsg(evtData interface{}, eventQueueId int32)(result  interface {},ok bool) {
//	waitChan:=make(chan interface {},1)
//	upEvent := new(events.ChannelUpStsEvent)
//	upEvent.Data = evtData
//	upEvent.EventQueueId = eventQueueId
//	upEvent.WaitChan = waitChan
//	FireEvent(upEvent)
//	result, ok = <- waitChan
//	return
//}
//
///*发送内部STS消息(玩家不一定在线)*/
//func FireEventStsConnectMsg( channelParams map[string]interface{}) {
//	upEvent := new(events.ChannelUpConnectEvent)
//	upEvent.ChannelParams = channelParams
//	FireEvent(upEvent)
//}
///*发送内部STS消息(玩家不一定在线)*/
//func FireEventStsCloseMsg(id int32, ip string, port int) {
//	upEvent := new(events.ChannelUpCloseEvent)
//	upEvent.Id = id
//	upEvent.Ip = ip
//	upEvent.Port = port
//	FireEvent(upEvent)
//}
