package main

import (
	"net/http"
)

// usage: curl localhost:30189/ping
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("pong"))
	})
	mux.HandleFunc("/kuchi", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("puchi"))
	})
	mux.HandleFunc("/goojy", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("boojy"))
	})
	mux.HandleFunc("/monchi", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("monchi"))
	})
	http.ListenAndServe("localhost:30189", mux)
}
