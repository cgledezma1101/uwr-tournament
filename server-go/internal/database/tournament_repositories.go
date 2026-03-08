package database

import (
	"github.com/uwr-tournament/server-go/internal/models"
	"gorm.io/gorm"
)

type GormTournamentRepository struct {
	db *gorm.DB
}

func NewTournamentRepository(db *gorm.DB) *GormTournamentRepository {
	return &GormTournamentRepository{db: db}
}

func (r *GormTournamentRepository) Create(tournament *models.Tournament) error {
	return r.db.Create(tournament).Error
}

func (r *GormTournamentRepository) GetByID(id int) (*models.Tournament, error) {
	var tournament models.Tournament
	if err := r.db.First(&tournament, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &tournament, nil
}

func (r *GormTournamentRepository) Update(tournament *models.Tournament) error {
	return r.db.Save(tournament).Error
}

func (r *GormTournamentRepository) Delete(id int) error {
	return r.db.Delete(&models.Tournament{}, id).Error
}

func (r *GormTournamentRepository) List(offset, limit int) ([]models.Tournament, error) {
	var tournaments []models.Tournament
	if err := r.db.Offset(offset).Limit(limit).Find(&tournaments).Error; err != nil {
		return nil, err
	}
	return tournaments, nil
}

type GormTournamentAdminRepository struct {
	db *gorm.DB
}

func NewTournamentAdminRepository(db *gorm.DB) *GormTournamentAdminRepository {
	return &GormTournamentAdminRepository{db: db}
}

func (r *GormTournamentAdminRepository) Create(admin *models.TournamentAdmin) error {
	return r.db.Create(admin).Error
}

func (r *GormTournamentAdminRepository) GetByTournamentID(tournamentID int) ([]models.TournamentAdmin, error) {
	var admins []models.TournamentAdmin
	if err := r.db.Where("tournament_id = ?", tournamentID).Find(&admins).Error; err != nil {
		return nil, err
	}
	return admins, nil
}

func (r *GormTournamentAdminRepository) DeleteByUserAndTournament(userID, tournamentID int) error {
	return r.db.Where("user_id = ? AND tournament_id = ?", userID, tournamentID).Delete(&models.TournamentAdmin{}).Error
}

type GormTournamentInvitationRepository struct {
	db *gorm.DB
}

func NewTournamentInvitationRepository(db *gorm.DB) *GormTournamentInvitationRepository {
	return &GormTournamentInvitationRepository{db: db}
}

func (r *GormTournamentInvitationRepository) Create(invitation *models.TournamentInvitation) error {
	return r.db.Create(invitation).Error
}

func (r *GormTournamentInvitationRepository) GetByTournamentID(tournamentID int) ([]models.TournamentInvitation, error) {
	var invitations []models.TournamentInvitation
	if err := r.db.Where("tournament_id = ?", tournamentID).Find(&invitations).Error; err != nil {
		return nil, err
	}
	return invitations, nil
}

func (r *GormTournamentInvitationRepository) GetByClubID(clubID int) ([]models.TournamentInvitation, error) {
	var invitations []models.TournamentInvitation
	if err := r.db.Where("club_id = ?", clubID).Find(&invitations).Error; err != nil {
		return nil, err
	}
	return invitations, nil
}

func (r *GormTournamentInvitationRepository) Delete(id int) error {
	return r.db.Delete(&models.TournamentInvitation{}, id).Error
}

type GormTournamentTeamRepository struct {
	db *gorm.DB
}

func NewTournamentTeamRepository(db *gorm.DB) *GormTournamentTeamRepository {
	return &GormTournamentTeamRepository{db: db}
}

func (r *GormTournamentTeamRepository) Create(tournamentTeam *models.TournamentTeam) error {
	return r.db.Create(tournamentTeam).Error
}

func (r *GormTournamentTeamRepository) GetByTournamentID(tournamentID int) ([]models.TournamentTeam, error) {
	var tournamentTeams []models.TournamentTeam
	if err := r.db.Where("tournament_id = ?", tournamentID).Find(&tournamentTeams).Error; err != nil {
		return nil, err
	}
	return tournamentTeams, nil
}

func (r *GormTournamentTeamRepository) GetByTeamID(teamID int) ([]models.TournamentTeam, error) {
	var tournamentTeams []models.TournamentTeam
	if err := r.db.Where("team_id = ?", teamID).Find(&tournamentTeams).Error; err != nil {
		return nil, err
	}
	return tournamentTeams, nil
}

func (r *GormTournamentTeamRepository) Delete(id int) error {
	return r.db.Delete(&models.TournamentTeam{}, id).Error
}
