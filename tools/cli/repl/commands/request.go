package commands

import (
	"errors"
	"fmt"
	"strings"

	"gitee.com/andyxt/gox/message"

	"gitee.com/andyxt/gox/tools/cli/msgproto"

	"gitee.com/andyxt/gox/service"

	"github.com/abiosoft/ishell/v2"
	"google.golang.org/protobuf/proto"
)

// Request request命令
func Request(c *ishell.Context) {
	err := request(c.RawArgs[1:])
	if err != nil {
		c.Err(err)
	}
}

// request 发请求
func request(args []string) error {
	if len(args) < 1 {
		return errors.New(`request should be in the format: request {route} [data]`)
	}

	msgPath := args[0]
	fmt.Println("msgPath:", msgPath)
	var byteArr []byte
	if len(args) > 1 {
		byteArr = []byte(strings.Join(args[1:], ""))
	}
	fmt.Println("msgData:", string(byteArr))
	handlerTypes := service.HandlerType(msgPath)
	if handlerTypes == nil {
		return errors.New(`no handler for route`)
	}
	adapterResult := msgproto.AdaptArgsFromJson(handlerTypes, byteArr)
	pb, ok := adapterResult.Interface().(proto.Message)
	if !ok {
		return errors.New(`adapterResult not pb`)
	}
	clientFacade.SendMessage(message.NewMessage(1, 0, 1, 1, service.Code(msgPath), pb), false, UID, CurrentChlCtx, "")
	return nil
}
