package session

type ISessionPoolBuilder interface {
	GetSessionPool(int64) ISessionPool
}
