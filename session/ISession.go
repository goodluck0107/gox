package session

type ISession interface {
	ID() string
	UID() int64
	SyncSet(key string, value interface{})
	SyncGet(key string) (value interface{})
	SyncRemove(key string)
	Set(key string, value interface{})
	Get(key string) interface{}
	Remove(key string)
}
