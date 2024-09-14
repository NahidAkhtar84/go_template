package main

import (
	"log"
	"net/http"
	"os"

	"go_template/db"
	"go_template/routes"
)

func main() {
	// Initialize the database connection
	db.InitDB()
	defer db.DB.Close()

	// Initialize all routes
	r := routes.InitRoutes()

	// Start server
	port := os.Getenv("APP_DOCKER_PORT")
	if port == "" {
		port = "8080" // Default to 8080 if PORT is not set
	}

	log.Printf("🚀 Server is running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
