package buffer

import (
	"errors"
	"fmt"
)

func ByteToInt16(buf []byte) (ret int16) {
	length := int32(len(buf))
	if length < Int16Size {
		err := errors.New(fmt.Sprint("ByteToInt16() 越界，length=", length))
		panic(err)
	}
	ret = int16(buf[0])<<8|int16(buf[1])
	return
}
func ByteToInt32(buf []byte) (ret int32) {
	length := int32(len(buf))
	if length < Int32Size {
		err := errors.New(fmt.Sprint("ByteToInt32() 越界，length=", length))
		panic(err)
	}
	ret = int32(buf[0])<<24|int32(buf[1])<<16|int32(buf[2])<<8|int32(buf[3])
	return
}
func ByteToInt64(buf []byte) (ret int64) {
	length := int32(len(buf))
	if length < Int64Size {
		err := errors.New(fmt.Sprint("ByteToInt64() 越界，length=", length))
		panic(err)
	}
	ret = 0
	ret = int64(buf[0])<<56|int64(buf[1])<<48|int64(buf[2])<<40|int64(buf[3])<<32|int64(buf[4])<<24|int64(buf[5])<<16|int64(buf[6])<<8|int64(buf[7])
	return
}

func Int16ToByte(v int16) (buf []byte) {
	buf = make([]byte, Int16Size)
	buf[0] = byte(v >> 8)
	buf[1] = byte(v)
	return buf
}
func Int32ToByte(v int32) (buf []byte) {
	buf = make([]byte, Int32Size)
	buf[0] = byte(v >> 24)
	buf[1] = byte(v >> 16)
	buf[2] = byte(v >> 8)
	buf[3] = byte(v)
	return buf
}
func Int64ToByte(v int64) (buf []byte) {
	buf = make([]byte, Int64Size)
	buf[0] = byte(v >> 56)
	buf[1] = byte(v >> 48)
	buf[2] = byte(v >> 40)
	buf[3] = byte(v >> 32)
	buf[4] = byte(v >> 24)
	buf[5] = byte(v >> 16)
	buf[6] = byte(v >> 8)
	buf[7] = byte(v)
	return buf
}
