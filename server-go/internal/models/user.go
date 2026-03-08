package models

import "time"

type User struct {
	ID                   int       `gorm:"primaryKey"`
	Email                string    `gorm:"uniqueIndex;not null;default:''"`
	EncryptedPassword    string    `gorm:"not null;default:''"`
	ResetPasswordToken   string
	ResetPasswordSentAt  *time.Time
	RememberCreatedAt    *time.Time
	SignInCount          int       `gorm:"not null;default:0"`
	CurrentSignInAt      *time.Time
	LastSignInAt         *time.Time
	CurrentSignInIP      string
	LastSignInIP         string
	CreatedAt            time.Time
	UpdatedAt            time.Time
	Name                 string
	Provider             string

	// Relationships
	UserClubs             []UserClub       `gorm:"foreignKey:UserID"`
	ClubAdmins            []ClubAdmin      `gorm:"foreignKey:UserID"`
	ClubJoinRequests      []ClubJoinRequest `gorm:"foreignKey:UserID"`
	Invitations           []Invitation     `gorm:"foreignKey:UserID"`
	Players               []Player         `gorm:"foreignKey:UserID"`
	TournamentAdmins      []TournamentAdmin `gorm:"foreignKey:UserID"`
}

func (User) TableName() string {
	return "users"
}
