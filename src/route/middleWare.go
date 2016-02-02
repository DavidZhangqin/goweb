package route

import (
	"net/http"

	log "github.com/cihub/seelog"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
)

func LogRequest(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Debugf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		h.ServeHTTP(w, r)
	})
}
