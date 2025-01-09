package buffer

type ByteOrder int8

const (
	ByteOrderBigEndian    ByteOrder = 0
	ByteOrderLittleEndian ByteOrder = 1
)

const (
	int8Size     int32 = 1
	int16Size    int32 = 2
	int32Size    int32 = 4
	int64Size    int32 = 8
	stringPrefix int32 = 2
	bytesPrefix  int32 = 4
)

// CreateBigEndianBuffer 创建大端序Buffer
func CreateBigEndianBuffer() Buffer {
	return CreateBuffer(ByteOrderBigEndian)
}

// CreateLittleEndianBuffer 创建小端序Buffer
func CreateLittleEndianBuffer() Buffer {
	return CreateBuffer(ByteOrderLittleEndian)
}

// CreateBuffer 指定字节序创建Buffer
func CreateBuffer(byteOrder ByteOrder) Buffer {
	if byteOrder == ByteOrderBigEndian {
		return new(beProtocolBuffer)
	}
	return new(leProtocolBuffer)
}

func FromBytes(v []byte, byteOrder ByteOrder) Buffer {
	if byteOrder == ByteOrderBigEndian {
		ret := new(beProtocolBuffer)
		ret.content = v
		return ret
	}
	ret := new(leProtocolBuffer)
	ret.content = v
	return ret
}

type Buffer interface {
	ReadInt8() int8
	ReadUInt8() uint8
	ReadInt8WithIndex(int32) int8
	ReadUInt8WithIndex(int32) uint8
	ReadInt16() int16
	ReadInt16WithIndex(int32) int16
	ReadInt32() int32
	ReadInt32WithIndex(int32) int32
	ReadInt64() int64
	ReadInt64WithIndex(int32) int64
	ReadBytes() []byte
	ReadBytesWithOutLength() []byte
	ReadBytesWithIndex(int32) []byte
	ReadString(size int32) string
	ReadStringWithIndex(int32, int32) string
	ReadStringWithoutSize() string
	ReadStringWithOutLength() string

	WriteInt8(v int8)
	WriteUInt8(v uint8)
	WriteInt8WithIndex(int32, int8)
	WriteUInt8WithIndex(int32, uint8)
	WriteInt16(v int16)
	WriteInt16WithIndex(int32, int16)
	WriteUInt16(v uint16)
	WriteInt32(v int32)
	WriteInt32WithIndex(int32, int32)
	WriteInt64(v int64)
	WriteInt64WithIndex(int32, int64)
	WriteBytes(v []byte)
	WriteBytesWithOutLength([]byte)
	WriteBytesWithIndex(int32, []byte)
	WriteString(size int32, v string)
	WriteStringWithoutSize(v string)
	WriteStringWithIndex(int32, string)
	ToBytes() []byte
	GetWriteIndex() int32
	GetReadIndex() int32
	GetContent() []byte
	GetBodyContent() []byte
}
