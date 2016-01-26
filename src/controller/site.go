package controller

import (
	"fmt"
	"net/http"
)

func (*CtrlStr) Get_Site_Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "get site/index")
}

func (*CtrlStr) Get_Site_Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "get site/hello")
}
