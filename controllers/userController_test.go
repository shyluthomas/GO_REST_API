package controllers

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"encoding/json"

	"example.com/database"
	"example.com/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func SetupTestDB() {
	database.DB = database.InitTestDB(":memmory")
}
func TestUserController(t *testing.T) {
	// define test db
	SetupTestDB()

	//set the test mode for gin
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.POST("/users", CreateUser)
	email := uuid.New()
	user := models.User{
		Name:  "shylu",
		Email: email.String() + `@gmail.com`,
		Profile: models.Profile{
			Bio: "Software developer from California",
		},
	}
	jsonData, err := json.Marshal(user)
	assert.NoError(t, err)

	// create http request
	req, err := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonData))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	respRecorder := httptest.NewRecorder()
	router.ServeHTTP(respRecorder, req)

	// test the request
	assert.Equal(t, http.StatusCreated, respRecorder.Code)

	// test the response
	var responseUser models.User
	err = json.Unmarshal(respRecorder.Body.Bytes(), &responseUser)
	assert.NoError(t, err)
	fmt.Printf("dddd %v", responseUser)
	assert.Equal(t, responseUser.Name, user.Name)

}
