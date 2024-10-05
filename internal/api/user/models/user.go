package models

import (
	"github.com/go-playground/validator/v10"
)

type User struct {
	ID       string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Username string `gorm:"unique;not null" validate:"required,min=6"` // Validação: requer ao menos 1 caractere
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null" validate:"required,min=6`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func (u *User) Validate() error {
	return validate.Struct(u)
}
