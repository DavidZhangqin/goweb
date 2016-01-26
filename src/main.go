package main

import (
	// "reflect"
	// "fmt"
	"log"
	"net/http"

	ctrl "controller"

	"github.com/julienschmidt/httprouter"
)

var actions *ctrl.CtrlStr

func main() {
	// controllers := [...]string{
	// 	"site",
	// 	"about",
	// }

	router := httprouter.New()
	// object := reflect.ValueOf(&ctrl.CtrlStr{})
	// for i := 0; i < object.NumMethod(); i++ {
	// 	log.Println(object.Method(i).Type())
	// 	router.GET("/site", httprouter.Handle(object.Method(i)))
	// 	break
	// }
	router.GET("/site", HandlerFunc(actions.Get_Site_Index))

	// for _, ctrlName := range controllers {
	// 	reflect.ValueOf(actions)
	// }

	// router.GET("/", Index)
	// router.GET("/hello/:name", Hello)
	// router.Handle("GET", "/test", Test)

	log.Println("ListenAndServe :8088")
	log.Fatal(http.ListenAndServe(":8088", router))
}

// object := reflect.ValueOf(MyMethods)
// 	f := object.MethodByName(method)
// 	if f.IsValid() {
// 		in := []reflect.Value{reflect.ValueOf(w), reflect.ValueOf(r)}
// 		f.Call(in)
// 	} else {
// 		Response(10001, "Method not exist", "", w)
// 	}

func HandlerFunc(h http.HandlerFunc) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		h.ServeHTTP(w, r)
		// h(w, r)
	}
}
