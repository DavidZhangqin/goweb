package controller

import (
	"fmt"
	"lib/dav"
)

type About struct{}

func (*About) Index(c *dav.Context) {
	fmt.Fprint(c.W, "get about/index\n")
}
