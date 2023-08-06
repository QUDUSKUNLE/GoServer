package handlers

import (
	"net/http"
	"server/internal/core/domain"
	"github.com/gin-gonic/gin"
)

// External Interractions
func (service *HTTPHandler) SaveUser(ctx *gin.Context) {
	user := domain.UserDto{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": service.CompileErrors(err) })
	}
	if err := service.ServicesAdapter.SaveUser(
		domain.User{ Email: user.Email, Password: user.Password },
	); err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error(), "status": false})
		return 
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": "User created successfully", "status": true})
}

func (service *HTTPHandler) ReadUsers(ctx *gin.Context) {
	result, err := service.ServicesAdapter.ReadUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "error": err.Error(), "status": false})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": result })
}

func (service *HTTPHandler) Login(ctx *gin.Context) {
	login := domain.UserDto{}
	if err := ctx.ShouldBindJSON(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": service.CompileErrors(err), "status": false })
		return
	}
	jwt, err := service.ServicesAdapter.Login(login)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": service.CompileErrors(err), "status": false })
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": jwt, "status": true })
}
