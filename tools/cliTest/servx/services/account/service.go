package account

import (
	"gitee.com/andyxt/gox/eventBus"
)

type AccountService struct{}

func NewService() *AccountService {
	listenEvent()
	return &AccountService{}
}

func listenEvent() {
	eventBus.On("Inactive", onInactive)
}
