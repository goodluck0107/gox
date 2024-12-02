package services

import (
	"gitee.com/andyxt/gox/mediator/server"
)

type RpcService struct{}

func NewService() *RpcService {
	listenEvent()
	return &RpcService{}
}

func listenEvent() {
	server.OnClose(onInactive)
}
