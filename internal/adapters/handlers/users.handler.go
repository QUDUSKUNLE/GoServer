package handlers

import (
	"net/http"
	"server/internal/core/domain"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

// External Interractions
func (service *HTTPHandler) SaveUser(ctx *gin.Context) {
	user := domain.UserDto{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": compileErrors(err) })
	}
	if err := service.ServicesAdapter.SaveUser(
		domain.User{ Email: user.Email, Password: user.Password },
	); err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": "User`s already exists.", "status": false})
		return 
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": "User created successfully.", "status": true})
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": false })
		return
	}
	jwt, err := service.ServicesAdapter.Login(login)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email/Password is incorrect.", "status": false })
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": jwt, "status": true })
}

func (service *HTTPHandler) Telementry(ctx *gin.Context) {
	otelgin.HTML(ctx, http.StatusOK, indexTempl, gin.H{"data": "Telementry"})
}
