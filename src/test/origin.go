package main

import (
	"log"
	"net/http"
	"runtime"

	ctrl "controller"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	var actions *ctrl.CtrlStr
	http.HandleFunc("/site/hello", actions.Get_Site_Hello)
	http.HandleFunc("/site/index", actions.Get_Site_Index)
	http.HandleFunc("/about/index", actions.Get_About_Index)

	log.Println("ListenAndServe :8089")
	log.Fatal(http.ListenAndServe(":8089", nil))
}
