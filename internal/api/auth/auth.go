package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// JWTService define o serviço para geração e validação de JWT
type JWTService struct {
	secretKey []byte
	issuer    string
}

// NewJWTService cria uma nova instância do JWTService
func NewJWTService() *JWTService {
	
	secret := os.Getenv("JWT_SECRET_KEY")
	if secret == "" {
		panic("JWT_SECRET_KEY não foi definido no .env")
	}
	return &JWTService{
		secretKey: []byte(secret),  // Chave como []byte
		issuer:    "your-app-name", // Ajuste conforme necessário
	}
}

// Claim define as informações que estarão no token JWT
type Claim struct {
	UserID uuid.UUID `json:"user_id"`
	jwt.RegisteredClaims
}

// GenerateToken gera um JWT com base no ID do usuário
func (s *JWTService) GenerateToken(userID uuid.UUID) (string, error) {
	claims := &Claim{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    s.issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Expira em 24 horas
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Cria um token com o método de assinatura HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Assina o token usando a chave secreta
	return token.SignedString(s.secretKey)
}

// ValidateToken valida um token JWT e retorna as reivindicações (claims)
func (s *JWTService) ValidateToken(tokenString string) (*Claim, error) {
	// Parse o token usando a chave secreta
	token, err := jwt.ParseWithClaims(tokenString, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		// Verifica se o método de assinatura é HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("método de assinatura inválido")
		}
		return s.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	// Verifica se o token é válido
	if claims, ok := token.Claims.(*Claim); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("token inválido ou expirado")
}

func (s *JWTService) IsValidToken(tokenString string) bool {
	_, err := s.ValidateToken(tokenString)
	return err == nil
}
