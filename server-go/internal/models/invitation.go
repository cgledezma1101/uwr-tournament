package models

import "time"

type Invitation struct {
	ID        int       `gorm:"primaryKey;autoIncrement:false"`
	ClubID    int
	UserID    int
	IsAdmin   *bool
	CreatedAt time.Time
	UpdatedAt time.Time

	// Relationships
	Club *Club `gorm:"foreignKey:ClubID"`
	User *User `gorm:"foreignKey:UserID"`
}

func (Invitation) TableName() string {
	return "invitations"
}
