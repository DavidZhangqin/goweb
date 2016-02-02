package main

import (
	"lib/session"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"runtime"

	// ctrl "controller"
	"route"
	"util"

	log "github.com/cihub/seelog"
	"github.com/julienschmidt/httprouter"
)

var config map[string]string

func init() {

	runtime.GOMAXPROCS(runtime.NumCPU())
	// log init
	util.NewLogs()

}

func main() {
	defer log.Flush()
	log.Infof("app start")

	// load config
	config = util.LoadConfig()
	log.Info(config)
	// session init
	session.LoadSession(config["session.name"], config["session.maxAge"])

	go route.Register()
	// router := ctrl.RouteRegister()
	// router := route.Register()
	// router.GET("/debug/pprof/", PprofIndex)
	// router.GET("/debug/pprof/:name", Pprof)
	// go func() {
	// 	log.Info("listen and serve 8089")
	// 	log.Info(http.ListenAndServe(":8089", router))
	// }()

	sigChan := make(chan int)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, os.Kill)
		for {
			select {
			case <-c:
				log.Info("App quit by signal")
				sigChan <- 1
			case <-util.ExitChan:
				log.Info("App quit manually")
				sigChan <- 1
			}
		}
	}()
	<-sigChan
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
