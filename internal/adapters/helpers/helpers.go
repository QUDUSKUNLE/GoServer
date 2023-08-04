package helpers

import (
	"os"
	"strconv"
	"time"
	"fmt"
	"strings"
	"errors"
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

func GenerateJWT(user domain.User) (string, error) {
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

// func CurrentUser(context *gin.Context) (domain.User, error) {
// 	err := ValidateJWT(context)
// 	if err != nil {
// 		return domain.User{}, err
// 	}
// 	token, _ := getToken(context)
// 	claims, _ := token.Claims.(jwt.MapClaims)
// 	userId := claims["id"].(string)

// 	user, err := domain.User.FindUserByID(userId)
// 	if err != nil {
// 		return domain.User{}, err
// 	}
// 	return user, nil
// }
