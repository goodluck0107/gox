package account

import (
	"github.com/goodluck0107/gox/mediator/server"
)

type AccountService struct{}

func NewService() *AccountService {
	listenEvent()
	return &AccountService{}
}

func listenEvent() {
	server.OnClose(onInactive)
}
