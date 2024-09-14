package routes

import (
	"github.com/gorilla/mux"
	"go_template/handlers"
)

// RegisterUserRoutes sets up routes related to user functionality
func RegisterUserRoutes(r *mux.Router) {
	r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")
}
