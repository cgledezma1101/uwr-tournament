package database

import (
	"fmt"
	"log"

	"github.com/uwr-tournament/server-go/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

// InitDatabase initializes the database connection and runs migrations
func InitDatabase(connectionString string) (*Database, error) {
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := runMigrations(db); err != nil {
		return nil, fmt.Errorf("failed to run migrations: %w", err)
	}

	log.Println("Database initialized successfully")
	return &Database{DB: db}, nil
}

// runMigrations runs all database migrations using GORM
func runMigrations(db *gorm.DB) error {
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

// Close closes the database connection
func (d *Database) Close() error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
