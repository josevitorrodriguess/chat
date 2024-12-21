package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/josevitorrodriguess/chat/internal/api/user/controller"
	"github.com/josevitorrodriguess/chat/internal/api/user/repository"
	"github.com/josevitorrodriguess/chat/internal/api/user/service"
	"gorm.io/gorm"
)

func MessageRoutes(r *gin.RouterGroup, db *gorm.DB) {
	messageRepo := repository.NewMessageRepository(db)
	messageService := service.NewMessageService(messageRepo)
	messageController := controller.NewMessageController(messageService)

	r.POST("/create", messageController.Create)
	r.GET("/findAll", messageController.GetAll)
	r.GET("/find/:id", messageController.GetByID)
	r.PUT("/update/:id", messageController.Update)
	r.DELETE("/delete/:id", messageController.Delete)
}
