package services

import (
	"gitee.com/andyxt/gox/eventBus"
)

type RpcService struct{}

func NewService() *RpcService {
	listenEvent()
	return &RpcService{}
}
func listenEvent() {
	eventBus.On("Inactive", onInactive)
}
