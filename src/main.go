package main

import (
	"reflect"
	"strings"
	// "fmt"
	"log"
	"net/http"

	ctrl "controller"

	"github.com/julienschmidt/httprouter"
)

func main() {

	router := httprouter.New()
	object := reflect.ValueOf(&ctrl.CtrlStr{})
	objectType := object.Type()
	for i := 0; i < object.NumMethod(); i++ {
		methodName := objectType.Method(i).Name
		log.Println(methodName)
		methodSlice := strings.Split(strings.ToLower(methodName), "_")
		if len(methodSlice) != 3 {
			continue
		}
		var (
			rMethod string
			rPath   []string
		)
		validMethos := map[string]string{
			"get":    "GET",
			"post":   "POST",
			"put":    "PUT",
			"delete": "DELETE",
			"patch":  "PATCH",
		}
		rMethod = "GET"
		if v, ok := validMethos[rMethod]; ok {
			rMethod = v
		} else {
			rMethod = "GET"
		}
		if methodSlice[2] == "index" {
			rPath = append(rPath, "/"+methodSlice[1])
		}
		rPath = append(rPath, "/"+methodSlice[1]+"/"+methodSlice[2])

		f := object.Method(i)
		fn := func(w http.ResponseWriter, r *http.Request) {
			in := []reflect.Value{reflect.ValueOf(w), reflect.ValueOf(r)}
			f.Call(in)
		}
		log.Println(rMethod, rPath)
		for _, v := range rPath {
			router.HandlerFunc(rMethod, v, fn)
		}
	}

	log.Println("ListenAndServe :8088")
	log.Fatal(http.ListenAndServe(":8088", router))

}
