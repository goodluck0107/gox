package commands

import (
	"strings"

	"github.com/abiosoft/ishell/v2"
)

// Greet greet命令
func Greet(c *ishell.Context) {
	c.Println("Hello", strings.Join(c.Args, " "))
}
