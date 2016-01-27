package controller

import (
	"fmt"
	"net/http"
)

func (*CtrlStr) Get_About_Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "get about/index\n")
}
