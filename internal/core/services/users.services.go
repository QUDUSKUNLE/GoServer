package services

import (
	"html"
	"strings"
	"github.com/satori/go.uuid"
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
