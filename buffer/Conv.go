package buffer

import (
	"errors"
	"strconv"
)

const (
	INT8_SIZE     int32 = 1
	INT16_SIZE    int32 = 2
	INT32_SIZE    int32 = 4
	INT64_SIZE    int32 = 8
	STRING_PREFIX int32 = 2
	BYTES_PREFIX  int32 = 2
)

func ByteToInt16(buf []byte) (ret int16) {
	length := int32(len(buf))
	if length < INT16_SIZE {
		err := errors.New("ByteToInt16() 越界,length=" + strconv.Itoa(int(length)))
		panic(err)
	}
	ret = int16(buf[0])<<8 | int16(buf[1])
	return
}
func ByteToInt32(buf []byte) (ret int32) {
	length := int32(len(buf))
	if length < INT32_SIZE {
		err := errors.New("ByteToInt32() 越界,length=" + strconv.Itoa(int(length)))
		panic(err)
	}
	ret = int32(buf[0])<<24 | int32(buf[1])<<16 | int32(buf[2])<<8 | int32(buf[3])
	return
}
func ByteToInt64(buf []byte) (ret int64) {
	length := int32(len(buf))
	if length < INT64_SIZE {
		err := errors.New("ByteToInt64() 越界,length=" + strconv.Itoa(int(length)))
		panic(err)
	}
	ret = 0
	for i, v := range buf {
		ret |= int64(v) << uint((7-i)*8)
	}
	return
}

func Int16ToByte(v int16) (buf []byte) {
	buf = make([]byte, INT16_SIZE)
	buf[0] = byte(v >> 8)
	buf[1] = byte(v)
	return buf
}
func Int32ToByte(v int32) (buf []byte) {
	buf = make([]byte, INT32_SIZE)
	buf[0] = byte(v >> 24)
	buf[1] = byte(v >> 16)
	buf[2] = byte(v >> 8)
	buf[3] = byte(v)
	return buf
}
func Int64ToByte(v int64) (buf []byte) {
	buf = make([]byte, INT64_SIZE)
	for i := range buf {
		buf[i] = byte(v >> uint((7-i)*8))
	}
	return buf
}
