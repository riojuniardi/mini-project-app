package main

import (
	"log"
	"mini-project/config"
	"mini-project/controllers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDB()

	server := gin.Default()

	server.POST("/movies", controllers.CreateBioskop)
	server.GET("/movies", controllers.GetBioskops)
	server.GET("/movies/:id", controllers.GetBioskopById)
	server.PUT("/movies/:id", controllers.UpdateBioskop)
	server.DELETE("/movies/:id", controllers.DeleteEvent)

	server.Run(":8080")
}
