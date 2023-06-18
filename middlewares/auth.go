package middlewares

import (
	"net/http"
	"server/helpers"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)
type States string
const (
	Required States = "required"
)

type ErrorMessage struct {
	Field string `json:"field"`
	Message string `json:"message"`
}

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		if err := helpers.ValidateJWT(context); err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{ "error": err.Error() })
			context.Abort()
			return
		}
		_, err := helpers.CurrentUser(context)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{ "error": "unauthorized request." })
			context.Abort()
			return
		}
		context.Next()
	}
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
