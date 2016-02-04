package controller

import (
	// "fmt"
	"lib/dav"
	"lib/view"
)

type About struct{}

type Person struct {
	Name   string
	Gender string
}

func (*About) Index(c *dav.Context) {
	data := Person{"david", "man"}
	// fmt.Fprint(c.W, "get about/index\n")
	// view.Render(c.W, "about/index")
	view.RenderData(c.W, "about/index", data)
}
