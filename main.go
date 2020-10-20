package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HomeHandler)
	http.ListenAndServe(":8080", nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "/")
}
