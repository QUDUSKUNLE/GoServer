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
		Email:    "qudus@gmail.com",
		Password: "test12345&",
	}
	writer := makeRequest("POST", "/v1/users/register", newUser, false, "", "")
	var responseBody map[string]interface{}
	json.Unmarshal(writer.Body.Bytes(), &responseBody)
	data, exists := responseBody["data"]
	assert.NotEmpty(t, data)
	assert.Equal(t,true, exists)
	assert.IsType(t, responseBody["data"], "string")
	assert.Equal(t, http.StatusCreated, writer.Code)
}

func TestLogin(t *testing.T) {
	user := models.UserInput{
		Email:    "qudus@gmail.com",
		Password: "test12345&",
	}

	writer := makeRequest("POST", "/v1/users/login", user, false, "", "")
	var response map[string]interface{}
	json.Unmarshal(writer.Body.Bytes(), &response)
	data, exists := response["token"]
	assert.NotEmpty(t, data)
	assert.Equal(t, true, exists)
	assert.Equal(t, http.StatusOK, writer.Code)
	assert.IsType(t, response["token"], "string")
}
