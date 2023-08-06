package services

import (
	"os"
	"strconv"
	"time"
	"html"
	"strings"
	"github.com/satori/go.uuid"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	domain "server/internal/core/domain"
)

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

func (externalServiceHandler *ServicesHandler) SaveUser(user domain.User) error {
	user.ID = uuid.NewV4()
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	user.Email = html.EscapeString(strings.TrimSpace(user.Email))
	return externalServiceHandler.External.SaveUser(user)
}

func (externalServiceHandler *ServicesHandler) ReadUser(UserID string) (*domain.User, error) {
	return externalServiceHandler.External.ReadUser(UserID)
}

func (externalServiceHandler *ServicesHandler) ReadUsers() ([]*domain.User, error) {
	return externalServiceHandler.External.ReadUsers()
}

func (externalServiceHandler *ServicesHandler) ReadUserByEmail(Email string) (*domain.User, error) {
	return externalServiceHandler.External.ReadUserByEmail(Email)
}

func (externalServiceHandler *ServicesHandler) Login(user domain.UserDto) (string, error) {
	userByEmail, err := externalServiceHandler.External.ReadUserByEmail(user.Email)
	if err != nil {
		return "", err
	}
	if err := userByEmail.ValidatePassword(user.Password); err != nil {
		return "", err
	}
	return externalServiceHandler.generateJWToken(userByEmail)
}

func (externalServiceHandler *ServicesHandler) generateJWToken(user *domain.User) (string, error) {
	tokenTTL, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"iat": time.Now().Unix(),
		"eat": time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
	})
	return token.SignedString(privateKey)
}
