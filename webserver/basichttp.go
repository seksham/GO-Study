package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a test: %v", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/test", handler)
	err := http.ListenAndServe(":8000", nil)
	fmt.Println(err)
}
