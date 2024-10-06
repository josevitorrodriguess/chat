package service

import (
	"fmt"

	"github.com/josevitorrodriguess/chat/internal/api/user/models"
	"github.com/josevitorrodriguess/chat/internal/api/user/repository"
	"github.com/josevitorrodriguess/chat/internal/utils"
)

type UserService interface {
	Create(*models.User) (models.User, error)
	FindAll() ([]models.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (us *userService) Create(user *models.User) (models.User, error) {

	if err := user.Validate(); err != nil {
		return models.User{}, fmt.Errorf("validation failed: %w", err)
	}

	hashedPassword, err := utils.HashPass(user.Password)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to encrypt password: %w", err)
	}
	user.Password = hashedPassword

	createdUser, err := us.repo.Create(user)
	if err != nil {
		return models.User{}, err
	}

	return createdUser, nil
}

func (us *userService) FindAll() ([]models.User, error) {
	return us.repo.FindAll()
}
