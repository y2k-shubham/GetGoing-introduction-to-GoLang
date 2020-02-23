package main

import (
	"net/http"
)

func main() {
	mux := Register()

	http.ListenAndServe("localhost:30189", mux)
}
