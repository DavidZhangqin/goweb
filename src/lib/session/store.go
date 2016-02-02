package session

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	"net/http"
)

func GetSessId(_ *http.Request) string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

var Name string
var MaxAge int

type Store interface {
	Get(key interface{}) (interface{}, error)
	Set(key interface{}, val interface{}) error
	Del(key interface{}) error
	Ept() error
}
