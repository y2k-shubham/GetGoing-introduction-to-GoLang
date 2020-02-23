package main

import (
	"encoding/json"
	"net/http"
)

func Register() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", Ping)
	return mux
}

func Ping(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		data := Response{
			Code: http.StatusOK,
			Body: "pong",
		}
		json.NewEncoder(writer).Encode(data)
	}
}
