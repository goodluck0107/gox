package services

import (
	"gitee.com/andyxt/gox/service"
	"gitee.com/andyxt/gox/tools/cliTest/servx/services/account"
)

func Register() {
	service.Register(account.NewService())
}
