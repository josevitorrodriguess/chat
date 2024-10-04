package service

import (
	"log"

	"github.com/josevitorrodriguess/chat/internal/api/user/models"
	"github.com/josevitorrodriguess/chat/internal/api/user/repository"
	"github.com/josevitorrodriguess/chat/internal/utils"
)

type UserService interface {
	Create(*models.User) (models.User, error)
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

	user.Password, _ = utils.HashPass(user.Password)

	if user.Password == "" {
		log.Println("ERROR: fail to encrpyt user password")
	}

	createdUser, err := us.repo.Create(user)
	if err != nil {
		return models.User{}, err
	}

	return createdUser, nil
}
