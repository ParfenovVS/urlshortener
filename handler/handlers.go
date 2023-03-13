package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/parfenovvs/urlshortener/shortener"
	"github.com/parfenovvs/urlshortener/store"
)

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	var creationRequest UrlCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := shortener.GenerateShortLink(creationRequest.LongUrl, creationRequest.UserId)
	store.SaveUrlMapping(shortUrl, creationRequest.LongUrl, creationRequest.UserId)

	host := "http://localhost:4000/"
	c.JSON(http.StatusOK, gin.H{
		"short_url": host + shortUrl,
	})
}

func HandleShortUrl(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	initialUrl := store.GetInitialUrl(shortUrl)
	c.Redirect(http.StatusFound, initialUrl)
}
