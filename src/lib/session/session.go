package session

import (
	"net/http"
	"os"
	"strconv"

	log "github.com/cihub/seelog"
)

type Session struct {
	Store
}

func LoadSession(name, maxAge string) {
	maxAgeInt, err := strconv.Atoi(maxAge)
	if err != nil {
		log.Criticalf("load session maxAge[%s] type error: %s", maxAge, err)
		os.Exit(2)
	}
	CacheInit(name, maxAgeInt)
}

func Start(w http.ResponseWriter, r *http.Request) *Session {
	sess := &Session{
		Store: NewCacheSession(w, r),
	}
	return sess
}
