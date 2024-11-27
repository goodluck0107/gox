package session

import (
	"errors"
	"math"
	"sync"
)

const (
	SessionMapsCount int64 = 1000 // 降低锁粒度，提高并发度
)

type SessionMap struct {
	mSessions map[int64]ISession
	mLock     *sync.Mutex
}

func NewSessionMap() (ret *SessionMap) {
	ret = new(SessionMap)
	ret.mSessions = make(map[int64]ISession)
	ret.mLock = new(sync.Mutex)
	return
}

func (sessionMap *SessionMap) add(value ISession) {
	if value == nil {
		return
	}
	defer sessionMap.mLock.Unlock()
	sessionMap.mLock.Lock()
	sessionMap.mSessions[value.UID()] = value
}

func (sessionMap *SessionMap) remove(uKey int64) {
	defer sessionMap.mLock.Unlock()
	sessionMap.mLock.Lock()
	delete(sessionMap.mSessions, uKey)
}

func (sessionMap *SessionMap) get(uKey int64) (ret ISession) {
	defer sessionMap.mLock.Unlock()
	sessionMap.mLock.Lock()
	ret = sessionMap.mSessions[uKey]
	return
}

func (sessionMap *SessionMap) TraverseDo(f func(ISession, interface{}), paramData interface{}) {
	defer sessionMap.mLock.Unlock()
	sessionMap.mLock.Lock()
	for _, session := range sessionMap.mSessions {
		f(session, paramData)
	}
}

type SessionPool struct {
	mSesionMaps      map[int64]*SessionMap
	mSessionCount    int64
	syncSessionCount *sync.Mutex
}

func NewSessionPool() (ret *SessionPool) {
	ret = new(SessionPool)
	ret.mSesionMaps = make(map[int64]*SessionMap)
	for i := int64(0); i < SessionMapsCount; i = i + 1 {
		ret.mSesionMaps[i] = NewSessionMap()
	}
	ret.syncSessionCount = new(sync.Mutex)
	return
}

func (sessionPool *SessionPool) addSession(value ISession) {
	if value == nil {
		panic(errors.New("AddSession 参数 value 不能为空"))
	}
	uKey := value.UID()
	mapsKey := int64(math.Abs(float64(uKey % SessionMapsCount)))
	if listChannels, ok := sessionPool.mSesionMaps[mapsKey]; ok {
		listChannels.add(value)
	}
	defer sessionPool.syncSessionCount.Unlock()
	sessionPool.syncSessionCount.Lock()
	sessionPool.mSessionCount = sessionPool.mSessionCount + 1
}

func (sessionPool *SessionPool) removeSession(uKey int64) {
	mapsKey := int64(math.Abs(float64(uKey % SessionMapsCount)))
	if listChannels, ok := sessionPool.mSesionMaps[mapsKey]; ok {
		listChannels.remove(uKey)
		defer sessionPool.syncSessionCount.Unlock()
		sessionPool.syncSessionCount.Lock()
		sessionPool.mSessionCount = sessionPool.mSessionCount - 1
	}
}

func (sessionPool *SessionPool) getSession(uKey int64) (ret ISession) {
	mapsKey := int64(math.Abs(float64(uKey % SessionMapsCount)))
	if listChannels, ok := sessionPool.mSesionMaps[mapsKey]; ok {
		ret = listChannels.get(uKey)
	}
	return
}

func (sessionPool *SessionPool) GetCount() int64 {
	defer sessionPool.syncSessionCount.Unlock()
	sessionPool.syncSessionCount.Lock()
	return sessionPool.mSessionCount
}

func (sessionPool *SessionPool) RemoveSession(uid int64) {
	sessionPool.removeSession(uid)
}

func (sessionPool *SessionPool) AddSession(pSession ISession) {
	sessionPool.addSession(pSession)
}

func (sessionPool *SessionPool) GetSession(uId int64) ISession {
	return sessionPool.getSession(uId)
}

func (sessionPool *SessionPool) TraverseDo(f func(ISession, interface{}), paramData interface{}) {
	for _, listChannels := range sessionPool.mSesionMaps {
		listChannels.TraverseDo(f, paramData)
	}
}
