package models

type AuthUserName struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthEmail struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}