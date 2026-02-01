package main

import (
	"blog-platform/database"
	"blog-platform/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	routes.RegisterRoutes(r)

	r.Run(":8080")

}
