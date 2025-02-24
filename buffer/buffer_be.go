package buffer

import (
	"errors"
	"fmt"
)

type beProtocolBuffer struct {
	content       []byte
	readPosition  int32
	writePosition int32
}

func (p *beProtocolBuffer) ReadInt8() (ret int8) {
	length := int32(len(p.content))
	if p.readPosition+int8Size > length {
		err := errors.New(fmt.Sprint("ReadInt8() 协议包读取越界，readPosition=", p.readPosition, ",length=", length))
		panic(err)
	}
	ret = int8(p.content[p.readPosition])
	p.readPosition = p.readPosition + 1
	return
}
func (p *beProtocolBuffer) ReadUInt8() (ret uint8) {
	length := int32(len(p.content))
	if p.readPosition+int8Size > length {
		err := errors.New(fmt.Sprint("ReadUInt8() 协议包读取越界，readPosition=", p.readPosition, ",length=", length))
		panic(err)
	}
	ret = uint8(p.content[p.readPosition])
	p.readPosition = p.readPosition + 1
	return
}

func (p *beProtocolBuffer) ReadInt8WithIndex(index int32) (ret int8) {
	length := int32(len(p.content))
	if index+int8Size > length {
		err := errors.New(fmt.Sprint("ReadInt8WithIndex() 协议包读取越界，readPosition=", index, ",length=", length))
		panic(err)
	}
	ret = int8(p.content[index])
	return
}

func (p *beProtocolBuffer) ReadUInt8WithIndex(index int32) (ret uint8) {
	length := int32(len(p.content))
	if index+int8Size > length {
		err := errors.New(fmt.Sprint("ReadUInt8WithIndex() 协议包读取越界，readPosition=", index, ",length=", length))
		panic(err)
	}
	ret = uint8(p.content[index])
	return
}

func (p *beProtocolBuffer) ReadInt16() (ret int16) {
	length := int32(len(p.content))
	if p.readPosition+int16Size > length {
		err := errors.New(fmt.Sprint("ReadUInt8WithIndex() 协议包读取越界，readPosition=", p.readPosition, ",length=", length))
		panic(err)
	}
	buf := p.content[p.readPosition : p.readPosition+int16Size]
	ret = ByteToInt16(buf)
	p.readPosition = p.readPosition + int16Size
	return
}
func (p *beProtocolBuffer) ReadUInt16() (ret uint16) {
	length := int32(len(p.content))
	if p.readPosition+int16Size > length {
		err := errors.New(fmt.Sprint("ReadUInt8WithIndex() 协议包读取越界，readPosition=", p.readPosition, ",length=", length))
		panic(err)
	}
	buf := p.content[p.readPosition : p.readPosition+int16Size]
	ret = ByteToUInt16(buf)
	p.readPosition = p.readPosition + int16Size
	return
}

func (p *beProtocolBuffer) ReadInt16WithIndex(index int32) (ret int16) {
	length := int32(len(p.content))
	if index+int16Size > length {
		err := errors.New(fmt.Sprint("ReadUInt8WithIndex() 协议包读取越界，readPosition=", index, ",length=", length))
		panic(err)
	}
	buf := p.content[index : index+int16Size]
	ret = ByteToInt16(buf)
	return
}

func (p *beProtocolBuffer) ReadInt32() (ret int32) {
	length := int32(len(p.content))
	if p.readPosition+int32Size > length {
		err := errors.New(fmt.Sprint("ReadUInt8WithIndex() 协议包读取越界，readPosition=", p.readPosition, ",length=", length))
		panic(err)
	}
	buf := p.content[p.readPosition : p.readPosition+int32Size]
	ret = ByteToInt32(buf)
	p.readPosition = p.readPosition + int32Size
	return
}

func (p *beProtocolBuffer) ReadInt32WithIndex(index int32) (ret int32) {
	length := int32(len(p.content))
	if index+int32Size > length {
		err := errors.New(fmt.Sprint("ReadUInt8WithIndex() 协议包读取越界，readPosition=", index, ",length=", length))
		panic(err)
	}
	buf := p.content[index : index+int32Size]
	ret = ByteToInt32(buf)
	return
}

