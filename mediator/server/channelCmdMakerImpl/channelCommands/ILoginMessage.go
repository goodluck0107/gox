package channelCommands

import "github.com/goodluck0107/gox/code/protocol"

/*
一条socket链接建立以后，发送给服务端的第一条消息必须实现该接口，用于分配处理该链接的协程
*/
type ILoginMessage interface {
	IsLoginMessage(protocol protocol.Protocol) bool
	IsValid(protocol protocol.Protocol) bool
	GetLoginUID(protocol protocol.Protocol) int64
	GetLngType(protocol protocol.Protocol) int8
	IsWhiteMessage(protocol protocol.Protocol) bool
}
