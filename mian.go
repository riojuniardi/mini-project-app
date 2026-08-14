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

	server.POST("/bioskop", controllers.CreateBioskop)
	server.GET("/bioskop", controllers.GetBioskops)
	server.GET("/bioskop/:id", controllers.GetBioskopById)
	server.PUT("/bioskop/:id", controllers.UpdateBioskop)
	server.DELETE("/bioskop/:id", controllers.DeleteEvent)

	server.Run(":8080")
}
