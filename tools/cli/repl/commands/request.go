package commands

import (
	"errors"
	"fmt"
	"strings"

	"github.com/goodluck0107/gox/extends"
	"github.com/goodluck0107/gox/messageImpl"

	"github.com/goodluck0107/gox/tools/cli/msgproto"

	"github.com/goodluck0107/gox/service"

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
	if msgPath != "/Login" && !extends.HasUserInfo(CurrentChlCtx) {
		return errors.New(`not HasUserInfo`)
	}
	clientFacade.SendMessage(UID, CurrentChlCtx, messageImpl.NewMessage(1, 0, 1, 1, service.Code(msgPath), pb), false, "")
	return nil
}
