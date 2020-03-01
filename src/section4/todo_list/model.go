package main

import (
	"database/sql"
	"fmt"
	"log"
)

var con *sql.DB
var db string = "todo_list"
var table string = "todo"
var name string = "name"

func GetConnection() *sql.DB {
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

	rows, err := GetConnection().Query(query)
	if rows != nil {
		defer rows.Close()
	}

	if err != nil {
		return err
	} else {
		return nil
	}
}

func ReadTodos() ([]Todo, error) {
	queryTmplt := "SELECT * FROM %s.%s"
	query := fmt.Sprintf(queryTmplt, db, table)

	return GetSelectQueryResults(query)
}

func ReadTodosByName(nameFilter string) ([]Todo, error) {
	queryTmplt := "SELECT * FROM %s.%s WHERE %s = '%s'"
	query := fmt.Sprintf(queryTmplt, db, table, name, nameFilter)

	return GetSelectQueryResults(query)
}

func GetSelectQueryResults(query string) ([]Todo, error) {
	rows, err := GetConnection().Query(query)
	var todos []Todo

	if rows != nil {
		defer rows.Close()
		if err == nil {
			for rows.Next() {
				var todo Todo = Todo{}
				err := rows.Scan(&todo.Name, &todo.Todo)
				if err != nil {
					log.Fatal(err)
				} else {
					todos = append(todos, todo)
				}
			}
		} else {
			log.Fatal(err)
		}
	} else {
		fmt.Println("Got no rows")
	}

	return todos, nil
}
