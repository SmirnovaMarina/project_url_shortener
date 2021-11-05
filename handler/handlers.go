package handler

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"project-url-shortener/encoder"
	"project-url-shortener/store"
)

type UrlEncodingRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
}

func HandleShortUrlCreate(c *gin.Context) {
	var request UrlEncodingRequest

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var id int64
	for used := true; used; used = store.IsUrlUnique(id) {
		id = rand.Int63()
	}

	shortUrl := encoder.Encode(id)
	store.SaveUrlMapping(id, request.LongUrl)

	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": shortUrl,
	})
}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	id := encoder.Decode(shortUrl)
	originalUrl := store.RetrieveOriginalUrl(id)
	c.Redirect(302, originalUrl)
}
