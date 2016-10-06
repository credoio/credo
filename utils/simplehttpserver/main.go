package main

import (
	"fmt"
	"net/http"
)

// Works like:
//      python -m SimpleHTTPServer
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	fmt.Println("serving HTTP at port 8000")
	http.ListenAndServe(":8000", nil)
}
