package commands

import (
	"gitee.com/andyxt/gox/mediator/client"
)

var clientFacade *client.ClientFacade

func Init() {
	clientFacade = client.BootClient(NewCallBack())
}
