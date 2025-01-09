package message

import (
	"fmt"
	"time"

	"gitee.com/andyxt/gox/buffer"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoiface"
)

type Type byte

// Message represents a unmarshaled message or a message which to be marshaled
type Message struct {
	RouteType   Type   // 路由类型: 0-close,1-client,2-hall
	MessageType Type   // 协议类型: 1-proto,2-json
	Verion      uint32 // 协议版本
	SeqID       uint32 // 消息唯一ID,从0开始递增
	Time        int64  // 时间戳
	MsgID       uint16 // 协议ID
	Data        []byte // 协议数据
	ProtoData   protoiface.MessageV1
}

func NewMessageDirect(routeType Type, messageType Type,
	verion uint32, seqID uint32, msgID uint16,
	protoData []byte) (this *Message) {
	this = new(Message)
	this.RouteType = routeType
	this.MessageType = messageType
	this.Verion = verion
	this.SeqID = seqID
	this.MsgID = msgID
	this.Data = protoData
	return
}
func NewMessage(routeType Type, messageType Type,
	verion uint32, seqID uint32, msgID uint16,
	protoData protoreflect.ProtoMessage) (this *Message) {
	this = new(Message)
	this.RouteType = routeType
	this.MessageType = messageType
	this.Verion = verion
	this.SeqID = seqID
	this.MsgID = msgID
	byteArr, err := proto.Marshal(protoData)
	if err != nil {
		panic(err)
	}
	this.Data = byteArr
	return
}
func NewMessageWith(data []byte) (this *Message) {
	this = new(Message)
	protocolBuffer := buffer.FromBytes(data, buffer.ByteOrderBigEndian)
	this.RouteType = Type(protocolBuffer.ReadUInt8())   // 路由类型（0-客户端，1-大厅）
	this.MessageType = Type(protocolBuffer.ReadUInt8()) // 消息类型（0-close,1-proto）
	this.Verion = uint32(protocolBuffer.ReadInt32())    // 协议版本
	this.SeqID = uint32(protocolBuffer.ReadInt32())     // 协议序列号
	protocolBuffer.ReadInt64()                          // 时间戳
	this.MsgID = uint16(protocolBuffer.ReadInt16())     // 协议ID
	this.Data = data[protocolBuffer.GetReadIndex():]
	return
}
func (msg *Message) Decode(e interface{}) (valid bool) {
	return
}
func (msg *Message) Encode() (ret interface{}) {
	protocolBuffer := buffer.CreateBigEndianBuffer()
	protocolBuffer.WriteUInt8(uint8(msg.RouteType))   // 路由类型（0-客户端，1-大厅）
	protocolBuffer.WriteUInt8(uint8(msg.MessageType)) // 消息类型（0-close,1-proto）
	protocolBuffer.WriteInt32(int32(msg.Verion))      // 协议版本
	protocolBuffer.WriteInt32(int32(msg.SeqID))       // 协议序列号
	protocolBuffer.WriteInt64(time.Now().Unix())      // 时间戳
	protocolBuffer.WriteInt16(int16(msg.MsgID))       // 协议ID
	data := protocolBuffer.GetContent()[:protocolBuffer.GetWriteIndex()]
	data = append(data, msg.Data...)
	return data
}
func (msg *Message) GetSerializeType() int8 {
	return 0
}
func (msg *Message) GetSecurityType() int8 {
	return 0
}

func (msg *Message) String() string {
	return fmt.Sprintf("Message RouteType=%+v,MessageType=%+v,ShortVer=%+v,SeqID=%+v,MsgID=%+v,Data=%+v,",
		msg.RouteType, msg.MessageType, msg.Verion, msg.SeqID, msg.MsgID, msg.Data)
}
