package main

import "net/http"

func main() {
	// to verify this is running, try 'telnet localhost 30189'
	http.ListenAndServe("localhost:30189", nil)
}
