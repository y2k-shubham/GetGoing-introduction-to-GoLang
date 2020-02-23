package main

import (
	"fmt"
	"net/http"
)

// usage: curl localhost:30189 --header "key:kuchi" / "ping"
// note that in code below, we still have to capitalize 'Key'
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		var responseString string = ""
		switch headerKey, headerKeyPresent := request.Header["Key"]; headerKeyPresent {
		case headerKeyPresent:
			if len(headerKey) > 0 {
				switch headerKey[0] {
				case "ping":
					responseString = "pong"
				case "kuchi":
					responseString = "puchi"
				default:
					responseString = fmt.Sprintf("unknown key: %s", headerKey[0])
				}
			} else {
				responseString = "Header key empty"
			}
		default:
			responseString = fmt.Sprintf("unknown request: header=%s", request.Header)
		}
		writer.Write([]byte(responseString))
	})
	http.ListenAndServe("localhost:30189", mux)
}
