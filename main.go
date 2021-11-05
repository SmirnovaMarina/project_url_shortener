package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"math/rand"
	"project-url-shortener/handler"
	"project-url-shortener/store"
	"time"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the URL Shortener API",
		})
	})
	router.POST("/create-short-url", func(c *gin.Context) {
		handler.HandleShortUrlCreate(c)
	})
	router.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	rand.Seed(time.Now().UnixNano())
	store.InitializeStore()

	err = router.Run(":8080")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
