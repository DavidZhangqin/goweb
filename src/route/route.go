package route

import (
	"net/http"

	ct "controller"
	"lib/dav"
	"lib/session"
	"util"

	"github.com/julienschmidt/httprouter"
)

func Register() *httprouter.Router {
	router := httprouter.New()

	about := &ct.About{}
	router.GET("/about", Bind(about.Index))

	site := &ct.Site{}
	router.GET("/site", Bind(site.Index))
	router.GET("/site/hello/:name", Bind(site.Hello))
	router.POST("/site/hello/:name", Bind(site.Hello))
	router.GET("/site/test/:name", Bind(site.Test))

	// debug pprof switch
	if util.IsDebug {
		pprof := &ct.Pprof{}
		router.GET("/debug/pprof/", Bind(pprof.Index))
		router.GET("/debug/pprof/:name", Bind(pprof.Cont))
	}
	return router
}

func Bind(h dav.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		s := session.Start(w, r)
		// auth TODO
		c := dav.NewContext(w, r, p, s)
		h(c)
	}
}

// middle handle
func MiddleHandle(h http.Handler) http.Handler {
	h = LogRequest(h)
	return h
}
