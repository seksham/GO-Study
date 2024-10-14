package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

// User struct represents the structure of a user object
type User struct {
	User string `json:"user"`
}

// mutex for safe concurrent access to the userCache
var mutex sync.RWMutex

// userCache stores users with their IDs as keys
var userCache = make(map[int]User)

// mainHandler responds with a simple "Hello" message
func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

// createUserHandler creates a new user and adds it to the userCache
func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	// Decode the JSON request body into a User struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("Error here:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return // Added return statement to exit the function on error
	}
	
	// Lock the mutex before modifying the userCache
	mutex.Lock()
	defer mutex.Unlock()
	
	// Add the new user to the cache with an auto-incremented ID
	userCache[len(userCache)+1] = user
	fmt.Println(userCache)
	w.WriteHeader(http.StatusAccepted)
}

// getUserHandler retrieves a user by ID from the userCache
func getUserHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the ID from the URL path
	id := r.PathValue("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(intId)
	
	// Lock the mutex for reading from the userCache
	mutex.RLock()
	defer mutex.RUnlock()
	
	// Retrieve the user from the cache
	user, ok := userCache[intId]
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	// Encode and send the user as JSON response
	json.NewEncoder(w).Encode(user)
	// using json marshal
	// userJson, err := json.Marshal(user)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// w.Write(userJson)

}

// deleteUserHandler removes a user from the userCache by ID
func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the ID from the URL path
	id := r.PathValue("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	// Lock the mutex before modifying the userCache
	mutex.Lock()
	defer mutex.Unlock()
	
	// Check if the user exists before deleting
	if _, ok := userCache[intId]; !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	delete(userCache, intId)
	w.WriteHeader(http.StatusAccepted)
}

func main() {
	// Create a new ServeMux for routing
	mux := http.NewServeMux()
	
	// Register handlers for different routes
	mux.HandleFunc("/", mainHandler)
	mux.HandleFunc("POST /users", createUserHandler)
	mux.HandleFunc("GET /users/{id}", getUserHandler)
	mux.HandleFunc("DELETE /users/{id}", deleteUserHandler)
	
	fmt.Println("Listening to port: 8080")
	// Start the HTTP server
	http.ListenAndServe(":8080", mux)
}
