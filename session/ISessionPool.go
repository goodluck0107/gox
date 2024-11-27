package session

type ISessionPool interface {
	GetCount() int64

	RemoveSession(uid int64)

	AddSession(pSession ISession)

	GetSession(uId int64) ISession

	TraverseDo(f func(ISession, interface{}), paramData interface{})
}
