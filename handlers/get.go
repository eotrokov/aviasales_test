package handlers

import (
	"aviasales/services"
	http2 "net/http"
)

type get struct {
}

func Get() *get {
	return &get{}
}

func (m *get) getMessage(req *http2.Request) interface{} {
	if req.Method != "GET" {
		return "Sorry, only GET methods are supported."
	}
	word := req.URL.Query().Get("word")
	if word == "" {
		return "Url Param 'word' is missing"
	}
	return services.GetAnnogram(word)
}
