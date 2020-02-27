package main

import (
	"encoding/json"
	"net/http"
)

func Register() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", Ping)
	mux.HandleFunc("/create", Create)
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

func Create(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		var todo Todo = Todo{
			Name: request.FormValue("Name"),
			Todo: request.FormValue("Todo"),
		}
		if err := CreateTodo(todo); err != nil {
			writer.Write([]byte("Some error occured: " + err.Error()))
		} else {
			writer.Write([]byte("Todo created"))
		}
	}
}
