package handlers

import (
	"aviasales/store"
	"encoding/json"
	http2 "net/http"
)

type load struct {
}

func Load() *load {
	return &load{}
}

func (m *load) getMessage(req *http2.Request) interface{} {
	if req.Method != "POST" {
		return "Sorry, only POST methods are supported."
	}
	decoder := json.NewDecoder(req.Body)
	var t []string
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	store.AddToStore(t)
	return t
}
