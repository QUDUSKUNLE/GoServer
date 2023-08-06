package handlers

import (
	"os"
	"errors"
	"github.com/go-playground/validator/v10"
)

type ErrorMessage struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

func setErrorMessage(message validator.FieldError) string {
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

func compileErrors(err error) []ErrorMessage {
	var ve validator.ValidationErrors
	var result []ErrorMessage
	if errors.As(err, &ve) {
		for _, fe := range ve {
			result = append(result, ErrorMessage{fe.Field(), setErrorMessage(fe)})
		}
	}
	return result
}
