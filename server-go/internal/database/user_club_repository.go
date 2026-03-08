package database

import (
	"github.com/uwr-tournament/server-go/internal/models"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{db: db}
}

func (r *GormUserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *GormUserRepository) GetByID(id int) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *GormUserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *GormUserRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *GormUserRepository) Delete(id int) error {
	return r.db.Delete(&models.User{}, id).Error
}

func (r *GormUserRepository) List(offset, limit int) ([]models.User, error) {
	var users []models.User
	if err := r.db.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

type GormClubRepository struct {
	db *gorm.DB
}

func NewClubRepository(db *gorm.DB) *GormClubRepository {
	return &GormClubRepository{db: db}
}

func (r *GormClubRepository) Create(club *models.Club) error {
	return r.db.Create(club).Error
}

func (r *GormClubRepository) GetByID(id int) (*models.Club, error) {
	var club models.Club
	if err := r.db.First(&club, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &club, nil
}

func (r *GormClubRepository) Update(club *models.Club) error {
	return r.db.Save(club).Error
}

func (r *GormClubRepository) Delete(id int) error {
	return r.db.Delete(&models.Club{}, id).Error
}

func (r *GormClubRepository) List(offset, limit int) ([]models.Club, error) {
	var clubs []models.Club
	if err := r.db.Offset(offset).Limit(limit).Find(&clubs).Error; err != nil {
		return nil, err
	}
	return clubs, nil
}
