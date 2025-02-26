package services

import (
	"fmt"
	"hash/fnv"

	"gitee.com/andyxt/gona/logger"
	"gitee.com/andyxt/gox/mediator/rpc/rpcServer/center"
	"gitee.com/andyxt/gox/mediator/server"
	"gitee.com/andyxt/gox/service"
)

type RpcService struct{}

func NewService() *RpcService {
	listenEvent()
	return &RpcService{}
}

func listenEvent() {
	server.OnClose(onInactive)
}

// onInactive 连接中断
func onInactive(playerID int64, chlCtx service.IChannelContext) {
	logger.Info(fmt.Sprintf("onInactive ctxID:%v 掉线成功", chlCtx.ID()))
	center.RemoveChannel(chlCtx)
}

// stringToInt64 根据字符串生成 int64
func stringToInt64(s string) int64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	return int64(h.Sum64())
}
