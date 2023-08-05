package handlers

import (
	"fmt"
	"os"
	"strconv"
	"time"
	services "server/internal/core/services"
	"net/http"
	"errors"
	"strings"
	"github.com/go-playground/validator/v10"
	"server/internal/core/domain"
	"github.com/google/uuid"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gin-gonic/gin"
)

type ErrorMessage struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

type HTTPHandler struct {
	ExternalServicesAdapter  services.ServicesHandler
	InternalServicesAdapter  services.ServicesHandler
}

func HTTPAdapter(services services.ServicesHandler) *HTTPHandler {
	return &HTTPHandler{
		ExternalServicesAdapter: services,
		InternalServicesAdapter: services,
	}
}

func (service *HTTPHandler) JWTAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := service.ExtractToken(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}
		if _, err := service.ValidateJWToken(token); err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

func (service *HTTPHandler) UUidMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		if !service.ValidateUUID(context.Param("id")) {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid id credential"})
			return
		}
		context.Next()
	}
}

func (service *HTTPHandler) ExtractToken(context *gin.Context) (*jwt.Token, error) {
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

func (service *HTTPHandler) ValidateJWToken(token *jwt.Token) (jwt.MapClaims, error) {
	claim, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return claim, nil
	}
	return nil, errors.New("invalid token")
}

func (service *HTTPHandler) ValidateUUID(uu string) bool {
	if _, err := uuid.Parse(uu); err != nil {
		return false
	}
	return true
}

func (service *HTTPHandler) GenerateJWToken(user domain.User) (string, error) {
	tokenTTL, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"iat": time.Now().Unix(),
		"eat": time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
	})
	return token.SignedString(privateKey)
}

func (service *HTTPHandler) SetErrorMessage(message validator.FieldError) string {
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

func (service *HTTPHandler) CompileErrors(err error) []ErrorMessage {
	var ve validator.ValidationErrors
	var result []ErrorMessage
	if errors.As(err, &ve) {
		for _, fe := range ve {
			result = append(result, ErrorMessage{fe.Field(), service.SetErrorMessage(fe)})
		}
	}
	return result
}
