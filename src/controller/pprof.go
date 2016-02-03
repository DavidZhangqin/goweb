package controller

import (
	"lib/dav"
	"net/http/pprof"
)

type Pprof struct{}

func (*Pprof) Index(c *dav.Context) {
	pprof.Index(c.W, c.R)
}

func (*Pprof) Cont(c *dav.Context) {
	switch c.P.ByName("name") {
	case "profile":
		pprof.Profile(c.W, c.R)
	case "symbol":
		pprof.Symbol(c.W, c.R)
	case "trace":
		pprof.Trace(c.W, c.R)
	case "cmdline":
		pprof.Cmdline(c.W, c.R)
	default:
		pprof.Index(c.W, c.R)
	}
}
