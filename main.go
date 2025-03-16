package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

var albums = []album {
	{ ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99 },
	{ ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99 },
	{ ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99 },
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbum(c *gin.Context) {
	id := c.Param("id")

	for _, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{ "message": "Album not found" })
}

func getArtists(c *gin.Context) {
	var artists = []string{}
	for	_, album := range albums {
		artists = append(artists, album.Artist)
	}

	c.IndentedJSON(http.StatusOK, gin.H{ "artists": artists })
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func updateAlbums(c *gin.Context) {
	id := c.Param("id")
	var updateAlbum struct {
		Title string `json:"title"`
		Artist string `json:"artist"`
		Price float64 `json:"price"`
	}

	if err := c.BindJSON(&updateAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for index, album := range albums {
		if album.ID == id {
			if updateAlbum.Title != "" {
				albums[index].Title = updateAlbum.Title
			}

			if updateAlbum.Artist != "" {
				albums[index].Artist = updateAlbum.Artist
			}

			if updateAlbum.Price != 0 {
				albums[index].Price = updateAlbum.Price
			}

			c.IndentedJSON(http.StatusOK, gin.H{ "message": "Updated album successfully", "album": albums[index] })
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{ "message": "Album not found" })
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbum)
	router.GET("/artists", getArtists)
	router.POST("/albums", postAlbums)
	router.PUT("/albums/:id", updateAlbums)

	router.Run("localhost:8080")
}
