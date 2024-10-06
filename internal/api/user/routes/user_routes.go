package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/josevitorrodriguess/chat/internal/api/user/controller"
	"github.com/josevitorrodriguess/chat/internal/api/user/repository"
	"github.com/josevitorrodriguess/chat/internal/api/user/service"
	"gorm.io/gorm"
)

func UserRoutes(r *gin.RouterGroup, db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	r.POST("/signUp", userController.CreateUser)
	r.GET("/findAll", userController.FinAll)
	r.GET("/find/:id", userController.FindById)
}
