package session

import (
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"
)

var SessionCache map[string]*CacheStore

type CacheStore struct {
	SessId string
	MaxAge int
	Mutex  *sync.Mutex
	Vals   map[interface{}]interface{}
}

func CacheInit(name string, maxAge int) {
	SessionCache = make(map[string]*CacheStore)
	Name = name
	MaxAge = maxAge

	// gc session
	go func() {
		for {
			select {
			case <-time.After(time.Second):
				for k, v := range SessionCache {
					v.MaxAge--
					if v.MaxAge <= 0 {
						v.Mutex.Lock()
						delete(SessionCache, k)
						v.Mutex.Unlock()
					}
				}
			}
		}
	}()
	return
}

func NewCacheSession(w http.ResponseWriter, r *http.Request) (cs *CacheStore) {
	sessCookie, err := r.Cookie(Name)
	if err == nil {
		sessId := sessCookie.Value
		if cs, ok := SessionCache[sessId]; ok {
			return cs
		}
	}

	//get sid & set cookie
	sessId := GetSessId(r)
	cs = &CacheStore{
		SessId: sessId,
		MaxAge: MaxAge,
		Mutex:  &sync.Mutex{},
		Vals:   make(map[interface{}]interface{}),
	}
	SessionCache[sessId] = cs
	cookie := &http.Cookie{
		Name:     Name,
		Value:    url.QueryEscape(sessId),
		Path:     "/",
		HttpOnly: true,
		Domain:   "",
		Secure:   false,
		MaxAge:   MaxAge,
	}
	http.SetCookie(w, cookie)
	return
}

func (cs *CacheStore) Get(key interface{}) (interface{}, error) {
	if v, ok := cs.Vals[key]; ok {
		return v, nil
	} else {
		return nil, fmt.Errorf("session %s not exists")
	}
}

func (cs *CacheStore) Set(key interface{}, val interface{}) error {
	cs.Vals[key] = val
	return nil
}

func (cs *CacheStore) Del(key interface{}) error {
	if _, ok := cs.Vals[key]; ok {
		delete(cs.Vals, key)
	}
	return nil
}

func (cs *CacheStore) Ept() error {
	for k, _ := range cs.Vals {
		delete(cs.Vals, k)
	}
	return nil
}
