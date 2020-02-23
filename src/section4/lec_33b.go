package main

import "net/http"

// usage: curl --data "key=goojy" localhost:30189
// reference: https://www.golangprograms.com/example-to-handle-get-and-post-request-in-golang.html
func main() {
	var mux *http.ServeMux = http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		var response string
		if request.Method == "POST" {
			switch request.FormValue("key") {
			case "kuchi": response = "puchi"
			case "goojy": response = "boojy"
			case "ping": response = "pong"
			case "monchi": response = "monchi"
			default:
				response = "unknown key"
			}
		} else {
			response = "please make POST request"
		}
		writer.Write([]byte(response))
	})
	http.ListenAndServe("localhost:30189", mux)
}
