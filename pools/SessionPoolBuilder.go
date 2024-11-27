package pools

import (
	"errors"
	"strconv"

	"gitee.com/andyxt/gox/session"
)

type SessionPoolBuilder struct {
	mSessionPools map[int64]session.ISessionPool
}

func NewSessionPoolBuilder() (ret *SessionPoolBuilder) {
	ret = new(SessionPoolBuilder)
	ret.mSessionPools = make(map[int64]session.ISessionPool)
	ret.Init()
	return
}

func (builder *SessionPoolBuilder) Init() {
	builder.mSessionPools[0] = session.NewSessionPool()
}

func (builder *SessionPoolBuilder) GetSessionPool(key int64) session.ISessionPool {
	if sessionPool, ok := builder.mSessionPools[key]; ok {
		return sessionPool
	}
	panic(errors.New("Invalid SessionPoolKey:" + strconv.Itoa(int(key))))
}
