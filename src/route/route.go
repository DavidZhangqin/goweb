package route

import (
	"fmt"
	"net/http"

	log "github.com/cihub/seelog"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
)

func Register() {
	router := httprouter.New()

	router.GET("/test", Test)
	router.GET("/hello/:name", Hello)

	log.Info("listen and serve 8089")
	log.Info(http.ListenAndServe(":8089", middleHandle(router)))
}

func Test(w http.ResponseWriter, r *http.Request) {
	log.Info("regist ", r.Method)
	fmt.Fprint(w, "register")
}
func Hello(w http.ResponseWriter, r *http.Request) {
	log.Info("hello ", r.Method)
	fmt.Fprint(w, "register")
}

func (r *httprouter.Router) Regist(method, path string, h http.HandleFunc) httprouter.Handle {
	h = func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		context.Set(r, "params", p)
		h.ServeHTTP(w, r)
	}
	r.Handle(method, path, h)
	return
}

// middle handle
func middleHandle(h http.Handler) http.Handler {
	h = context.ClearHandler(h)
	h = LogRequest(h)
	return h
}
