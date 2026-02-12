package outhttpor

import (
	"github.com/goodluck0107/gona/boot/channel"
	"github.com/goodluck0107/gox/code/protocol"
	"gitlab.yq-dev-inner.com/yq-game-developer/main-server/ck-common.git/ck-logger/logger"
	"net/http"
)

func NewProtocolRawEncoder() (this *ProtocolRawEncoder) {
	this = new(ProtocolRawEncoder)
	return
}

// json ---> []byte
type ProtocolRawEncoder struct {
}

func (encoder *ProtocolRawEncoder) ExceptionCaught(ctx channel.ChannelContext, err error) {
	//	logger.Debug("MessageEncoder ExceptionCaught")
}

func (encoder *ProtocolRawEncoder) Write(ctx channel.ChannelContext, e interface{}) interface{} {
	httpWriter := ctx.ContextAttr().Get(channel.KeyForResponse).(http.ResponseWriter)
	if httpWriter != nil {
		// 	header := fmt.Sprintf("HTTP/1.0 200 OK\r\n
		//	Content-Length: %d\r\n
		//	X-Powered-By: Jetty\r\n
		//	Access-Control-Max-Age: 86400\r\n
		//	Access-Control-Allow-Credentials: true\r\n
		//	Access-Control-Allow-Origin: *\r\n
		//	Access-Control-Allow-Methods: GET,PUT,POST,GET,DELETE,OPTIONS\r\n
		//	Access-Control-Allow-Headers: Origin,X-Requested-With,Content-Type,Content-Length,Accept,Authorization,X-Request-Info\r\n
		//	Content-Type: text/plain; charset=UTF-8\r\n\r\n", len(b))
		httpWriter.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	}
	//	logger.Debug("MessageEncoder Write")
	proto := e.(protocol.Protocol)
	buf, err := proto.Encode()
	if err != nil {
		logger.Error("ProtocolEncoder Write err=", err)
		return nil
	}
	return buf
}

func (encoder *ProtocolRawEncoder) Close(ctx channel.ChannelContext) {
	//	logger.Debug("MessageEncoder Close")
}
