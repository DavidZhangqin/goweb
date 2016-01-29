package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"runtime"
	"strings"

	ctrl "controller"
)

var Actions *ctrl.CtrlStr

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	http.HandleFunc("/", mainHandle)

	log.Println("ListenAndServe :8089")
	log.Fatal(http.ListenAndServe(":8089", nil))
}

func mainHandle(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path[0] == []byte("/")[0] {
		path = path[1:]
	}
	urls := strings.Split(path, "/")
	methodName := strings.Title(strings.ToLower(r.Method))
	for _, v := range urls {
		methodName += "_" + strings.Title(v)
	}

	object := reflect.ValueOf(Actions)
	f := object.MethodByName(methodName)
	if f.IsValid() {
		in := []reflect.Value{reflect.ValueOf(w), reflect.ValueOf(r)}
		f.Call(in)
	} else {
		fmt.Fprint(w, "404 not found!\n")
	}

}
