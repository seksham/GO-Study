package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Main page")
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "User id is: %s", vars["id"])
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/user/{id}", userHandler)
	// Remove this line: http.Handle("/", r)
	fmt.Println("Server is running on http://localhost:8000")
	http.ListenAndServe(":8000", r)
}
