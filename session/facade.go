package session

func GetCount() int64 {
	return sessionPoolInstance.GetCount()
}

func RemoveSession(sessionID int64) {
	sessionPoolInstance.RemoveSession(sessionID)
}

func AddSession(session ISession) {
	sessionPoolInstance.AddSession(session)
}

func GetSession(sessionID int64) ISession {
	return sessionPoolInstance.GetSession(sessionID)
}

func TraverseDo(f func(ISession, interface{}), paramData interface{}) {
	sessionPoolInstance.TraverseDo(f, paramData)
}

var sessionPoolInstance ISessionPool = NewSessionPool()
