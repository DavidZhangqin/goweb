package view

import (
	"html/template"
	"net/http"
	"strings"
)

type View struct {
	PostFix  string
	Layout   string
	Folder   string
	IsCache  bool
	CacheTmp map[string]*template.Template
}

var viewConf *View

func LoadView(postFix, layout, folder string, isCache bool) {
	viewConf = &View{
		PostFix:  postFix,
		Layout:   layout,
		Folder:   folder,
		IsCache:  isCache,
		CacheTmp: make(map[string]*template.Template),
	}
}

func Render(w http.ResponseWriter, name string) {
	RenderData(w, name, nil)
}

func RenderData(w http.ResponseWriter, name string, data interface{}) {
	templates := strings.Split(name, "/")
	if len(templates)
}
