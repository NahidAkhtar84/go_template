package main

import (
	"log"
	"net/http"

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
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
