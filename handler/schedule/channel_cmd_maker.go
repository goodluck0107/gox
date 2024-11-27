package schedule

import "gitee.com/andyxt/gox/service"

// IChannelInboundCommandMaker 创建的所有Commands全部都在Tcp读协程中执行
// IChannelOutboundCommandMaker 创建的所有Commands全部都在Tcp写协程中执行

type ICommand interface {
	Exec()
}

type IChannelInboundCommandMaker interface {
	//触发异常
	MakeExceptionCommand(ctx service.IChannelContext, err error) ICommand

	//新连接
	MakeActiveCommand(Ctx service.IChannelContext) ICommand
	//连接中断
	MakeInActiveCommand(Ctx service.IChannelContext) ICommand
	//收到消息包
	MakeMessageReceivedCommand(Ctx service.IChannelContext, Data interface{}) ICommand
}

type IChannelOutboundCommandMaker interface {
	// 触发异常
	MakeExceptionCommand(ctx service.IChannelContext, err error) ICommand

	// 请求关闭连接
	MakeCloseCommand(Ctx service.IChannelContext) ICommand
	// 下发消息包
	MakeMessageSendCommand(Ctx service.IChannelContext, Data interface{}) ICommand
}
