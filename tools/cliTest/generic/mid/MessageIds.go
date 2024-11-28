package mid

// dictionary is used for compressed route.
const (
	// account c>s ----------------------------------------------------------------
	LoginRequest        uint16 = 101
	LogoutRequest       uint16 = 102
	HeartbeatRequest    uint16 = 103
	InactiveRequest     uint16 = 104
	KickOfflineRequest  uint16 = 105
	ForbidPlayerRequest uint16 = 106

	// account s>c
	LoginResponse     uint16 = 201
	LogoutResponse    uint16 = 202
	HeartbeatResponse uint16 = 204
	ActionFailed      uint16 = 208
	LoginConflictPush uint16 = 209
)
