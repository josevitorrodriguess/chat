package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/josevitorrodriguess/chat/internal/api/auth"
)

func Auth(jwtService *auth.JWTService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BearerSchema = "Bearer "

		// Obter o cabeçalho Authorization da requisição
		header := ctx.GetHeader("Authorization")
		if header == "" || !strings.HasPrefix(header, BearerSchema) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid Authorization header"})
			ctx.Abort()
			return
		}

		// Extrair o token JWT removendo o prefixo "Bearer "
		token := strings.TrimPrefix(header, BearerSchema)

		// Validar o token JWT usando o serviço JWT
		claims, err := jwtService.ValidateToken(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid JWT", "details": err.Error()})
			ctx.Abort()
			return
		}

		// Adicionar as claims ao contexto
		ctx.Set("userID", claims.UserID.String())

		// Continuar para o próximo middleware ou manipulador
		ctx.Next()
	}
}
