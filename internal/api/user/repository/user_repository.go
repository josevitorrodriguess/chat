package repository

import (
	"github.com/josevitorrodriguess/chat/internal/api/user/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) (models.User, error)
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
	createdUser := models.User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}
	return createdUser, nil
}
