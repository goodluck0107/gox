package main

import (
	"fmt"

	"github.com/goodluck0107/gox/tools/cli/repl"
	"github.com/goodluck0107/gox/tools/cliTest/clix/cliService"
	"github.com/goodluck0107/gox/tools/cliTest/servx/services/account"

	"github.com/goodluck0107/gox/service"
)

func main() {
	fmt.Println("cli")
	service.Register(cliService.NewService())
	service.Register(account.NewService())
	repl.Repl()
}
