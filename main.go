package main

import (
	"example.com/database"
	_ "example.com/docs" // replace with the actual package path
	"example.com/routes"
)

// @title Albums API
// @version 1.0
// @description API to manage albums
// @host localhost:8080
// @BasePath /

func main() {
	database.InitDatabase()
	routes.SetupRoutes()

}
