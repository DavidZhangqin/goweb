package controller

import (
	"fmt"
	"net/http"

	log "github.com/cihub/seelog"
	"github.com/gorilla/context"
)

func (*CtrlStr) Get_Site_Index(w http.ResponseWriter, r *http.Request) {
	// log.Trace("site index")
	p := context.Get(r, "params")
	log.Tracef("sitep: %v", p)
	fmt.Fprint(w, "get site/index\n")
}

func (*CtrlStr) Get_Site_Hello(w http.ResponseWriter, r *http.Request) {
	// log.Println("site hello")
	fmt.Fprint(w, "hello world!\n")
}
