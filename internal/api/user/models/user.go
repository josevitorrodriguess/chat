package models

type User struct {
	ID       string `json:"id"`       //MAKES UUID
	Username string `json:"username"` //UNIQUE
	Email    string `json:"email"`    //UNIQUE
	Password string `json:"password"` //SHA256
}
