package controller

import (
// "net/http"
// "reflect"
// "strings"

// log "github.com/cihub/seelog"
// "github.com/gorilla/context"
// "github.com/julienschmidt/httprouter"
)

// type CtrlStr struct{}

// func RouteRegister() *httprouter.Router {
// 	router := httprouter.New()

// 	object := reflect.ValueOf(&CtrlStr{})
// 	objectType := object.Type()

// 	validMethos := map[string]string{
// 		"get":    "GET",
// 		"post":   "POST",
// 		"put":    "PUT",
// 		"delete": "DELETE",
// 		"patch":  "PATCH",
// 	}
// 	for i := 0; i < object.NumMethod(); i++ {
// 		methodName := objectType.Method(i).Name
// 		methodSlice := strings.Split(strings.ToLower(methodName), "_")
// 		if len(methodSlice) != 3 {
// 			continue
// 		}
// 		var (
// 			rMethod string
// 			rPath   []string
// 		)
// 		rMethod = "GET"
// 		if v, ok := validMethos[rMethod]; ok {
// 			rMethod = v
// 		} else {
// 			rMethod = "GET"
// 		}
// 		if methodSlice[2] == "index" {
// 			rPath = append(rPath, "/"+methodSlice[1])
// 		}
// 		rPath = append(rPath, "/"+methodSlice[1]+"/"+methodSlice[2])

// 		f := object.Method(i)
// 		fn := func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 			log.Tracef("request url: %s", r.URL)
// 			context.Set(r, "params", p)
// 			in := []reflect.Value{reflect.ValueOf(w), reflect.ValueOf(r)}
// 			f.Call(in)
// 		}
// 		for _, v := range rPath {
// 			router.Handle(rMethod, v, fn)
// 		}
// 	}
// 	return router
// }
