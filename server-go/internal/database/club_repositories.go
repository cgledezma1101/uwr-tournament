package database

import (
	"github.com/uwr-tournament/server-go/internal/models"
	"gorm.io/gorm"
)

type GormClubAdminRepository struct {
	db *gorm.DB
}

func NewClubAdminRepository(db *gorm.DB) *GormClubAdminRepository {
	return &GormClubAdminRepository{db: db}
}

func (r *GormClubAdminRepository) Create(admin *models.ClubAdmin) error {
	return r.db.Create(admin).Error
}

func (r *GormClubAdminRepository) GetByClubID(clubID int) ([]models.ClubAdmin, error) {
	var admins []models.ClubAdmin
	if err := r.db.Where("club_id = ?", clubID).Find(&admins).Error; err != nil {
		return nil, err
	}
	return admins, nil
}

func (r *GormClubAdminRepository) DeleteByUserAndClub(userID, clubID int) error {
	return r.db.Where("user_id = ? AND club_id = ?", userID, clubID).Delete(&models.ClubAdmin{}).Error
}

type GormClubJoinRequestRepository struct {
	db *gorm.DB
}

func NewClubJoinRequestRepository(db *gorm.DB) *GormClubJoinRequestRepository {
	return &GormClubJoinRequestRepository{db: db}
}

func (r *GormClubJoinRequestRepository) Create(request *models.ClubJoinRequest) error {
	return r.db.Create(request).Error
}

func (r *GormClubJoinRequestRepository) GetByID(id int) (*models.ClubJoinRequest, error) {
	var request models.ClubJoinRequest
	if err := r.db.First(&request, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &request, nil
}

func (r *GormClubJoinRequestRepository) GetByClubID(clubID int) ([]models.ClubJoinRequest, error) {
	var requests []models.ClubJoinRequest
	if err := r.db.Where("club_id = ?", clubID).Find(&requests).Error; err != nil {
		return nil, err
	}
	return requests, nil
}

func (r *GormClubJoinRequestRepository) Delete(id int) error {
	return r.db.Delete(&models.ClubJoinRequest{}, id).Error
}

type GormInvitationRepository struct {
	db *gorm.DB
}

func NewInvitationRepository(db *gorm.DB) *GormInvitationRepository {
	return &GormInvitationRepository{db: db}
}

func (r *GormInvitationRepository) Create(invitation *models.Invitation) error {
	return r.db.Create(invitation).Error
}

func (r *GormInvitationRepository) GetByID(id int) (*models.Invitation, error) {
	var invitation models.Invitation
	if err := r.db.First(&invitation, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &invitation, nil
}

func (r *GormInvitationRepository) GetByUserID(userID int) ([]models.Invitation, error) {
	var invitations []models.Invitation
	if err := r.db.Where("user_id = ?", userID).Find(&invitations).Error; err != nil {
		return nil, err
	}
	return invitations, nil
}

func (r *GormInvitationRepository) GetByClubID(clubID int) ([]models.Invitation, error) {
	var invitations []models.Invitation
	if err := r.db.Where("club_id = ?", clubID).Find(&invitations).Error; err != nil {
		return nil, err
	}
	return invitations, nil
}

func (r *GormInvitationRepository) Delete(id int) error {
	return r.db.Delete(&models.Invitation{}, id).Error
}

type GormUserClubRepository struct {
	db *gorm.DB
}

func NewUserClubRepository(db *gorm.DB) *GormUserClubRepository {
	return &GormUserClubRepository{db: db}
}

func (r *GormUserClubRepository) Create(userClub *models.UserClub) error {
	return r.db.Create(userClub).Error
}

func (r *GormUserClubRepository) GetByUserID(userID int) ([]models.UserClub, error) {
	var userClubs []models.UserClub
	if err := r.db.Where("user_id = ?", userID).Find(&userClubs).Error; err != nil {
		return nil, err
	}
	return userClubs, nil
}

func (r *GormUserClubRepository) GetByClubID(clubID int) ([]models.UserClub, error) {
	var userClubs []models.UserClub
	if err := r.db.Where("club_id = ?", clubID).Find(&userClubs).Error; err != nil {
		return nil, err
	}
	return userClubs, nil
}

func (r *GormUserClubRepository) DeleteByUserAndClub(userID, clubID int) error {
	return r.db.Where("user_id = ? AND club_id = ?", userID, clubID).Delete(&models.UserClub{}).Error
}
