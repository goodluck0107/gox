package serializier

import (
	"github.com/goodluck0107/gox/code/protocol"

	"github.com/goodluck0107/gox/buffer"
)

type IMessageFactory interface {
	/*
		@Return reuse 返回的结构ret是否可以重用，如果为true代表可以重用，则不需要调用ret的Decode方法解码，节约性能
	*/
	GetMessage(buf buffer.Buffer) (ret protocol.Protocol, reuse bool)
}
