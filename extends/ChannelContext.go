package extends

import (
	"fmt"

	"github.com/goodluck0107/gox/service"
)

func ChannelContextToString(chlCtx service.IChannelContext) string {
	chlCtxID := chlCtx.ID()
	chlCtxUID := UID(chlCtx)
	return fmt.Sprintf("chlCtxID=%s chlCtxUID=%d ", chlCtxID, chlCtxUID)
}

func ChannelContextEquals(chlCtx service.IChannelContext, other service.IChannelContext) (ret bool) {
	return chlCtx.ID() == other.ID()
}

// 放入用户信息代表后续Command事件中可以执行派发Event事件的行为了
func PutInUserInfo(chlCtx service.IChannelContext, uID int64, lngType int8) {
	chlCtx.ContextAttr().Set("poolKey", uID)     // 用户在连接列表中的Key，目前使用用户UID标识
	chlCtx.ContextAttr().Set("lngType", lngType) // 用户语言类型
	chlCtx.ContextAttr().Set("isInPool", true)   // 连接是否添加到连接池
}

// ResetUserInfo 登出释放用户
func ResetUserInfo(chlCtx service.IChannelContext) {
	chlCtx.ContextAttr().Set("poolKey", 0)      // 用户在连接列表中的Key，目前使用用户UID标识
	chlCtx.ContextAttr().Set("lngType", 0)      // 用户语言类型
	chlCtx.ContextAttr().Set("isInPool", false) // 连接是否添加到连接池
}

func HasUserInfo(chlCtx service.IChannelContext) bool {
	return chlCtx.ContextAttr().GetBool("isInPool")
}

func UID(chlCtx service.IChannelContext) int64 {
	return chlCtx.ContextAttr().GetInt64("poolKey")
}

func GetLngType(chlCtx service.IChannelContext) int8 {
	return chlCtx.ContextAttr().GetInt8("lngType")
}

/**
 * 设置以前的用户连接为废弃，废弃连接的后续消息都将不处理
 * */
func Conflict(chlCtx service.IChannelContext) {
	chlCtx.ContextAttr().Set("userConflict", true) // 连接是否废弃,当有用户重连与异地登陆时候，以前的连接会被置为废弃
}

func IsConflict(chlCtx service.IChannelContext) (ret bool) {
	return chlCtx.ContextAttr().GetBool("userConflict")
}

func IsClose(chlCtx service.IChannelContext) (ret bool) {
	return chlCtx.ContextAttr().GetBool("isClose")
}

func Close(chlCtx service.IChannelContext) {
	chlCtx.Close()
	chlCtx.ContextAttr().Set("isClose", true) //是否已经关闭
}

func IsLogout(chlCtx service.IChannelContext) (ret bool) {
	return chlCtx.ContextAttr().GetBool("isLogout")
}

func Logout(chlCtx service.IChannelContext) {
	chlCtx.ContextAttr().Set("isLogout", true) //是否已经登出
}

func IsSystemKick(chlCtx service.IChannelContext) (ret bool) {
	return chlCtx.ContextAttr().GetBool("isSystemKick")
}

func SystemKick(chlCtx service.IChannelContext) {
	chlCtx.ContextAttr().Set("isSystemKick", true) //是否已经被系统踢出
}

func IsOfflie(chlCtx service.IChannelContext) (ret bool) {
	return chlCtx.ContextAttr().GetBool("isOfflie")
}

func Offlie(chlCtx service.IChannelContext) {
	chlCtx.ContextAttr().Set("isOfflie", true) //是否已经断开
}
