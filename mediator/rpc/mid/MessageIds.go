package mid

// dictionary is used for compressed route.
const (
	// rpc g>h ----------------------------------------------------------------
	RPCLoginRequest     uint16 = 63001
	RPCLogoutRequest    uint16 = 63002
	RPCHeartbeatRequest uint16 = 63003
	RPCCallRequest      uint16 = 63004
	RPCBroadcastRequest uint16 = 63005
	EchoRequest         uint16 = 63006
	// rpc h>g ----------------------------------------------------------------
	RPCLoginResponse     uint16 = 63101
	RPCLogoutResponse    uint16 = 63102
	RPCHeartbeatResponse uint16 = 63103
	RPCLoginConflictPush uint16 = 63104
	RPCCallPush          uint16 = 63105
)
