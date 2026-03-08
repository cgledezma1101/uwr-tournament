package models

import "time"

type Stage struct {
	ID           int       `gorm:"primaryKey"`
	TournamentID int
	Name         string
	CreatedAt    time.Time
	UpdatedAt    time.Time

	// Relationships
	Tournament *Tournament `gorm:"foreignKey:TournamentID"`
	Games      []Game      `gorm:"foreignKey:StageID"`
}

func (Stage) TableName() string {
	return "stages"
}

type Game struct {
	ID                    int       `gorm:"primaryKey"`
	BlueTeamID            int
	WhiteTeamID           int
	CreatedAt             time.Time
	UpdatedAt             time.Time
	WinningColor          string
	StageID               int
	Status                string
	StartsAt              *time.Time
	BlueTeamCalculation   string
	WhiteTeamCalculation  string

	// Relationships
	BlueTeam  *Team        `gorm:"foreignKey:BlueTeamID;references:ID"`
	WhiteTeam *Team        `gorm:"foreignKey:WhiteTeamID;references:ID"`
	Stage     *Stage       `gorm:"foreignKey:StageID"`
	GameEvents []GameEvent `gorm:"foreignKey:GameID"`
	PlayerGames []PlayerGame `gorm:"foreignKey:GameID"`
	Scores    []Score      `gorm:"foreignKey:GameID"`
}

func (Game) TableName() string {
	return "games"
}

type GameEvent struct {
	ID        int       `gorm:"primaryKey"`
	Text      string
	GameID    int
	CreatedAt time.Time
	UpdatedAt time.Time

	// Relationships
	Game *Game `gorm:"foreignKey:GameID"`
}

func (GameEvent) TableName() string {
	return "game_events"
}

type Score struct {
	ID        int       `gorm:"primaryKey"`
	PlayerID  int
	GameID    int
	CreatedAt time.Time
	UpdatedAt time.Time

	// Relationships
	Player *Player `gorm:"foreignKey:PlayerID"`
	Game   *Game   `gorm:"foreignKey:GameID"`
}

func (Score) TableName() string {
	return "scores"
}
