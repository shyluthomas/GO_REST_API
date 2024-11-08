package router

import (
	"example.com/controllers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes() {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/users", controllers.GetUsers)
	router.POST("/users", controllers.CreateUser)
	router.Run("localhost:8080")

}
