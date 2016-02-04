package controller

import (
	// "lib/dav"
	"net/http"
	"net/http/pprof"

	"github.com/julienschmidt/httprouter"
)

type Pprof struct{}

// func (*Pprof) Index(c *dav.Context) {
// 	pprof.Index(c.W, c.R)
// }

func (*Pprof) Index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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
