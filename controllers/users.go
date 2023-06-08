package controllers

import (
	"server/models"
	"server/helpers"
	"github.com/gin-gonic/gin"
  "net/http"
)

func Register(context *gin.Context) {
	var userInput models.UserInput

	if err := context.ShouldBindJSON(&userInput); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}

	user := models.User{
		Username: userInput.Username,
		Password: userInput.Password,
	}

	savedUser, err := user.Save()
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}

	token, err := helpers.GenerateJWT(*savedUser)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.IndentedJSON(http.StatusCreated, gin.H{ "data": savedUser, "jwt": token })
}

func Login(context *gin.Context) {
	var loginInput models.UserInput

	if err := context.ShouldBindJSON(&loginInput); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}
	var user models.User
	_, err := user.FindUserByUsername(loginInput.Username)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}

	if err := user.ValidatePassword(loginInput.Password); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}

	jwt, err := helpers.GenerateJWT(user)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}
	context.IndentedJSON(http.StatusOK, gin.H{ "token": jwt })
}
