package main

import (
	"database/sql"
	"fmt"
	"log"
)

var con *sql.DB
var db string = "todo_list"
var table string = "todo"

func Connect() *sql.DB {
	if con == nil {
		db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3307)/todo_list")
		if err != nil {
			log.Println("Error while making database connection: " + err.Error())
			log.Fatal(err)
		} else {
			log.Println("Connected to database")
			con = db
		}
	} else {
		log.Println("Reusing database connection")
	}
	return con
}

func CreateTodo(todo Todo) error {
	queryTmplt := "INSERT INTO %s.%s VALUES ('%s', '%s')"
	query := fmt.Sprintf(queryTmplt, db, table, todo.Name, todo.Todo)

	rows, err := Connect().Query(query)
	if rows != nil {
		defer rows.Close()
	}

	if err != nil {
		return err
	} else {
		return nil
	}
}
