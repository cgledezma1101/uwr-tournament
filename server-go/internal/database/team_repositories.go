package database

import (
	"github.com/uwr-tournament/server-go/internal/models"
	"gorm.io/gorm"
)

type GormTeamRepository struct {
	db *gorm.DB
}

func NewTeamRepository(db *gorm.DB) *GormTeamRepository {
	return &GormTeamRepository{db: db}
}

func (r *GormTeamRepository) Create(team *models.Team) error {
	return r.db.Create(team).Error
}

func (r *GormTeamRepository) GetByID(id int64) (*models.Team, error) {
	var team models.Team
	if err := r.db.First(&team, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &team, nil
}

func (r *GormTeamRepository) Update(team *models.Team) error {
	return r.db.Save(team).Error
}

func (r *GormTeamRepository) Delete(id int64) error {
	return r.db.Delete(&models.Team{}, id).Error
}

func (r *GormTeamRepository) List(offset, limit int) ([]models.Team, error) {
	var teams []models.Team
	if err := r.db.Offset(offset).Limit(limit).Find(&teams).Error; err != nil {
		return nil, err
	}
	return teams, nil
}

func (r *GormTeamRepository) GetByClubID(clubID int64) ([]models.Team, error) {
	var teams []models.Team
	if err := r.db.Where("club_id = ?", clubID).Find(&teams).Error; err != nil {
		return nil, err
	}
	return teams, nil
}

type GormPlayerRepository struct {
	db *gorm.DB
}

func NewPlayerRepository(db *gorm.DB) *GormPlayerRepository {
	return &GormPlayerRepository{db: db}
}

func (r *GormPlayerRepository) Create(player *models.Player) error {
	return r.db.Create(player).Error
}

func (r *GormPlayerRepository) GetByID(id int64) (*models.Player, error) {
	var player models.Player
	if err := r.db.First(&player, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &player, nil
}

func (r *GormPlayerRepository) Update(player *models.Player) error {
	return r.db.Save(player).Error
}

func (r *GormPlayerRepository) Delete(id int64) error {
	return r.db.Delete(&models.Player{}, id).Error
}

func (r *GormPlayerRepository) List(offset, limit int) ([]models.Player, error) {
	var players []models.Player
	if err := r.db.Offset(offset).Limit(limit).Find(&players).Error; err != nil {
		return nil, err
	}
	return players, nil
}

func (r *GormPlayerRepository) GetByTeamID(teamID int) ([]models.Player, error) {
	var players []models.Player
	if err := r.db.Where("team_id = ?", teamID).Find(&players).Error; err != nil {
		return nil, err
	}
	return players, nil
}

func (r *GormPlayerRepository) GetByUserID(userID int64) ([]models.Player, error) {
	var players []models.Player
	if err := r.db.Where("user_id = ?", userID).Find(&players).Error; err != nil {
		return nil, err
	}
	return players, nil
}

type GormPlayerGameRepository struct {
	db *gorm.DB
}

func NewPlayerGameRepository(db *gorm.DB) *GormPlayerGameRepository {
	return &GormPlayerGameRepository{db: db}
}

func (r *GormPlayerGameRepository) Create(playerGame *models.PlayerGame) error {
	return r.db.Create(playerGame).Error
}

func (r *GormPlayerGameRepository) GetByGameID(gameID int) ([]models.PlayerGame, error) {
	var playerGames []models.PlayerGame
	if err := r.db.Where("game_id = ?", gameID).Find(&playerGames).Error; err != nil {
		return nil, err
	}
	return playerGames, nil
}

func (r *GormPlayerGameRepository) GetByPlayerID(playerID int64) ([]models.PlayerGame, error) {
	var playerGames []models.PlayerGame
	if err := r.db.Where("player_id = ?", playerID).Find(&playerGames).Error; err != nil {
		return nil, err
	}
	return playerGames, nil
}

func (r *GormPlayerGameRepository) Delete(id int64) error {
	return r.db.Delete(&models.PlayerGame{}, id).Error
}
