package session

import "sync"

const sessionKeyID = "txxiongtao-sessionId"
const sessionKeyUID = "txxiongtao-sessionUId"

type Session struct {
	lock     *sync.Mutex // 扯犊子玩意，同协程都不可重入
	syncAttr map[string]interface{}
	values   map[string]interface{}
}

func NewSession(sessionID string, sessionUID int64) (s *Session) {
	s = new(Session)
	s.lock = new(sync.Mutex)
	s.syncAttr = make(map[string]interface{})
	s.values = make(map[string]interface{})
	s.SyncSet(sessionKeyID, sessionID)
	s.SyncSet(sessionKeyUID, sessionUID)
	return
}

func (s *Session) ID() string {
	if v := s.SyncGet(sessionKeyID); v != nil {
		stringValue, isStringValue := v.(string) // Alt. non panicking version
		if isStringValue {
			return stringValue
		}
	}
	return ""
}

func (s *Session) UID() int64 {
	if v := s.SyncGet(sessionKeyUID); v != nil {
		intValue, isIntValue := v.(int64) // Alt. non panicking version
		if isIntValue {
			return intValue
		}
	}
	return -1
}

func (s *Session) SyncSet(key string, value interface{}) {
	defer s.lock.Unlock()
	s.lock.Lock()
	s.syncAttr[key] = value
}

func (s *Session) SyncGet(key string) (value interface{}) {
	defer s.lock.Unlock()
	s.lock.Lock()
	if v, ok := s.syncAttr[key]; ok {
		return v
	}
	return nil
}

func (s *Session) SyncRemove(key string) {
	defer s.lock.Unlock()
	s.lock.Lock()
	delete(s.syncAttr, key)
}

func (s *Session) Set(key string, value interface{}) {
	s.values[key] = value
}

func (s *Session) Get(key string) interface{} {
	return s.values[key]
}

func (s *Session) Remove(key string) {
	delete(s.values, key)
}
