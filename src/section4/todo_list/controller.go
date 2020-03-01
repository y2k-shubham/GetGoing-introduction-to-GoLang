package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func Register() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", Ping)
	mux.HandleFunc("/create", Create)
	mux.HandleFunc("/read", Read)
	mux.HandleFunc("/read/", ReadByName)
	mux.HandleFunc("/delete/", Delete)
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
		var todo = Todo{
			Name: request.FormValue("name"),
			Todo: request.FormValue("todo"),
		}
		if err := CreateTodo(todo); err != nil {
			writer.Write([]byte("Some error occured: " + err.Error()))
		} else {
			writer.WriteHeader(http.StatusCreated)
			json.NewEncoder(writer).Encode("Todo created " + todo.ToString())
		}
	}
}

func Read(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		todos, err := ReadTodos()
		ReadHelper(writer, todos, err)
	}
}

func ReadByName(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		nameFilter := request.URL.Query().Get("name")
		todos, err := ReadTodosByName(nameFilter)
		ReadHelper(writer, todos, err)
	}
}

func ReadHelper(writer http.ResponseWriter, todos []Todo, err error) {
	if err != nil {
		writer.Write([]byte("Some error occured: " + err.Error()))
	} else {
		writer.WriteHeader(http.StatusOK)
		if todos == nil {
			writer.Write([]byte("Result set empty!"))
		} else {
			json.NewEncoder(writer).Encode(todos)
		}
	}
}

func Delete(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodDelete {
		nameFilter := strings.Replace(request.URL.Path, "/delete/", "", -1)
		if err := DeleteTodo(nameFilter); err != nil {
			writer.Write([]byte("Some error occured: " + err.Error()))
		} else {
			writer.WriteHeader(http.StatusOK)
			json.NewEncoder(writer).Encode("Deleted records for name = " + nameFilter)
		}
	}
}
