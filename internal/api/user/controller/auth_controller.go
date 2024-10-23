package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/josevitorrodriguess/chat/internal/api/user/models"
	"github.com/josevitorrodriguess/chat/internal/api/user/service"
)

type authController struct {
	serv service.AuthService
}

func NewAuthController(serv service.AuthService) *authController {
	return &authController{
		serv: serv,
	}
}

func (ac *authController) LoginWithEmail(ctx *gin.Context) {
	var l models.AuthEmail
	if err := ctx.BindJSON(&l); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "cannot bind JSON:" + err.Error()})
	}

	token, err := ac.serv.LoginWithEmail(l.Email, l.Password)
	if err != nil {
		if err.Error() == "cannot find user with this email" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else if err.Error() == "password are incorrect" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
