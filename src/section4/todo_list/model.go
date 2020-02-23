package main

import (
	"database/sql"
	"log"
)

var con *sql.DB

func Connect() *sql.DB {
	if con == nil {
		db, err := sql.Open("mysql", "root:root@(tcp:127.0.0.1:3307)/todo_list")
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Connected to database")
		con = db
	} else {
		log.Println("Reusing database connection")
	}
	return con
}
