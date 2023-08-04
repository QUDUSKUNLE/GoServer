package handlers

import (
	"net/http"
	"server/internal/adapters/helpers"
	"server/internal/core/domain"

	"github.com/gin-gonic/gin"
)

func (service *HTTPHandler) SaveUser(ctx *gin.Context) {
	user := domain.UserInputDto{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": helpers.CompileErrors(err) })
	}
	if err := service.svc.SaveUser(
		domain.User{ Email: user.Email, Password: user.Password },
	); err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error() })
		return 
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": "User created successfully"})
}

func (service *HTTPHandler) ReadUsers(ctx *gin.Context) {
	result, err := service.svc.ReadUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "error": helpers.CompileErrors(err)})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": result })
}

func (service *HTTPHandler) Login(ctx *gin.Context) {
	login := domain.UserInputDto{}
	if err := ctx.ShouldBindJSON(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": helpers.CompileErrors(err) })
		return
	}
	user, err := service.svc.ReadUserByEmail(login.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": helpers.CompileErrors(err)})
		return
	}
	if err := user.ValidatePassword(login.Password); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error() })
		return 
	}
	jwt, err := helpers.GenerateJWT(*user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": helpers.CompileErrors(err) })
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": jwt})
}
