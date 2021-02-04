// http_server.go
package main

import (
	//"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Second * 10)

		w.Header().Add("Accept-Encoding", "gzip")
		w.Write([]byte("hello world!"))
	})

	http.ListenAndServe("127.0.0.1:8080", nil)
}
