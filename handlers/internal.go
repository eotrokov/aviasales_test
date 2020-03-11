package handlers

import (
	"encoding/json"
	"net/http"
)

type messageHandler interface {
	getMessage(*http.Request) interface{}
}

type Handlers struct {
	*http.ServeMux
	Links []string
}

func NewHandlers(mux *http.ServeMux) *Handlers {
	return &Handlers{
		ServeMux: mux,
		Links:    make([]string, 0),
	}
}

func (hh *Handlers) AddHandler(path string, h messageHandler) *Handlers {
	hh.Links = append(hh.Links, path)
	hh.ServeMux.HandleFunc(path, getHandler(h))
	return hh
}

func getHandler(ms messageHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		res := ms.getMessage(r)
		enc := json.NewEncoder(w)
		enc.SetIndent("", "   ")
		enc.Encode(res)
	}
}
