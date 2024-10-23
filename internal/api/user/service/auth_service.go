package service

import (
	"errors"
	"fmt"

	"github.com/josevitorrodriguess/chat/internal/api/auth"
	"github.com/josevitorrodriguess/chat/internal/api/user/repository"
	"github.com/josevitorrodriguess/chat/internal/utils"
)

type AuthService interface {
	LoginWithEmail(string, string) (string, error)
}

type authService struct {
	repo repository.UserRepository
}

func NewAuthServie(repo repository.UserRepository) AuthService {
	return &authService{
		repo: repo,
	}
}

func (as *authService) LoginWithEmail(email, password string) (string, error) {
	user, err := as.repo.FindByEmail(email)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", errors.New("cannot find user with this email")
	}

	if !utils.CheckPassword(user.Password, password) {
		return "", fmt.Errorf("password is incorrect")
	}

	token, err := auth.NewJWTService().GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
