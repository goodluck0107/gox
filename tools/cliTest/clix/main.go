package main

import (
	"fmt"

	"gitee.com/andyxt/gox/tools/cli/repl"
	"gitee.com/andyxt/gox/tools/cliTest/clix/cliService"
	"gitee.com/andyxt/gox/tools/cliTest/servx/services/account"

	"gitee.com/andyxt/gox/service"
)

func main() {
	fmt.Println("cli")
	service.Register(cliService.NewService())
	service.Register(account.NewService())
	repl.Repl()
}
