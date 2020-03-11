package main

import (
	"aviasales/handlers"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	hh := handlers.NewHandlers(mux)
	hh.AddHandler("/get", handlers.Get())
	hh.AddHandler("/load", handlers.Load())
	http.ListenAndServe(":8080", mux)
}
