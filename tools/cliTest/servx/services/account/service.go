package account

import (
	"gitee.com/andyxt/gox/mediator/server"
)

type AccountService struct{}

func NewService() *AccountService {
	listenEvent()
	return &AccountService{}
}

func listenEvent() {
	server.OnClose(onInactive)
}
