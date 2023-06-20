package helpers

import (
	"errors"
	"fmt"
	"os"
	"server/models"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

func GenerateJWT(user models.User) (string, error) {
	tokenTTL, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"iat": time.Now().Unix(),
		"eat": time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
	})
	return token.SignedString(privateKey)
}

func ValidateJWT(context *gin.Context) error {
	token, err := getToken(context)
	if err != nil {
		return errors.New("invalid token provided")
	}
	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return nil
	}
	return errors.New("invalid token")
}

func getToken(context *gin.Context) (*jwt.Token, error) {
	tokenString, _ := extractToken(context)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return privateKey, nil
	})
	return token, err
}

func extractToken(context *gin.Context) (string, error) {
	bearerToken := context.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 && splitToken[0] == os.Getenv("TOKEN_TYPE") {
		return splitToken[1], nil
	}
	return "", errors.New("invalid authorization token")
}

func CurrentUser(context *gin.Context) (models.User, error) {
	err := ValidateJWT(context)
	if err != nil {
		return models.User{}, err
	}
	token, _ := getToken(context)
	claims, _ := token.Claims.(jwt.MapClaims)
	userId := claims["id"].(string)

	user, err := models.FindUserByID(userId)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func ValidateUUID(uu string) bool {
	if _, err := uuid.Parse(uu); err != nil {
		return false
	}
	return true
}
