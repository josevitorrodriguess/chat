package models

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Username string    `gorm:"unique;not null" validate:"required,min=6"` 
	Email    string    `gorm:"unique;not null"`
	Password string    `gorm:"not null" validate:"required,min=6"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func (u *User) Validate() error {
	return validate.Struct(u)
}
