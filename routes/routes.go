package routes

import (
	"github.com/gorilla/mux"
)

// InitRoutes initializes all the routes for the application
func InitRoutes() *mux.Router {
	r := mux.NewRouter()

	// Register routes for different modules
	RegisterUserRoutes(r)

	return r
}
