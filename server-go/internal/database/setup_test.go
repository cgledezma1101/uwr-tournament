package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/uwr-tournament/server-go/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	testDBUser     = "test"
	testDBPassword = "test"
	testDBName     = "uwr-tournaments-test"
	databasePort   = "5433"
	maxRetries     = 30
	retryDelay     = 1 * time.Second
)

var Db *gorm.DB

func TestMain(m *testing.M) {
	fmt.Println("Starting unit test setup")

	// Locate the docker-compose.test.yml file
	composeFile := locateComposeFile()
	fmt.Printf("Using docker-compose file: %s\n", composeFile)

	// Start Docker Compose services
	fmt.Println("Starting Docker Compose services")
	upCmd := exec.Command("docker-compose", "-f", composeFile, "up", "-d")
	if output, err := upCmd.CombinedOutput(); err != nil {
		log.Fatalf("Failed to start docker-compose: %v\nOutput: %s", err, string(output))
	}

	// Wait for the database to be ready
	connString := fmt.Sprintf("postgresql://%s:%s@localhost:%s/%s?sslmode=disable",
		testDBUser, testDBPassword, databasePort, testDBName)

	// fmt.Println("Waiting for database to be ready")
	db, err := tryOpenDatabase(context.Background(), connString)
	if err != nil {
		log.Fatalf("Database failed to open database connection: %v", err)
	}

	Db = db

	// Run GORM migrations for all models
	fmt.Println("Running GORM migrations")
	if err := runTestMigrations(Db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	exitCode := m.Run()

	fmt.Println("Stopping Docker Compose services")
	downCmd := exec.Command("docker-compose", "-f", composeFile, "down")
	if output, err := downCmd.CombinedOutput(); err != nil {
		log.Printf("Warning: Failed to stop docker-compose: %v\nOutput: %s", err, string(output))
	}

	os.Exit(exitCode)
}

// locateComposeFile finds the docker-compose.test.yml file
func locateComposeFile() string {
	// Try PWD first
	pwdPath := fmt.Sprintf("%s/docker-compose.test.yml", os.Getenv("PWD"))
	fmt.Printf("Loading docker-compose from: %s\n", pwdPath)
	if fileExists(pwdPath) {
		return pwdPath
	}

	log.Fatal("docker-compose.test.yml not found.")
	return ""
}

// fileExists checks if a file exists
func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// tryOpenDatabase attempts to connect to the database with retries and exponential backoff.
// It uses the healthcheck approach similar to docker-compose.yml
func tryOpenDatabase(ctx context.Context, connString string) (*gorm.DB, error) {
	var lastErr error

	for attempt := 0; attempt < maxRetries; attempt++ {
		db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
		if err != nil {
			lastErr = err
			fmt.Printf("Attempt %d/%d: Failed to open connection - %v\n", attempt+1, maxRetries, err)
			time.Sleep(retryDelay)
			continue
		}
		fmt.Println("Database connection opened")
		return db, nil
	}

	return nil, fmt.Errorf("database failed to become ready after %d attempts: %w", maxRetries, lastErr)
}

// runTestMigrations runs all database migrations for testing using GORM
func runTestMigrations(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.Club{},
		&models.ClubAdmin{},
		&models.ClubJoinRequest{},
		&models.UserClub{},
		&models.Invitation{},
		&models.Team{},
		&models.Player{},
		&models.PlayerGame{},
		&models.Tournament{},
		&models.TournamentAdmin{},
		&models.TournamentInvitation{},
		&models.TournamentTeam{},
		&models.Stage{},
		&models.Game{},
		&models.GameEvent{},
		&models.Score{},
	)
}
