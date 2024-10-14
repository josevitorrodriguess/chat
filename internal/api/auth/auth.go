package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWT(issuer string) *jwtService {
	return &jwtService{
		secretKey: os.Getenv("JWT_SECRET_KEY"),
		issuer:    "chat_api",
	}
}

type Claim struct {
	Sum uuid.UUID `json:"sum"`
	jwt.StandardClaims
}

func (s *jwtService) GenerateToken(id uuid.UUID) (string, error) {

	claim := &Claim{
		Sum: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    s.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claim)

	tokenString, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *jwtService) ValidateToken(tokenString string) bool {

	token, err := jwt.ParseWithClaims(tokenString, &Claim{}, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid token")
		}
		return []byte(s.secretKey), nil

	})

	return err == nil && token.Valid
}

func (s *jwtService) GetClaims(tokenString string) (Claim, error) {
	claims := Claim{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		// Verifica se o método de assinatura do token é HMAC (HS256 neste caso)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid token")
		}
		// Retorna a chave secreta para validar o token
		return []byte(s.secretKey), nil
	})

	if err != nil {
		return Claim{}, err
	}

	if !token.Valid {
		return Claim{}, fmt.Errorf("token is not valid")
	}

	return claims, nil
}
