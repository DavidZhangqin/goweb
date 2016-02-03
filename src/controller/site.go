package controller

import (
	"fmt"
	"net/http"

	"lib/dav"

	log "github.com/cihub/seelog"
)

type Site struct{}

func (*Site) Test(c *dav.Context) {
	name := c.P.ByName("name")
	log.Tracef("hello %s", name)
	fmt.Fprintf(c.W, "hello %s", name)
}

func (*Site) Index(c *dav.Context) {
	// log.Trace("site index")
	http.Error(c.W, "it is a http error", 403)
	// fmt.Fprint(w, "get site/index\n")
}

func (*Site) Hello(c *dav.Context) {
	// log.Println("site hello")
	log.Tracef("hello %s\n", c.P.ByName("name"))
	fmt.Fprintf(c.W, "hello %s!\n", c.P.ByName("name"))
}
