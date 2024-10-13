package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

type User struct {
	User string `json:"user"`
}

var mutex sync.RWMutex

var userCache = make(map[int]User)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("Error here:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	userCache[len(userCache)+1] = user
	fmt.Println(userCache)
	w.WriteHeader(http.StatusAccepted)
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	fmt.Println(intId)
	user, ok := userCache[intId]
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	if _, ok := userCache[intId]; !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	delete(userCache, intId)
	w.WriteHeader(http.StatusAccepted)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", mainHandler)
	mux.HandleFunc("POST /users", createUserHandler)
	mux.HandleFunc("GET /users/{id}", getUserHandler)
	fmt.Println("Listening to port: 8080")
	http.ListenAndServe(":8080", mux)
}
