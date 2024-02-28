package main

import (
	"fmt"
	"net/http"
)

// a simple server responding with "Hello, World!"
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "v.0.0.1")
	})

	http.ListenAndServe(":8080", nil)
}
