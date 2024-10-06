package repository

import (
	"github.com/josevitorrodriguess/chat/internal/api/user/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(*models.User) (models.User, error)
	FindAll() ([]models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) Create(user *models.User) (models.User, error) {
	if err := ur.db.Create(user).Error; err != nil {
		return models.User{}, err
	}
	return *user, nil
}

func (ur *userRepository) FindAll() ([]models.User, error) {
	var users []models.User

	if err := ur.db.Find(&users).Error; err != nil {
		return []models.User{}, err
	}

	return users, nil
}
