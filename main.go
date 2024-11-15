package main

import (
	"net/http"
	"test/internal/handler"
)

func main() {
	// mux - стандартный гошный роутер
	mux := http.NewServeMux()
	mux.HandleFunc("/message", handler.HandleMessage)

	http.ListenAndServe(":80", mux)
}
