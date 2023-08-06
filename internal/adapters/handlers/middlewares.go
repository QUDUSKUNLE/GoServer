package handlers

import (
	"fmt"
	"net/http"
	"errors"
	"strings"
	"github.com/google/uuid"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gin-gonic/gin"
)

func (service *HTTPHandler) JWTAuthentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := extractToken(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}
		cla, err := validateJWToken(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}
		UserID := cla["id"].(string)
		user, err := service.ServicesAdapter.ReadUser(UserID)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}
		ctx.Set("user", user)
		ctx.Next()
	}
}

func (service *HTTPHandler) UUIDMiddleware() gin.HandlerFunc {
	return func (ctx *gin.Context) {
		if !validateUUID(ctx.Param("id")) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid id credential"})
			return
		}
		ctx.Next()
	}
}

func extractToken(ctx *gin.Context) (*jwt.Token, error) {
	bearerToken := ctx.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	token, err := jwt.Parse(splitToken[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return privateKey, nil
	})
	return token, err
}

func validateJWToken(token *jwt.Token) (jwt.MapClaims, error) {
	claim, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return claim, nil
	}
	return nil, errors.New("invalid token")
}

func validateUUID(str string) bool {
	if _, err := uuid.Parse(str); err != nil {
		return false
	}
	return true
}

