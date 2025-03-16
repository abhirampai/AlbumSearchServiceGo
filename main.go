package main

import (
	"github.com/gin-gonic/gin"
	"album_search_go_service/albums"
	"album_search_go_service/artists"
)

func main() {
	router := gin.Default()
	router.GET("/albums", albums.Index)
	router.GET("/albums/:id", albums.Show)
	router.GET("/artists", artists.Index)
	router.POST("/albums", albums.Create)
	router.PUT("/albums/:id", albums.Update)
	router.DELETE("/albums/:id", albums.Destroy)

	router.Run("localhost:8080")
}
