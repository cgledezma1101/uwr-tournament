package models

import "time"

type Team struct {
	ID        int64     `gorm:"primaryKey"`
	Name      string
	ClubID    int64
	CreatedAt time.Time `gorm:"autoCreateTime:milli"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli"`

	// Relationships
	Club                 *Club                  `gorm:"foreignKey:ClubID"`
	Players              []Player               `gorm:"foreignKey:TeamID"`
	BlueTeamGames        []Game                 `gorm:"foreignKey:BlueTeamID"`
	WhiteTeamGames       []Game                 `gorm:"foreignKey:WhiteTeamID"`
	TournamentTeams      []TournamentTeam       `gorm:"foreignKey:TeamID"`
}

func (Team) TableName() string {
	return "teams"
}

type Player struct {
	ID        int64     `gorm:"primaryKey"`
	Number    int
	TeamID    int
	UserID    int64
	IsActive  *bool
	CreatedAt time.Time `gorm:"autoCreateTime:milli"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli"`

	// Relationships
	Team         *Team         `gorm:"foreignKey:TeamID"`
	User         *User         `gorm:"foreignKey:UserID"`
	PlayerGames  []PlayerGame  `gorm:"foreignKey:PlayerID"`
	Scores       []Score       `gorm:"foreignKey:PlayerID"`
}

func (Player) TableName() string {
	return "players"
}

type PlayerGame struct {
	ID        int64 `gorm:"primaryKey"`
	PlayerID  int
	GameID    int
	TeamColor string
	CreatedAt time.Time `gorm:"autoCreateTime:milli"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli"`

	// Relationships
	Player *Player `gorm:"foreignKey:PlayerID"`
	Game   *Game   `gorm:"foreignKey:GameID"`
}

func (PlayerGame) TableName() string {
	return "player_games"
}
