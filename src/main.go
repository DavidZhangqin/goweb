package main

import (
	"net/http"
	"net/http/pprof"
	_ "net/http/pprof"
	"runtime"

	ctrl "controller"
	"util"

	log "github.com/cihub/seelog"
	"github.com/julienschmidt/httprouter"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	util.NewLogs()
}

func main() {
	defer log.Flush()
	log.Infof("app start")

	router := ctrl.RouteRegister()
	router.GET("/debug/pprof/", PprofIndex)
	router.GET("/debug/pprof/:name", Pprof)
	log.Info(http.ListenAndServe(":8089", router))
}

func PprofIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	pprof.Index(w, r)
}

func Pprof(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	switch p.ByName("name") {
	case "profile":
		pprof.Profile(w, r)
	case "symbol":
		pprof.Symbol(w, r)
	case "trace":
		pprof.Trace(w, r)
	case "cmdline":
		pprof.Cmdline(w, r)
	default:
		pprof.Index(w, r)
	}
}
