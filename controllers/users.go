package controllers

import (
	"net/http"
	"server/helpers"
	"server/middlewares"
	"server/models"

	"github.com/gin-gonic/gin"
)

func Register(context *gin.Context) {
	userInputModel := models.UserInput{}

	if err := context.ShouldBindJSON(&userInputModel); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": middlewares.CompileErrors(err)})
		return
	}

	userModel := models.User{
		Email:    userInputModel.Email,
		Password: userInputModel.Password,
	}

	if err := userModel.Save(); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": middlewares.CompileErrors(err) })
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": "User created successfully"})
}

func Login(context *gin.Context) {
	loginInputModel := models.UserInput{}

	if err := context.ShouldBindJSON(&loginInputModel); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": middlewares.CompileErrors(err)})
		return
	}

	userModel := models.User{}
	_, err := userModel.FindUserByEmail(loginInputModel.Email)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": middlewares.CompileErrors(err)})
		return
	}

	if err := userModel.ValidatePassword(loginInputModel.Password); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": middlewares.CompileErrors(err)})
		return
	}

	jwt, err := helpers.GenerateJWT(userModel)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": middlewares.CompileErrors(err)})
		return
	}
	context.JSON(http.StatusOK, gin.H{"token": jwt})
}
