package controllers

import (
	"encoding/json"
	"net/http"
	"server/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	newUser := models.UserInput{
		Email:    "quduskunle",
		Password: "test",
	}
	writer := makeRequest("POST", "/v1/users/register", newUser, false, "", "")
	assert.Equal(t, http.StatusCreated, writer.Code)
}

func TestLogin(t *testing.T) {
	user := models.UserInput{
		Email:    "quduskunle",
		Password: "test",
	}

	writer := makeRequest("POST", "/v1/users/login", user, false, "", "")
	assert.Equal(t, http.StatusOK, writer.Code)
	var response map[string]string
	json.Unmarshal(writer.Body.Bytes(), &response)
	_, exists := response["token"]
	assert.Equal(t, true, exists)
}
