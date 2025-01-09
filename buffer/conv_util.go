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

// ByteToInt16LD 小端字节序字节数组转int16
func ByteToInt16LD(buf []byte) (ret int16) {
	length := int32(len(buf))
	if length < INT16_SIZE {
		err := errors.New("ByteToInt16LD() 越界,length=" + strconv.Itoa(int(length)))
		panic(err)
	}
	ret = int16(buf[1])<<8 | int16(buf[0])
	return
}

// ByteToInt16 大端字节序字节数组转int16
func ByteToInt16(buf []byte) (ret int16) {
	length := int32(len(buf))
	if length < INT16_SIZE {
		err := errors.New("ByteToInt16() 越界,length=" + strconv.Itoa(int(length)))
		panic(err)
	}
	ret = int16(buf[0])<<8 | int16(buf[1])
	return
}

// ByteToInt32LD 小端字节序字节数组转int32
func ByteToInt32LD(buf []byte) (ret int32) {
	length := int32(len(buf))
	if length < INT32_SIZE {
		err := errors.New("ByteToInt32LD() 越界,length=" + strconv.Itoa(int(length)))
		panic(err)
	}
	ret = int32(buf[3])<<24 | int32(buf[2])<<16 | int32(buf[1])<<8 | int32(buf[0])
	return
}

// ByteToInt32 大端字节序字节数组转int32
func ByteToInt32(buf []byte) (ret int32) {
	length := int32(len(buf))
	if length < INT32_SIZE {
		err := errors.New("ByteToInt32() 越界,length=" + strconv.Itoa(int(length)))
		panic(err)
	}
	ret = int32(buf[0])<<24 | int32(buf[1])<<16 | int32(buf[2])<<8 | int32(buf[3])
	return
}

// ByteToInt64LD 小端字节序字节数组转int64
func ByteToInt64LD(buf []byte) (ret int64) {
	length := int32(len(buf))
	if length < INT64_SIZE {
		err := errors.New("ByteToInt64() 越界,length=" + strconv.Itoa(int(length)))
		panic(err)
	}
	ret = 0
	for i, v := range buf {
		ret |= int64(v) << uint((i)*8)
	}
	return
}

// ByteToInt64 大端字节序字节数组转int64
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

// Int16ToByteLD 小端字节序int16转字节数组
func Int16ToByteLD(v int16) (buf []byte) {
	buf = make([]byte, INT16_SIZE)
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	return buf
}

// UInt16ToByteLD 小端字节序uint16转字节数组
func UInt16ToByteLD(v uint16) (buf []byte) {
	buf = make([]byte, INT16_SIZE)
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	return buf
}

// Int16ToByte 大端字节序int16转字节数组
func Int16ToByte(v int16) (buf []byte) {
	buf = make([]byte, INT16_SIZE)
	buf[0] = byte(v >> 8)
	buf[1] = byte(v)
	return buf
}

// Int32ToByteLD 小端字节序int32转字节数组
func Int32ToByteLD(v int32) (buf []byte) {
	buf = make([]byte, INT32_SIZE)
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v >> 16)
	buf[3] = byte(v >> 24)
	return buf
}

// Int32ToByte 大端字节序int32转字节数组
func Int32ToByte(v int32) (buf []byte) {
	buf = make([]byte, INT32_SIZE)
	buf[0] = byte(v >> 24)
	buf[1] = byte(v >> 16)
	buf[2] = byte(v >> 8)
	buf[3] = byte(v)
	return buf
}

// Int64ToByteLD 小端字节序int64转字节数组
func Int64ToByteLD(v int64) (buf []byte) {
	buf = make([]byte, INT64_SIZE)
	// buf[0] = byte(v)
	// buf[1] = byte(v >> 8)
	// buf[2] = byte(v >> 16)
	// buf[3] = byte(v >> 24)
	// buf[4] = byte(v >> 32)
	// buf[5] = byte(v >> 40)
	// buf[6] = byte(v >> 48)
	// buf[7] = byte(v >> 56)
	for i := range buf {
		buf[i] = byte(v >> uint(i*8))
	}
	return buf
}

// Int64ToByte 大端字节序int64转字节数组
func Int64ToByte(v int64) (buf []byte) {
	buf = make([]byte, INT64_SIZE)
	for i := range buf {
		buf[i] = byte(v >> uint((7-i)*8))
	}
	return buf
}
