package repl

import (
	"github.com/goodluck0107/gox/tools/cli/repl/commands"
)

// Repl start a shell for user
func Repl() {
	commands.Init()
	shell := ishell.New()
	shell.Println("Sample Interactive Shell")
	shell.AddCmd(&ishell.Cmd{
		Name: "greet",
		Help: "greet user",
		Func: commands.Greet,
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "connect",
		Help: "connects to server",
		Func: commands.Connect,
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "request",
		Help: "makes a request to server",
		Func: commands.Request,
	})
	shell.Run()
}
