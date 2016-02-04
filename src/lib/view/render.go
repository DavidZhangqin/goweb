package view

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var (
	funcMaps template.FuncMap
	cacheTmp map[string]*template.Template
)

type View struct {
	PostFix       string
	LayoutFolder  string
	DefaultLayout string
	Folder        string
	IsCache       bool
}

var viewConf *View

func LoadView(postFix, layoutFolder, defaultLayout, folder string, isCache bool) {
	viewConf = &View{
		PostFix:       postFix,
		LayoutFolder:  layoutFolder,
		DefaultLayout: defaultLayout,
		Folder:        folder,
		IsCache:       isCache,
	}

	cacheTmp = make(map[string]*template.Template)
	funcMaps = make(template.FuncMap)
}

func LoadPlugins(fm ...template.FuncMap) {
	// fm := make(template.FuncMap)
	for _, v := range fm {
		for m, i := range v {
			funcMaps[m] = i
		}
	}
	// log.Info(funcMaps)
	// funcMaps = fm
}

func Render(w http.ResponseWriter, name string) {
	RenderData(w, name, nil)
}

func RenderData(w http.ResponseWriter, name string, data interface{}) {
	layout := viewConf.DefaultLayout
	RenderTemplate(w, layout, name, data)
}

func RenderTemplate(w http.ResponseWriter, layout, name string, data interface{}) {
	templates := strings.Split(name, "/")
	if len(templates) != 2 {
		http.Error(w, "Wrong Render Path: "+name, http.StatusInternalServerError)
		return
	}
	var ts *template.Template
	var err error
	mainPath, err := filepath.Abs(viewConf.Folder + string(os.PathSeparator) + viewConf.LayoutFolder + string(os.PathSeparator) + layout + "." + viewConf.PostFix)
	if err != nil {
		http.Error(w, "Main File Abs Error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	renderPath, err := filepath.Abs(viewConf.Folder + string(os.PathSeparator) + name + "." + viewConf.PostFix)
	if err != nil {
		http.Error(w, "Render File Abs Error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if viewConf.IsCache {
		tmpKey := layout + "_" + name
		var ok bool
		if ts, ok = cacheTmp[tmpKey]; ok {
			ts = cacheTmp[tmpKey]
		} else {
			if ts, err = template.ParseFiles(mainPath, renderPath); err != nil {
				http.Error(w, "Template Parse Error: "+err.Error(), http.StatusInternalServerError)
				return
			}
			cacheTmp[tmpKey] = ts
		}
	} else {
		if ts, err = template.ParseFiles(mainPath, renderPath); err != nil {
			http.Error(w, "Template Parse Error: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	err = ts.Funcs(funcMaps).ExecuteTemplate(w, layout, data)
	if err != nil {
		http.Error(w, "Template Execute Error: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
