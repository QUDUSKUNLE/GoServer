package handlers

import (
	"net/http"
	"server/internal/core/domain"

	"github.com/gin-gonic/gin"
)

func (service *HTTPHandler) SaveUser(ctx *gin.Context) {
	var user domain.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error() })
	}

	if err := service.svc.SaveUser(user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error() })
		return 
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": "User created successfully"})
}
