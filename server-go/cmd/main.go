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

	"github.com/uwr-tournament/server-go/internal/auth"
	"github.com/uwr-tournament/server-go/internal/config"
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
	// Server config flags
	port := flag.String("port", "8080", "Port to run the server on")

	// Database connection flags
	dbHost := flag.String("db-host", "localhost", "Database host")
	dbPort := flag.String("db-port", "5432", "Database port")
	dbUser := flag.String("db-user", "postgres", "Database user")
	dbName := flag.String("db-name", "uwr_tournament", "Database name")
	dbSSLMode := flag.String("db-sslmode", "disable", "Database SSL mode")

	// OAuth flags
	googleClientID := flag.String("google-client-id", "", "Google OAuth Client ID")
	facebookClientID := flag.String("facebook-client-id", "", "Facebook OAuth Client ID")
	redirectBaseURL := flag.String("redirect-base-url", "http://localhost:8080", "Base URL for OAuth redirects")
	flag.Parse()

	envConfig, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config from environment: %v", err)
	}

	// Build connection string dynamically
	dbConnString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		*dbHost, *dbPort, *dbUser, envConfig.DatabasePassword, *dbName, *dbSSLMode,
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

	// Initialize IoC dependencies
	repos = database.NewRepositoryProvider(db.DB)

	// Initialize our IoC data structures
	userRepo := database.NewUserRepository(db.DB)
	passwordHasher := auth.NewPasswordHasher()
	authService := auth.NewAuthService(userRepo, passwordHasher)
	googleProvider := auth.NewGoogleOAuthProvider(
		*googleClientID,
		envConfig.GoogleClientSecret,
		*redirectBaseURL+"/oauth/google/callback",
		"https://accounts.google.com/",
		"https://oauth2.googleapis.com",
		"https://www.googleapis.com",
	)
	facebookProvider := auth.NewFacebookOAuthProvider(
		*facebookClientID,
		envConfig.FacebookClientSecret,
		*redirectBaseURL+"/oauth/facebook/callback",
		"https://www.facebook.com",
		"https://graph.instagram.com",
		"https://graph.instagram.com",
	)
	authHandler := auth.NewHTTPHandler(
		authService,
		passwordHasher,
		googleProvider,
		facebookProvider,
	)

	// Setup routes
	mux := createMux(authHandler)

	// Create server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", *port),
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

func createMux(authHandler *auth.HTTPHandler) *http.ServeMux {
	mux := http.NewServeMux()

	// Healthcheck routes
	mux.HandleFunc("/health", healthCheckHandler)

	// Authentication routes
	mux.HandleFunc("POST /api/v2/auth/login", authHandler.LoginHandler)
	mux.HandleFunc("POST /api/v2/auth/register", authHandler.RegisterHandler)
	mux.HandleFunc("GET /oauth/google", authHandler.GoogleOAuthHandler)
	mux.HandleFunc("GET /oauth/google/callback", authHandler.GoogleOAuthCallbackHandler)
	mux.HandleFunc("GET /oauth/facebook", authHandler.FacebookOAuthHandler)
	mux.HandleFunc("GET /oauth/facebook/callback", authHandler.FacebookOAuthCallbackHandler)

	return mux
}
