package serializier

import (
	"gitee.com/andyxt/gox/code/protocol"

	"gitee.com/andyxt/gona/logger"
	"gitee.com/andyxt/gox/buffer"
)

const (
	CommonSerilizeType int8 = 1
)

type Serializier interface {
	Serialize(protocol.Protocol) []byte
	Deserialize(b []byte) (bool, protocol.Protocol)
}

type DefualtSerializier struct {
	mFactory IMessageFactory
}

func NewDefualtSerializier(mFactory IMessageFactory) (this *DefualtSerializier) {
	this = new(DefualtSerializier)
	this.mFactory = mFactory
	return this
}

func (serializier *DefualtSerializier) Serialize(b protocol.Protocol) []byte {
	pBuf, pE := b.Encode()
	if pE != nil {
		logger.Error("序列化失败")
	}
	return pBuf
}

func (serializier *DefualtSerializier) Deserialize(b []byte) (bool, protocol.Protocol) {
	//logger.Error("MessageDecoder:",byteSlice)
	buf := buffer.FromBytes(b, buffer.ByteOrderBigEndian)
	//VersionId, UserId, AppId, MessageId := protocolDefine.GetHeadFiledValue(buf)
	msg, reuse := serializier.mFactory.GetMessage(buf)
	if msg != nil {
		if reuse {
			return true, msg
		} else {
			valid := msg.Decode(b)
			if valid == nil {
				return true, msg
			} else {
				logger.Info("无效协议") //, " , 协议号：", VersionId, UserId, AppId, MessageId)
				return false, nil
			}
		}
	} else {
		logger.Info("非法协议") //, " , 协议号：", VersionId, UserId, AppId, MessageId)
		return false, nil
	}
}
