package buffer

import (
	"errors"
	"fmt"
)

const (
	Int8Size     int32 = 1
	Int16Size    int32 = 2
	Int32Size    int32 = 4
	Int64Size    int32 = 8
	StringPrefix int32 = 2
	BytesPrefix  int32 = 4
)

type ProtocolBuffer struct {
	content       []byte
	readPosition  int32
	writePosition int32
}

func (p *ProtocolBuffer) ReadInt8() (ret int8) {
	length := int32(len(p.content))
	if p.readPosition+Int8Size > length {
		err := errors.New(fmt.Sprint("ReadInt8() 协议包读取越界，readPosition=", p.readPosition, ",length=", length))
		panic(err)
	}
	ret = int8(p.content[p.readPosition])
	p.readPosition = p.readPosition + 1
	return
}
func (p *ProtocolBuffer) ReadUInt8() (ret uint8) {
	length := int32(len(p.content))
	if p.readPosition+Int8Size > length {
		err := errors.New(fmt.Sprint("ReadUInt8() 协议包读取越界，readPosition=", p.readPosition, ",length=", length))
		panic(err)
	}
	ret = uint8(p.content[p.readPosition])
	p.readPosition = p.readPosition + 1
	return
}

func (p *ProtocolBuffer) ReadInt8WithIndex(index int32) (ret int8) {
	length := int32(len(p.content))
	if index+Int8Size > length {
		err := errors.New(fmt.Sprint("ReadInt8WithIndex() 协议包读取越界，readPosition=", index, ",length=", length))
		panic(err)
	}
	ret = int8(p.content[index])
	return
}

func (p *ProtocolBuffer) ReadUInt8WithIndex(index int32) (ret uint8) {
	length := int32(len(p.content))
	if index+Int8Size > length {
		err := errors.New(fmt.Sprint("ReadUInt8WithIndex() 协议包读取越界，readPosition=", index, ",length=", length))
		panic(err)
	}
	ret = uint8(p.content[index])
	return
}

func (p *ProtocolBuffer) ReadInt16() (ret int16) {
	length := int32(len(p.content))
	if p.readPosition+Int16Size > length {
		err := errors.New(fmt.Sprint("ReadUInt8WithIndex() 协议包读取越界，readPosition=", p.readPosition, ",length=", length))
		panic(err)
	}
	buf := p.content[p.readPosition : p.readPosition+Int16Size]
	ret = ByteToInt16(buf)
	p.readPosition = p.readPosition + Int16Size
	return
}

func (p *ProtocolBuffer) ReadInt16WithIndex(index int32) (ret int16) {
	length := int32(len(p.content))
	if index+Int16Size > length {
		err := errors.New(fmt.Sprint("ReadUInt8WithIndex() 协议包读取越界，readPosition=", index, ",length=", length))
		panic(err)
	}
	buf := p.content[index : index+Int16Size]
	ret = ByteToInt16(buf)
	return
}

func (p *ProtocolBuffer) ReadInt32() (ret int32) {
	length := int32(len(p.content))
	if p.readPosition+Int32Size > length {
		err := errors.New(fmt.Sprint("ReadUInt8WithIndex() 协议包读取越界，readPosition=", p.readPosition, ",length=", length))
		panic(err)
	}
	buf := p.content[p.readPosition : p.readPosition+Int32Size]
	ret = ByteToInt32(buf)
	p.readPosition = p.readPosition + Int32Size
	return
}

func (p *ProtocolBuffer) ReadInt32WithIndex(index int32) (ret int32) {
	length := int32(len(p.content))
	if index+Int32Size > length {
		err := errors.New(fmt.Sprint("ReadUInt8WithIndex() 协议包读取越界，readPosition=", index, ",length=", length))
		panic(err)
	}
	buf := p.content[index : index+Int32Size]
	ret = ByteToInt32(buf)
	return
}

func (p *ProtocolBuffer) ReadInt64() (ret int64) {
	length := int32(len(p.content))
	if p.readPosition+Int64Size > length {
		err := errors.New(fmt.Sprint("ReadUInt8WithIndex() 协议包读取越界，readPosition=", p.readPosition, ",length=", length))
		panic(err)
	}
	buf := p.content[p.readPosition : p.readPosition+Int64Size]
	ret = ByteToInt64(buf)
	p.readPosition = p.readPosition + Int64Size
	return
}

func (p *ProtocolBuffer) ReadInt64WithIndex(index int32) (ret int64) {
	length := int32(len(p.content))
	if index+Int64Size > length {
		err := errors.New(fmt.Sprint("ReadUInt8WithIndex() 协议包读取越界，readPosition=", index, ",length=", length))
		panic(err)
	}
	buf := p.content[index : index+Int64Size]
	ret = ByteToInt64(buf)
	return
}

func (p *ProtocolBuffer) ReadBytes() (ret []byte) {
	length := int32(len(p.content))
	size := p.ReadInt32()
	if p.readPosition+size > length {
		err := errors.New(fmt.Sprint("ReadBytes() 协议包读取越界，readPosition=", p.readPosition, ",length=", length))
		panic(err)
	}
	ret = p.content[p.readPosition : p.readPosition+size]
	p.readPosition = p.readPosition + size
	return
}

func (p *ProtocolBuffer) ReadBytesWithOutLength() (ret []byte) {
	length := int32(len(p.content))
	ret = p.content[p.readPosition:length]
	p.readPosition = length
	return
}

