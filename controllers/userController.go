package controllers

import (
	"net/http"

	"example.com/database"
	"example.com/models"
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

// @Summary Get Users
// @Description Get Users
// @Tags Users
// @Produce json
// @Success 200 {object} models.User
// @Failure 404 {object} ErrorResponse
// @Router /users [get]
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&user).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, user)
}

// Get All Users
func GetUsers(c *gin.Context) {
	var users []models.User
	if err := database.DB.Preload("Profile").Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}
