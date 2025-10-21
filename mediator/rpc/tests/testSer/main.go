package main

import (
	"github.com/goodluck0107/gox/mediator/rpc/rpcServer"
)

func main() {
	rpcServer.Start(10086)
}
