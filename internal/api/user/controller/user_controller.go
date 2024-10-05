package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/josevitorrodriguess/chat/internal/api/user/models"
	"github.com/josevitorrodriguess/chat/internal/api/user/service"
)

type userController struct {
	serv service.UserService
}

func NewUserController(serv service.UserService) *userController {
	return &userController{
		serv: serv,
	}
}

func (uc *userController) CreateUser(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := uc.serv.Create(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusCreated, createdUser)
}
