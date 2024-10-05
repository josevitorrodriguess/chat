package api

import (
	"github.com/gin-gonic/gin"
	"github.com/josevitorrodriguess/chat/internal/api/user/routes"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Grupo principal /user
	userGroup := r.Group("/user")
	{
		routes.UserRoutes(userGroup, db) // Passa o banco de dados para as rotas de usuário
	}

	return r
}
