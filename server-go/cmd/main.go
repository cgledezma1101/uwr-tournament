package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/uwr-tournament/server-go/internal/database"
)

type HealthResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

var repos *database.RepositoryProvider //nolint:unused

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := HealthResponse{
		Status: "healthy",
		Code:   http.StatusOK,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding health response: %v", err)
	}
}

func main() {
	// Define CLI flags for database connection
	dbHost := flag.String("db-host", "localhost", "Database host")
	dbPort := flag.String("db-port", "5432", "Database port")
	dbUser := flag.String("db-user", "postgres", "Database user")
	dbName := flag.String("db-name", "uwr_tournament", "Database name")
	dbSSLMode := flag.String("db-sslmode", "disable", "Database SSL mode")
	flag.Parse()

	// Get password from environment variable
	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		log.Fatal("Error: DB_PASSWORD environment variable not set")
	}

	// Build connection string dynamically
	dbConnString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		*dbHost, *dbPort, *dbUser, dbPassword, *dbName, *dbSSLMode,
	)

	// Initialize database
	db, err := database.InitDatabase(dbConnString)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()

	// Initialize repositories
	repos = database.NewRepositoryProvider(db.DB)
	log.Println("Repository provider initialized successfully")

	// Setup routes
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthCheckHandler)

	// Create server
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Starting server on %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Handle graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan
	log.Println("Shutting down server...")

	if err := server.Close(); err != nil {
		log.Fatalf("Error closing server: %v", err)
	}

	log.Println("Server stopped")
}
