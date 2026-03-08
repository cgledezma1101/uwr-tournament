package database

import (
	"github.com/uwr-tournament/server-go/internal/models"
	"gorm.io/gorm"
)

type GormGameRepository struct {
	db *gorm.DB
}

func NewGameRepository(db *gorm.DB) *GormGameRepository {
	return &GormGameRepository{db: db}
}

func (r *GormGameRepository) Create(game *models.Game) error {
	return r.db.Create(game).Error
}

func (r *GormGameRepository) GetByID(id int) (*models.Game, error) {
	var game models.Game
	if err := r.db.First(&game, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &game, nil
}

func (r *GormGameRepository) Update(game *models.Game) error {
	return r.db.Save(game).Error
}

func (r *GormGameRepository) Delete(id int) error {
	return r.db.Delete(&models.Game{}, id).Error
}

func (r *GormGameRepository) List(offset, limit int) ([]models.Game, error) {
	var games []models.Game
	if err := r.db.Offset(offset).Limit(limit).Find(&games).Error; err != nil {
		return nil, err
	}
	return games, nil
}

func (r *GormGameRepository) GetByStageID(stageID int) ([]models.Game, error) {
	var games []models.Game
	if err := r.db.Where("stage_id = ?", stageID).Find(&games).Error; err != nil {
		return nil, err
	}
	return games, nil
}

type GormGameEventRepository struct {
	db *gorm.DB
}

func NewGameEventRepository(db *gorm.DB) *GormGameEventRepository {
	return &GormGameEventRepository{db: db}
}

func (r *GormGameEventRepository) Create(event *models.GameEvent) error {
	return r.db.Create(event).Error
}

func (r *GormGameEventRepository) GetByGameID(gameID int) ([]models.GameEvent, error) {
	var events []models.GameEvent
	if err := r.db.Where("game_id = ?", gameID).Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}

func (r *GormGameEventRepository) Delete(id int) error {
	return r.db.Delete(&models.GameEvent{}, id).Error
}

type GormScoreRepository struct {
	db *gorm.DB
}

func NewScoreRepository(db *gorm.DB) *GormScoreRepository {
	return &GormScoreRepository{db: db}
}

func (r *GormScoreRepository) Create(score *models.Score) error {
	return r.db.Create(score).Error
}

func (r *GormScoreRepository) GetByGameID(gameID int) ([]models.Score, error) {
	var scores []models.Score
	if err := r.db.Where("game_id = ?", gameID).Find(&scores).Error; err != nil {
		return nil, err
	}
	return scores, nil
}

func (r *GormScoreRepository) GetByPlayerID(playerID int) ([]models.Score, error) {
	var scores []models.Score
	if err := r.db.Where("player_id = ?", playerID).Find(&scores).Error; err != nil {
		return nil, err
	}
	return scores, nil
}

func (r *GormScoreRepository) Delete(id int) error {
	return r.db.Delete(&models.Score{}, id).Error
}

type GormStageRepository struct {
	db *gorm.DB
}

func NewStageRepository(db *gorm.DB) *GormStageRepository {
	return &GormStageRepository{db: db}
}

func (r *GormStageRepository) Create(stage *models.Stage) error {
	return r.db.Create(stage).Error
}

func (r *GormStageRepository) GetByID(id int) (*models.Stage, error) {
	var stage models.Stage
	if err := r.db.First(&stage, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &stage, nil
}

func (r *GormStageRepository) Update(stage *models.Stage) error {
	return r.db.Save(stage).Error
}

func (r *GormStageRepository) Delete(id int) error {
	return r.db.Delete(&models.Stage{}, id).Error
}

func (r *GormStageRepository) GetByTournamentID(tournamentID int) ([]models.Stage, error) {
	var stages []models.Stage
	if err := r.db.Where("tournament_id = ?", tournamentID).Find(&stages).Error; err != nil {
		return nil, err
	}
	return stages, nil
}
