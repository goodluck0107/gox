package main

import (
	"fmt"
	"time"

	"github.com/goodluck0107/gox/mediator/rpc/mid"
	"github.com/goodluck0107/gox/mediator/rpc/pb/rpc"
	"github.com/goodluck0107/gox/mediator/rpc/rpcClient"
	"github.com/goodluck0107/gox/mediator/rpc/tests/testCli/services"
	"github.com/goodluck0107/gox/service"
)

func main() {
	service.Register(services.NewService())
	rpcClient.Start("testCli", "127.0.0.1", 10086)
	for {
		time.Sleep(time.Second * 5)
		call()
	}
}

func call() {
	go func() {
		for {
			time.Sleep(20 * time.Second)
			playerID := int64(1)
			rpcClient.Publish(fmt.Sprintf("%v_%v_%v", playerID, "testCli", mid.EchoRequest), &rpc.EchoRequest{
				Param1: 1,
				Param2: "2",
				Param3: []byte("Echo"),
			})
		}
	}()
}
