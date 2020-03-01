package main

import "fmt"

type Response struct {
	Code int         `json:"code"`
	Body interface{} `json:"body"`
}

type Todo struct {
	Name string `json:"name"`
	Todo string `json:"todo"`
}

func (todo Todo) ToString() string {
	return fmt.Sprintf("{\"Name\": \"%s\", \"Todo\": \"%s\"}", todo.Name, todo.Todo)
}
