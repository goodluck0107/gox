package protocol

type IProtocol interface {
	Decode(e interface{}) (valid bool)
	Encode() (ret interface{})
	GetSerializeType() int8
	GetSecurityType() int8
	// GetMessageId() int16
	// SetMessageId(MessageId int16)

	// GetInt8(paramKey int8) int8
	// GetInt16(paramKey int8) int16
	// GetInt32(paramKey int8) int32
	// GetInt64(paramKey int8) int64
	// GetString(paramKey int8) string
}
