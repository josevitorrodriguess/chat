package models

type User struct {
	ID       string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"` // UUID como chave primária
	Username string `gorm:"unique;not null"`                                 // Nome de usuário único e não nulo
	Email    string `gorm:"unique;not null"`                                 // E-mail único e não nulo
	Password string `gorm:"not null"`                                        // Senha não nula
}
