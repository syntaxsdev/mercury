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
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")

	// Set defaults
	if dbHost == "" {
		dbHost = "localhost"
		log.Printf("INFO: No `DB_HOST` set. Defaulting to %s\n", dbHost)
	}

	if dbName == "" {
		dbName = "mercury_db"
		log.Printf("INFO: No `DN_NAME` set. Defaulting to %s\n", dbName)
	}

	if dbUser == "" {
		dbUser = "mercury"
		log.Printf("INFO: No `DB_USER` set. Defaulting to %s\n", dbUser)
	}

	if dbPass == "" {
		dbPass = "mercury_db"
		log.Printf("INFO: No `DB_PASSWORD` set. Defaulting...")
	}

	if dbPort == "" {
		dbPort = "27017"
		log.Printf("INFO: No `DB_USER` set. Defaulting to %s\n", dbUser)
	}

	// Create MongoDB client
	connString := fmt.Sprintf("mongodb://%s:%s@%s:%s", dbUser, dbPass, dbHost, dbPort)
	log.Println(connString)
	mongo, cleanup := repositories.NewMongoClient(connString)
	defer cleanup()
	ms := services.MongoService{Client: mongo, DatabaseName: dbName}

	// Create the Factory and pass it into our Routes
	factory := services.Factory{MongoService: &ms}

	// Start Routes
	router := api.InitRoutes(&factory)

	log.Printf("INFO: Starting Server... Listening on localhost:%s\n", dbPort)
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Serve failed to start: %v", err)
	}
}
