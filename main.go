package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go_template/db"
	"go_template/handlers"
)

func main() {
	// Initialize the database connection
	db.InitDB()
	defer db.DB.Close()

	// Set up router
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

	// Start server
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
