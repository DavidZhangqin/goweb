package dav

import (
	"net/http"

	"lib/session"

	"github.com/julienschmidt/httprouter"
)

type Handle func(*Context)

type Context struct {
	W http.ResponseWriter
	R *http.Request
	P httprouter.Params
	S *session.Session
}

func NewContext(w http.ResponseWriter, r *http.Request, p httprouter.Params, s *session.Session) *Context {
	return &Context{W: w, R: r, P: p, S: s}
}
