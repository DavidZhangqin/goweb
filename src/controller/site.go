package controller

import (
	"fmt"
	// "log"
	"net/http"
)

func (*CtrlStr) Get_Site_Index(w http.ResponseWriter, r *http.Request) {
	// log.Println("site index")
	fmt.Fprint(w, "get site/index\n")
}

func (*CtrlStr) Get_Site_Hello(w http.ResponseWriter, r *http.Request) {
	// log.Println("site hello")
	fmt.Fprint(w, "hello world!\n")
}
