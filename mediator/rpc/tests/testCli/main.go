package main

import (
	"fmt"
	"time"

	"gitee.com/andyxt/gox/mediator/rpc/mid"
	"gitee.com/andyxt/gox/mediator/rpc/pb/rpc"
	"gitee.com/andyxt/gox/mediator/rpc/rpcClient"
	"gitee.com/andyxt/gox/mediator/rpc/tests/testCli/services"
	"gitee.com/andyxt/gox/service"
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
