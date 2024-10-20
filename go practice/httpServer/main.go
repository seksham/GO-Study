package main

import (
	"fmt"
	"net/http"
	"time"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second * 30)
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintf(w, "done")
}

func main() {
	fmt.Println("Starting server on port 8080")
	mux := http.NewServeMux()
	mux.HandleFunc("/", mainHandler)
	http.ListenAndServe(":8080", mux)
}
