package routes

import (
	"example.com/controllers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes() {
	routes := gin.Default()
	routes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routes.GET("/users", controllers.GetUsers)
	routes.POST("/users", controllers.CreateUser)
	routes.Run("localhost:8080")

}
