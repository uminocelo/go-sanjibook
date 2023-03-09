package main

import (
	"log"
	"net/http"
	"os"
	"sanjibook/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDb()
	loadRoutes()
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(2)
	}
}

func loadDb() {
	database.Connect()
}

func loadRoutes() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.Run(":3000")
}
