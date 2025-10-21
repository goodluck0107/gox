package services

import (
	"github.com/goodluck0107/gox/service"
	"github.com/goodluck0107/gox/tools/cliTest/servx/services/account"
)

func Register() {
	service.Register(account.NewService())
}
