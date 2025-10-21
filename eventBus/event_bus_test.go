package eventBus

import (
	"fmt"
	"testing"

	"github.com/goodluck0107/gona/boot/channel"
	"github.com/goodluck0107/gona/utils"
	"github.com/goodluck0107/gox/service"
)

func TestEventBus(t *testing.T) {
	listenEvt()
	triggerEvt()
}

func listenEvt() {
	On("Inactive", onInactive)
}

func triggerEvt() {
	Trigger("Inactive", int64(10086), NewMockChannelHandlerContext("name"))
	TriggerCross("Inactive", 123456, int64(10086), NewMockChannelHandlerContext("name"))
	TriggerCrossWait("Inactive", 123456, int64(10086), NewMockChannelHandlerContext("name"))
}

// onInactive 新的玩家下注
func onInactive(data ...interface{}) {
	PlayerID := data[0].(int64)
	Context := data[1].(service.IChannelContext)
	fmt.Println("onInactive", "PlayerID:", PlayerID)
	fmt.Println("onInactive", "key1:", Context.ContextAttr().GetString("key1"))
}

type MockChannelHandlerContext struct {
	mAttr *channel.Attr
	Name  string
	mID   string
}

func NewMockChannelHandlerContext(name string) (context *MockChannelHandlerContext) {
	context = new(MockChannelHandlerContext)
	context.Name = name
	context.mID = utils.UUID()
	context.mAttr = channel.NewAttr(nil)
	context.mAttr.Set("key1", "value1")
	return
}

/*发起关闭事件，消息将被送往管道处理*/
func (context *MockChannelHandlerContext) Close() {
	fmt.Println("MockChannelHandlerContext", context.Name, "链接被关闭")
}

/*发起写事件，消息将被送往管道处理*/
func (context *MockChannelHandlerContext) Write(e interface{}) {
	fmt.Println("MockChannelHandlerContext", context.Name, "链接收到回包", fmt.Sprintf("%v", e))
}

func (context *MockChannelHandlerContext) ContextAttr() channel.IAttr {
	return context.mAttr
}

func (context *MockChannelHandlerContext) ID() string {
	return context.mID
}

func (context *MockChannelHandlerContext) RemoteAddr() string {
	return ""
}
