package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodGet {
			data := Response{
				Code: http.StatusOK,
				Body: "pong",
			}
			json.NewEncoder(writer).Encode(data)
		}
	})

	http.ListenAndServe("localhost:30189", mux)
}
