package route

import (
	"net/http"

	ct "controller"
	"lib/dav"
	"lib/session"
	"util"

	"github.com/julienschmidt/httprouter"
)

var staticBase string

func Register() *httprouter.Router {
	router := httprouter.New()

	router.GET("/public/*path", F2R(StaticHandle))

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
		router.GET("/debug/pprof/*name", pprof.Index)
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

func LoadRoute(sPath string) {
	staticBase = sPath
}

func StaticHandle(w http.ResponseWriter, r *http.Request) {
	// Disable listing directories
	// if strings.HasSuffix(r.URL.Path, "/") {
	// 	http.NotFound(w, r)
	// 	return
	// }
	var filePath string
	if staticBase == "/" {
		filePath = r.URL.Path[1:]
	} else {
		filePath = staticBase + r.URL.Path[1:]
	}
	http.ServeFile(w, r, filePath)
}
