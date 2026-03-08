package models

import (
	"time"
)

type Tournament struct {
	ID        int       `gorm:"primaryKey"`
	Name      string
	StartDate *time.Time
	EndDate   *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time

	// Relationships
	Stages                   []Stage                   `gorm:"foreignKey:TournamentID"`
	TournamentAdmins         []TournamentAdmin         `gorm:"foreignKey:TournamentID"`
	TournamentInvitations    []TournamentInvitation    `gorm:"foreignKey:TournamentID"`
	TournamentTeams          []TournamentTeam          `gorm:"foreignKey:TournamentID"`
}

func (Tournament) TableName() string {
	return "tournaments"
}

type TournamentAdmin struct {
	ID           int       `gorm:"primaryKey;autoIncrement:false"`
	UserID       int
	TournamentID int
	CreatedAt    time.Time
	UpdatedAt    time.Time

	// Relationships
	User       *User       `gorm:"foreignKey:UserID"`
	Tournament *Tournament `gorm:"foreignKey:TournamentID"`
}

func (TournamentAdmin) TableName() string {
	return "tournament_admins"
}

type TournamentInvitation struct {
	ID           int       `gorm:"primaryKey;autoIncrement:false"`
	ClubID       int
	TournamentID int
	CreatedAt    time.Time
	UpdatedAt    time.Time

	// Relationships
	Club       *Club       `gorm:"foreignKey:ClubID"`
	Tournament *Tournament `gorm:"foreignKey:TournamentID"`
}

func (TournamentInvitation) TableName() string {
	return "tournament_invitations"
}

type TournamentTeam struct {
	ID           int       `gorm:"primaryKey;autoIncrement:false"`
	TournamentID int
	TeamID       int
	Password     string
	CreatedAt    time.Time
	UpdatedAt    time.Time

	// Relationships
	Tournament *Tournament `gorm:"foreignKey:TournamentID"`
	Team       *Team       `gorm:"foreignKey:TeamID"`
}

func (TournamentTeam) TableName() string {
	return "tournament_teams"
}
