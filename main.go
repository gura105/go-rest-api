package main

import (
	"github.com/gin-gonic/gin"

	"example/web-service-gin/handler"
)

func main() {
	router := gin.Default()
	router.GET("/albums", handler.GetAlbums)
	router.GET("/albums/:id", handler.GetAlbumByID)
	router.POST("/albums", handler.PostAlbums)

	router.Run("localhost:8080")
}
