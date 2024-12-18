package repository

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/josevitorrodriguess/chat/internal/api/user/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(*models.User) (models.User, error)
	FindAll() ([]models.User, error)
	FindById(uuid.UUID) (models.User, error)
	Update(uuid.UUID, models.User) (models.User, error)
	Delete(uuid.UUID) error
	FindByEmail(email string) (*models.User, error)
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

func (ur *userRepository) FindById(id uuid.UUID) (models.User, error) {
	var user models.User

	if err := ur.db.Where("id = ?", id).First(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (ur *userRepository) Update(id uuid.UUID, user models.User) (models.User, error) {

	if err := ur.db.Model(&models.User{}).Where("id = ?", id).Updates(user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (ur *userRepository) Delete(id uuid.UUID) error {
	if err := ur.db.Delete(models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User

	query := "SELECT id, username, email, password FROM users WHERE email = $1"
	result := ur.db.Raw(query, email).Scan(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found")
		}
		return nil, result.Error
	}

	return &user, nil
}
