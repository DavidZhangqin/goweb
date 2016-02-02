package main

import (
	"fmt"
	"log"
	"net/http"

	"lib/session"
)

func main() {
	session.LoadSession()

	http.HandleFunc("/login", Login)
	log.Println("Start listen 8089")
	http.ListenAndServe(":8089", nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	sess := session.Start(w, r)
	sess.Set("id", 123)
	user_id, _ := sess.Get("id")
	fmt.Fprintf(w, "user id %d", user_id.(int))
}
