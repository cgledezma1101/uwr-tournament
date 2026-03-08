package models

import "time"

type Club struct {
	ID        int       `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time

	// Relationships
	UserClubs            []UserClub           `gorm:"foreignKey:ClubID"`
	ClubAdmins           []ClubAdmin          `gorm:"foreignKey:ClubID"`
	ClubJoinRequests     []ClubJoinRequest    `gorm:"foreignKey:ClubID"`
	Invitations          []Invitation         `gorm:"foreignKey:ClubID"`
	Teams                []Team               `gorm:"foreignKey:ClubID"`
	TournamentInvitations []TournamentInvitation `gorm:"foreignKey:ClubID"`
}

func (Club) TableName() string {
	return "clubs"
}

type ClubAdmin struct {
	ID        int       `gorm:"primaryKey;autoIncrement:false"`
	UserID    int
	ClubID    int
	CreatedAt time.Time
	UpdatedAt time.Time

	// Relationships
	User *User `gorm:"foreignKey:UserID"`
	Club *Club `gorm:"foreignKey:ClubID"`
}

func (ClubAdmin) TableName() string {
	return "club_admins"
}

type ClubJoinRequest struct {
	ID        int       `gorm:"primaryKey;autoIncrement:false"`
	UserID    int
	ClubID    int
	CreatedAt time.Time
	UpdatedAt time.Time

	// Relationships
	User *User `gorm:"foreignKey:UserID"`
	Club *Club `gorm:"foreignKey:ClubID"`
}

func (ClubJoinRequest) TableName() string {
	return "club_join_requests"
}

type UserClub struct {
	ID        int       `gorm:"primaryKey;autoIncrement:false"`
	UserID    int
	ClubID    int
	CreatedAt time.Time
	UpdatedAt time.Time

	// Relationships
	User *User `gorm:"foreignKey:UserID"`
	Club *Club `gorm:"foreignKey:ClubID"`
}

func (UserClub) TableName() string {
	return "user_clubs"
}
