package extends

import (
	"fmt"

	"gitee.com/andyxt/gox/service"
	"gitee.com/andyxt/gox/session"
)

const sessionKeyChlCtx = "ChlCtx"

func SessionEquals(m session.ISession, n session.ISession) (ret bool) {
	return m.ID() == n.ID()
}

func SessionToString(s session.ISession) string {
	sessionID := s.ID()
	sessionUID := s.UID()
	chlCtx := GetChlCtx(s)
	chlCtxID := chlCtx.ID()
	chlCtxUID := UID(chlCtx)
	return fmt.Sprintf("sessionID=%s sessionUID=%d chlCtxID=%s chlCtxUID=%d ", sessionID, sessionUID,
		chlCtxID, chlCtxUID)
}

func ChangeChlCtx(s session.ISession, chlCtx service.IChannelContext) {
	s.SyncSet(sessionKeyChlCtx, chlCtx)
}

func GetChlCtx(s session.ISession) service.IChannelContext {
	if v := s.SyncGet(sessionKeyChlCtx); v != nil {
		contextValue, isContextValue := v.(service.IChannelContext) // Alt. non panicking version
		if isContextValue {
			return contextValue
		}
	}
	return nil
}
