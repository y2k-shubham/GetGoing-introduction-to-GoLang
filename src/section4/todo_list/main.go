package main

import (
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// database connection
	db := Connect()
	defer db.Close()

	// request multiplexer
	mux := Register()

	// start server
	log.Println("Serving")
	http.ListenAndServe("localhost:30189", mux)
}
