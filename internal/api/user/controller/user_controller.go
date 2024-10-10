package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/josevitorrodriguess/chat/internal/api/user/models"
	"github.com/josevitorrodriguess/chat/internal/api/user/service"
	"gorm.io/gorm"
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

func (uc *userController) FinAll(ctx *gin.Context) {

	users, err := uc.serv.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (uc *userController) FindById(ctx *gin.Context) {

	idParam := ctx.Param("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	user, err := uc.serv.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (uc *userController) Update(ctx *gin.Context) {
	var user models.User

	idParam := ctx.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedUser, err := uc.serv.Update(id, user)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	ctx.JSON(http.StatusOK, updatedUser)
}

func (uc *userController) DeleteUser(ctx *gin.Context) {

    idParam := ctx.Param("id")

    id, err := uuid.Parse(idParam)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
        return
    }

    if err := uc.serv.Delete(id); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
