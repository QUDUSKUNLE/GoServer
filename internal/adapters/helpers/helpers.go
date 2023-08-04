package helpers

import (
	"os"
	"strconv"
	"time"
	"fmt"
	"strings"
	"errors"
	"github.com/google/uuid"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"server/internal/core/domain"
	"github.com/go-playground/validator/v10"
)


var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

type ErrorMessage struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func GenerateJWToken(user domain.User) (string, error) {
	tokenTTL, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"iat": time.Now().Unix(),
		"eat": time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
	})
	return token.SignedString(privateKey)
}

func ValidateJWToken(token *jwt.Token) (jwt.MapClaims, error) {
	claim, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return claim, nil
	}
	return nil, errors.New("invalid token")
}

func ExtractToken(context *gin.Context) (*jwt.Token, error) {
	bearerToken := context.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	token, err := jwt.Parse(splitToken[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return privateKey, nil
	})
	return token, err
}

func getErrorMessage(message validator.FieldError) string {
	switch message.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + message.Param()
	case "gte":
		return "Should be greater than " + message.Param()
	}
	return "unknown"
}

func CompileErrors(err error) []ErrorMessage {
	var ve validator.ValidationErrors
	var result []ErrorMessage
	if errors.As(err, &ve) {
		for _, fe := range ve {
			result = append(result, ErrorMessage{fe.Field(), getErrorMessage(fe)})
		}
	}
	return result
}

func ValidateUUID(uu string) bool {
	if _, err := uuid.Parse(uu); err != nil {
		return false
	}
	return true
}
