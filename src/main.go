package main

import (
	"net/http"
	"os"
	"os/signal"
	"runtime"

	"lib/session"
	"lib/view"
	"lib/view/plugin"
	"route"
	"util"

	log "github.com/cihub/seelog"
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
	log.Debug(config)
	// session init
	session.LoadSession(config["session.name"], config["session.maxAge"])
	// route init
	route.LoadRoute(config["static.base"])
	router := route.Register()
	// view init
	var isCache bool
	if config["view.isCache"] == "true" {
		isCache = true
	} else {
		isCache = false
	}
	view.LoadView(config["view.postFix"], config["view.layoutFolder"], config["view.defaultLayout"],
		config["view.folder"], isCache)
	view.LoadPlugins(
		plugin.NoEscape(),
	)

	go func() {
		log.Info("listen and serve 8089")
		log.Info(http.ListenAndServe(":8089", route.MiddleHandle(router)))
	}()

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
