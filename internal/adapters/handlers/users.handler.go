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
	if err := service.ExternalServicesAdapter.SaveUser(
		domain.User{ Email: user.Email, Password: user.Password },
	); err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error(), "status": false})
		return 
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": "User created successfully", "status": true})
}

func (service *HTTPHandler) ReadUsers(ctx *gin.Context) {
	result, err := service.ExternalServicesAdapter.ReadUsers()
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
	user, err := service.InternalServicesAdapter.ReadUserByEmail(login.Email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error(), "status": false})
		return
	}
	if err := user.ValidatePassword(login.Password); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": false })
		return 
	}
	jwt, err := service.GenerateJWToken(*user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": service.CompileErrors(err), "status": false })
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": jwt, "status": true })
}

func (service *HTTPHandler) CurrentUser(ctx *gin.Context) (domain.User, error) {
	token, err := service.ExtractToken(ctx)
	if err != nil {
		return domain.User{}, err
	}
	claim, err := service.ValidateJWToken(token);
	if err != nil {
		return domain.User{}, err
	}
	UserID := claim["id"].(string)
	user, err := service.ExternalServicesAdapter.ReadUser(UserID)
	if err != nil {
		return domain.User{}, err
	}
	return *user, nil
}
