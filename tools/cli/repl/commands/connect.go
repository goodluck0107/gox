package commands

import (
	"strconv"

	"github.com/abiosoft/ishell/v2"
)

// Connect connect命令
func Connect(c *ishell.Context) {
	if err := connect(c.RawArgs[1:]); err != nil {
		c.Err(err)
	}
}

// connect 连接server
func connect(args []string) (err error) {
	port, err1 := strconv.ParseInt(args[0], 10, 64)
	if err1 != nil {
		port = 30000
	}
	uID, err2 := strconv.ParseInt(args[1], 10, 64)
	if err2 != nil {
		uID = 0
	}
	clientFacade.Connect("127.0.0.1", int(port), uID)
	return nil
}
