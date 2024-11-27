package schedule

// IEventMaker创建的所有Events全部都在Event的PoolId对应的协程池中QueueId对应的协程中执行
import (
	"gitee.com/andyxt/gox/handler/protocol"
	"gitee.com/andyxt/gox/service"

	"gitee.com/andyxt/gox/executor"
)

// Inbound
type IRoutineInboundEventMaker interface {
	//收到消息包
	MakeMessageReceivedEvent(routineId int64, Data protocol.IProtocol, Ctx service.IChannelContext) executor.Event
	//新连接
	MakeActiveEvent(routineId int64, Ctx service.IChannelContext) executor.Event
	//连接中断
	MakeInActiveEvent(routineId int64, Ctx service.IChannelContext) executor.Event
}

// Outbound
type IRoutineOutboundEventMaker interface {
	//发起连接
	MakeConnectEvent(routineId int64, ip string, port int, uID int64, params map[string]interface{}) executor.Event
	//关闭连接
	MakeCloseEvent(routineId int64, uID int64, Desc string) executor.Event
	//下发消息包:OnClose是否在消息发送完毕后关闭连接
	MakeSendMessageEvent(routineId int64, Data protocol.IProtocol, OnClose bool, PoolKey int64, ChlCtx service.IChannelContext, Desc string) executor.Event
}
