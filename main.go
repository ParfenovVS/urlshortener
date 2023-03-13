package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/parfenovvs/urlshortener/handler"
	"github.com/parfenovvs/urlshortener/store"
)

func main() {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hey Go URL shortener!",
		})
	})

	r.POST("/generate", handler.CreateShortUrl)
	r.GET("/:shortUrl", handler.HandleShortUrl)

	storageService := store.InitializeStore()
	defer storageService.PostgresModel.DB.Close()

	err := r.Run(":4000")
	if err != nil {
		panic(fmt.Sprintf("Failed to start server - Error: %v", err))
	}
}
