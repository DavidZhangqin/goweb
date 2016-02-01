package session

import (
	"net/http"
)

func GetSessId(r *http.Request) string {
	return "abcdefg"
}

var Name string
var MaxAge int

type Store interface {
	Get(key interface{}) (interface{}, error)
	Set(key interface{}, val interface{}) error
	Del(key interface{}) error
	Ept() error
}
