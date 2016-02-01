package session

import (
	"net/http"
)

type Session struct {
	Store
}

func LoadSession() {
	CacheInit("GOSESSID", 28800)
}

func Start(w http.ResponseWriter, r *http.Request) *Session {
	sess := &Session{
		CacheStore: NewCacheSession(w, r),
	}
	return sess
}