func (p *ProtocolBuffer) ReadBytesWithIndex(index int32) (ret []byte) {
	length := int32(len(p.content))
	size := p.ReadInt32WithIndex(index)
	if index+BytesPrefix+size > length {
		err := errors.New(fmt.Sprint("ReadBytesWithIndex() 协议包读取越界，readPosition=", index+BytesPrefix, ",length=", length))
		panic(err)
	}
	ret = p.content[index+BytesPrefix : index+BytesPrefix+size]
	return
}

func (p *ProtocolBuffer) ReadString() (ret string) {
	length := int32(len(p.content))
	size := int32(p.ReadInt16())
	if p.readPosition+size > length {
		err := errors.New(fmt.Sprint("ReadString() 协议包读取越界，readPosition=", p.readPosition," readSize=",size, ",length=", length))
		panic(err)
	}
	bytes := p.content[p.readPosition : p.readPosition+size]
	p.readPosition = p.readPosition + size
	ret = string(bytes)
	return
}

func (p *ProtocolBuffer) ReadStringWithOutLength() (ret string) {
	length := int32(len(p.content))
	bytes := p.content[p.readPosition:length]
	p.readPosition = length
	ret = string(bytes)
	return
}

func (p *ProtocolBuffer) ReadStringWithIndex(index int32) (ret string) {
	length := int32(len(p.content))
	size := int32(p.ReadInt16WithIndex(index))
	if index+StringPrefix+size > length {
		err := errors.New(fmt.Sprint("ReadStringWithIndex() 协议包读取越界，readPosition=", index+BytesPrefix, ",length=", length))
		panic(err)
	}
	bytes := p.content[index+StringPrefix : index+StringPrefix+size]
	ret = string(bytes)
	return
}

func (p *ProtocolBuffer) WriteInt8(v int8) {
	p.content = append(p.content, byte(v))
	p.writePosition = p.writePosition + Int8Size
}

func (p *ProtocolBuffer) WriteUInt8(v uint8) {
	p.content = append(p.content, byte(v))
	p.writePosition = p.writePosition + Int8Size
}

func (p *ProtocolBuffer) WriteInt8WithIndex(index int32, v int8) {
	p.content[index] = byte(v)
}

func (p *ProtocolBuffer) WriteUInt8WithIndex(index int32, v uint8) {
	p.content[index] = byte(v)
}

func (p *ProtocolBuffer) WriteInt16(v int16) {
	buf := Int16ToByte(v)
	p.content = append(p.content, buf...)
	p.writePosition = p.writePosition + Int16Size
}
func (p *ProtocolBuffer) WriteInt16WithIndex(index int32, v int16) {
	buf := Int16ToByte(v)
	for i, v := range buf {
		p.content[index+int32(i)] = v
	}
}

func (p *ProtocolBuffer) WriteInt32(v int32) {
	buf := Int32ToByte(v)
	p.content = append(p.content, buf...)
	p.writePosition = p.writePosition + Int32Size
}

func (p *ProtocolBuffer) WriteInt32WithIndex(index int32, v int32) {
	buf := Int32ToByte(v)
	for i, v := range buf {
		p.content[index+int32(i)] = v
	}
}

func (p *ProtocolBuffer) WriteInt64(v int64) {
	buf := Int64ToByte(v)
	p.content = append(p.content, buf...)
	p.writePosition = p.writePosition + Int64Size
}

func (p *ProtocolBuffer) WriteInt64WithIndex(index int32, v int64) {
	buf := Int64ToByte(v)
	for i, v := range buf {
		p.content[index+int32(i)] = v
	}
}

func (p *ProtocolBuffer) WriteBytes(v []byte) {
	length := int32(len(v))
	p.WriteInt32(int32(length))
	p.content = append(p.content, v...)
	p.writePosition = p.writePosition + length
}
func (p *ProtocolBuffer) WriteBytesWithOutLength(v []byte) {
	length := int32(len(v))
	p.content = append(p.content, v...)
	p.writePosition = p.writePosition + length
}
func (p *ProtocolBuffer) WriteBytesWithIndex(index int32, v []byte) {
	length := int32(len(v))
	p.WriteInt32WithIndex(index, int32(length))
	for i, v1 := range v {
		p.content[index+BytesPrefix+int32(i)] = v1
	}
}

func (p *ProtocolBuffer) WriteString(v string) {
	bytes := []byte(v)
	length := int32(len(bytes))
	p.WriteInt16(int16(length))
	p.content = append(p.content, bytes...)
	p.writePosition = p.writePosition + length
}

func (p *ProtocolBuffer) WriteStringWithIndex(index int32, v string) {
	bytes := []byte(v)
	length := int32(len(bytes))
	p.WriteInt16WithIndex(index, int16(length))
	for i, v1 := range bytes {
		p.content[index+StringPrefix+int32(i)] = v1
	}
}

func (p *ProtocolBuffer) ToBytes() (ret []byte) {
	length := int32(len(p.content))
	ret = make([]byte, 0, length)
	ret = append(ret, p.content...)
	return
}

func FromBytes(v []byte) (ret *ProtocolBuffer) {
	ret = new(ProtocolBuffer)
	ret.content = v
	return
}

func (p *ProtocolBuffer) GetWriteIndex() (ret int32) {
	return p.writePosition
}

func (p *ProtocolBuffer) GetReadIndex() (ret int32) {
	return p.readPosition
}
func (p *ProtocolBuffer) GetContent() (ret []byte) {
	return p.content
}

//TOFIX 
func (p *ProtocolBuffer) GetBodyContent() (ret []byte) {
	totalLength := int32(len(p.content))
	length := totalLength - 8
	ret = make([]byte, length, length)
	for i := int32(8); i < totalLength; i = i + 1 {
		ret[i-8] = p.content[i]
	}
	return ret
}
