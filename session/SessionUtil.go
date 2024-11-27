package session

type SessionUtil struct {
	sessionPoolBuilder ISessionPoolBuilder
}

func newSessionUtil(sessionPoolBuilder ISessionPoolBuilder) (this *SessionUtil) {
	this = new(SessionUtil)
	this.sessionPoolBuilder = sessionPoolBuilder
	return
}

func (sessionUtil *SessionUtil) GetCount(poolID int64) int64 {
	sessionPool := sessionUtil.sessionPoolBuilder.GetSessionPool(poolID)
	return sessionPool.GetCount()
}

func (sessionUtil *SessionUtil) RemoveSession(poolID int64, sessionID int64) {
	sessionPool := sessionUtil.sessionPoolBuilder.GetSessionPool(poolID)
	sessionPool.RemoveSession(sessionID)
}

func (sessionUtil *SessionUtil) AddSession(poolID int64, session ISession) {
	sessionPool := sessionUtil.sessionPoolBuilder.GetSessionPool(poolID)
	sessionPool.AddSession(session)
}

func (sessionUtil *SessionUtil) GetSession(poolID int64, sessionID int64) ISession {
	sessionPool := sessionUtil.sessionPoolBuilder.GetSessionPool(poolID)
	return sessionPool.GetSession(sessionID)
}

func (sessionUtil *SessionUtil) TraverseDo(poolID int64, f func(ISession, interface{}), paramData interface{}) {
	sessionPool := sessionUtil.sessionPoolBuilder.GetSessionPool(poolID)
	sessionPool.TraverseDo(f, paramData)
}

var sessionUtilInstance *SessionUtil

func Init(sessionPoolBuilder ISessionPoolBuilder) {
	sessionUtilInstance = newSessionUtil(sessionPoolBuilder)
}

func GetCount(poolID int64) int64 {
	return sessionUtilInstance.GetCount(poolID)
}

func RemoveSession(poolID int64, sessionID int64) {
	sessionUtilInstance.RemoveSession(poolID, sessionID)
}

func AddSession(poolID int64, session ISession) {
	sessionUtilInstance.AddSession(poolID, session)
}

func GetSession(poolID int64, sessionID int64) ISession {
	return sessionUtilInstance.GetSession(poolID, sessionID)
}

func TraverseDo(poolID int64, f func(ISession, interface{}), paramData interface{}) {
	sessionUtilInstance.TraverseDo(poolID, f, paramData)
}
