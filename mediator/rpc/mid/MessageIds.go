package mid

// dictionary is used for compressed route.
const (
	// rpc c>s ----------------------------------------------------------------
	HeartbeatRequest   uint16 = 63001
	PublishRequest     uint16 = 63002
	SubscribeRequest   uint16 = 63003
	UnsubscribeRequest uint16 = 63004
	EchoRequest        uint16 = 63005
	// rpc s>c ----------------------------------------------------------------
	HeartbeatResponse uint16 = 63101
	MessagePush       uint16 = 63102
)
