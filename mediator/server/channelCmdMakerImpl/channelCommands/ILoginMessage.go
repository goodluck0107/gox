package channelCommands

import "gitee.com/andyxt/gox/message"

/*
一条socket链接建立以后，发送给服务端的第一条消息必须实现该接口，用于分配处理该链接的协程
*/
type ILoginMessage interface {
	IsLoginMessage(msg *message.Message) bool
	IsValid(msg *message.Message) bool
	GetLoginUID(msg *message.Message) int64
	GetLngType(msg *message.Message) int8
}
