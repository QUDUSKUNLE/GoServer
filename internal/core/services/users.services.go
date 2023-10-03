package services

import (
	"strconv"
	"time"
	"fmt"
	"html"
	"strings"
	"github.com/satori/go.uuid"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	domain "server/internal/core/domain"
)

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
	return externalServiceHandler.Internal.SaveUser(user)
}

func (externalServiceHandler *ServicesHandler) ReadUser(UserID string) (*domain.User, error) {
	return externalServiceHandler.Internal.ReadUser(UserID)
}

func (externalServiceHandler *ServicesHandler) ReadUsers() ([]*domain.User, error) {
	return externalServiceHandler.Internal.ReadUsers()
}

func (externalServiceHandler *ServicesHandler) ReadUserByEmail(Email string) (*domain.User, error) {
	return externalServiceHandler.Internal.ReadUserByEmail(Email)
}

func (externalServiceHandler *ServicesHandler) Login(user domain.UserDto) (string, error) {
	userByEmail, err := externalServiceHandler.Internal.ReadUserByEmail(user.Email)
	fmt.Println(userByEmail, err, "@@@@@@@")
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	if err := userByEmail.ValidatePassword(user.Password); err != nil {
		return "", err
	}
	return externalServiceHandler.generateJWToken(userByEmail)
}

func (externalServiceHandler *ServicesHandler) generateJWToken(user *domain.User) (string, error) {
	tokenTTL, _ := strconv.Atoi(tokenKey)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"iat": time.Now().Unix(),
		"eat": time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
	})
	return token.SignedString(privateKey)
}
