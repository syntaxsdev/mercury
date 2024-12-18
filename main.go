package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/syntaxsdev/mercury/internal/api"
	"github.com/syntaxsdev/mercury/internal/repositories"
	"github.com/syntaxsdev/mercury/internal/services"
)

func main() {
	fmt.Println("Starting Mercury Core...")

	// Get environment variables
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	// Set defaults
	if dbHost == "" {
		dbHost = "localhost:27017"
		log.Printf("INFO: No `DB_HOST` set. Defaulting to %s\n", dbHost)
	}

	if dbName == "" {
		dbName = "mercury_db"
		log.Printf("INFO: No `DN_NAME` set. Defaulting to %s\n", dbName)
	}

	// Create MongoDB client
	mongo, cleanup := repositories.NewMongoClient(dbHost)
	defer cleanup()
	ms := services.MongoService{Client: mongo, DatabaseName: dbName}

	// Create the Factory and pass it into our Routes
	factory := services.Factory{MongoService: &ms}

	// Start Routes
	router := api.InitRoutes(&factory)

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Serve failed to start: %v", err)
	}
}