func (p *beProtocolBuffer) ReadInt64() (ret int64) {
	length := int32(len(p.content))
	if p.readPosition+int64Size > length {
		err := errors.New(fmt.Sprint("ReadUInt8WithIndex() 协议包读取越界，readPosition=", p.readPosition, ",length=", length))
		panic(err)
	}
	buf := p.content[p.readPosition : p.readPosition+int64Size]
	ret = ByteToInt64(buf)
	p.readPosition = p.readPosition + int64Size
	return
}

func (p *beProtocolBuffer) ReadInt64WithIndex(index int32) (ret int64) {
	length := int32(len(p.content))
	if index+int64Size > length {
		err := errors.New(fmt.Sprint("ReadUInt8WithIndex() 协议包读取越界，readPosition=", index, ",length=", length))
		panic(err)
	}
	buf := p.content[index : index+int64Size]
	ret = ByteToInt64(buf)
	return
}

func (p *beProtocolBuffer) ReadBytes() (ret []byte) {
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

func (p *beProtocolBuffer) ReadBytesWithOutLength() (ret []byte) {
	length := int32(len(p.content))
	ret = p.content[p.readPosition:length]
	p.readPosition = length
	return
}

func (p *beProtocolBuffer) ReadBytesWithIndex(index int32) (ret []byte) {
	length := int32(len(p.content))
	size := p.ReadInt32WithIndex(index)
	if index+bytesPrefix+size > length {
		err := errors.New(fmt.Sprint("ReadBytesWithIndex() 协议包读取越界，readPosition=", index+bytesPrefix, ",length=", length))
		panic(err)
	}
	ret = p.content[index+bytesPrefix : index+bytesPrefix+size]
	return
}

func (p *beProtocolBuffer) ReadStringWithoutSize() (ret string) {
	length := int32(len(p.content))
	size := int32(p.ReadInt16())
	if p.readPosition+size > length {
		err := errors.New(fmt.Sprint("ReadStringWithoutSize() 协议包读取越界，readPosition=", p.readPosition, " readSize=", size, ",length=", length))
		panic(err)
	}
	bytes := p.content[p.readPosition : p.readPosition+size]
	p.readPosition = p.readPosition + size
	ret = string(bytes)
	return
}

func (p *beProtocolBuffer) ReadString(size int32) (ret string) {
	length := int32(len(p.content))
	if p.readPosition+size > length {
		err := errors.New(fmt.Sprint("ReadString() 协议包读取越界，readPosition=", p.readPosition, " readSize=", size, ",length=", length))
		panic(err)
	}
	bytes := make([]byte, 0)
	for k := p.readPosition; k < p.readPosition+size; k++ {
		if p.content[k] != 0 {
			bytes = append(bytes, p.content[k])
		} else {
			break
		}
	}
	// p.content[p.readPosition : p.readPosition+size]
	p.readPosition = p.readPosition + size
	ret = string(bytes)
	return
}

func (p *beProtocolBuffer) ReadStringWithOutLength() (ret string) {
	length := int32(len(p.content))
	bytes := p.content[p.readPosition:length]
	p.readPosition = length
	ret = string(bytes)
	return
}

func (p *beProtocolBuffer) ReadStringWithIndex(index int32, size int32) (ret string) {
	length := int32(len(p.content))
	if index+size > length {
		err := errors.New(fmt.Sprint("ReadStringWithIndex() 协议包读取越界，readPosition=", index+bytesPrefix, ",length=", length))
		panic(err)
	}
	bytes := p.content[index : index+size]
	ret = string(bytes)
	return
}

func (p *beProtocolBuffer) WriteInt8(v int8) {
	p.content = append(p.content, byte(v))
	p.writePosition = p.writePosition + int8Size
}

func (p *beProtocolBuffer) WriteUInt8(v uint8) {
	p.content = append(p.content, byte(v))
	p.writePosition = p.writePosition + int8Size
}

func (p *beProtocolBuffer) WriteInt8WithIndex(index int32, v int8) {
	p.content[index] = byte(v)
}

func (p *beProtocolBuffer) WriteUInt8WithIndex(index int32, v uint8) {
	p.content[index] = byte(v)
}

func (p *beProtocolBuffer) WriteInt16(v int16) {
	buf := Int16ToByte(v)
	p.content = append(p.content, buf...)
	p.writePosition = p.writePosition + int16Size
}
func (p *beProtocolBuffer) WriteUInt16(v uint16) {
	buf := UInt16ToByte(v)
	p.content = append(p.content, buf...)
	p.writePosition = p.writePosition + int16Size
}
func (p *beProtocolBuffer) WriteInt16WithIndex(index int32, v int16) {
	buf := Int16ToByte(v)
	for i, v := range buf {
		p.content[index+int32(i)] = v
	}
}

func (p *beProtocolBuffer) WriteInt32(v int32) {
	buf := Int32ToByte(v)
	p.content = append(p.content, buf...)
	p.writePosition = p.writePosition + int32Size
}

func (p *beProtocolBuffer) WriteInt32WithIndex(index int32, v int32) {
	buf := Int32ToByte(v)
	for i, v := range buf {
		p.content[index+int32(i)] = v
	}
}

func (p *beProtocolBuffer) WriteInt64(v int64) {
	buf := Int64ToByte(v)
	p.content = append(p.content, buf...)
	p.writePosition = p.writePosition + int64Size
}

func (p *beProtocolBuffer) WriteInt64WithIndex(index int32, v int64) {
	buf := Int64ToByte(v)
	for i, v := range buf {
		p.content[index+int32(i)] = v
	}
}

func (p *beProtocolBuffer) WriteBytes(v []byte) {
	length := int32(len(v))
	p.WriteInt32(int32(length))
	p.content = append(p.content, v...)
	p.writePosition = p.writePosition + length
}
func (p *beProtocolBuffer) WriteBytesWithOutLength(v []byte) {
	length := int32(len(v))
	p.content = append(p.content, v...)
	p.writePosition = p.writePosition + length
}
func (p *beProtocolBuffer) WriteBytesWithIndex(index int32, v []byte) {
	length := int32(len(v))
	p.WriteInt32WithIndex(index, int32(length))
	for i, v1 := range v {
		p.content[index+bytesPrefix+int32(i)] = v1
	}
}
func (p *beProtocolBuffer) WriteString(size int32, v string) {
	bytes := make([]byte, size)
	vBytes := []byte(v)
	copy(bytes, vBytes)
	p.content = append(p.content, bytes...)
	p.writePosition = p.writePosition + size
}

func (p *beProtocolBuffer) WriteStringWithoutSize(v string) {
	bytes := []byte(v)
	length := int32(len(bytes))
	p.WriteInt16(int16(length))
	p.content = append(p.content, bytes...)
	p.writePosition = p.writePosition + length
}

func (p *beProtocolBuffer) WriteStringWithIndex(index int32, v string) {
	bytes := []byte(v)
	length := int32(len(bytes))
	p.WriteInt16WithIndex(index, int16(length))
	for i, v1 := range bytes {
		p.content[index+stringPrefix+int32(i)] = v1
	}
}

func (p *beProtocolBuffer) ToBytes() (ret []byte) {
	length := int32(len(p.content))
	ret = make([]byte, 0, length)
	ret = append(ret, p.content...)
	return
}

func (p *beProtocolBuffer) GetWriteIndex() (ret int32) {
	return p.writePosition
}

func (p *beProtocolBuffer) GetReadIndex() (ret int32) {
	return p.readPosition
}
func (p *beProtocolBuffer) GetContent() (ret []byte) {
	return p.content
}

// TOFIX
func (p *beProtocolBuffer) GetBodyContent() (ret []byte) {
	totalLength := int32(len(p.content))
	length := totalLength - 8
	ret = make([]byte, length)
	for i := int32(8); i < totalLength; i = i + 1 {
		ret[i-8] = p.content[i]
	}
	return ret
}
