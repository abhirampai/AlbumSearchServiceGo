package albums

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"album_search_go_service/models"
)

func Index(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Albums)
}

func Show(c *gin.Context) {
	id := c.Param("id")

	for _, album := range models.Albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{ "message": "Album not found" })
}

func Create(c *gin.Context) {
	var newAlbum models.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.Albums = append(models.Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func Update(c *gin.Context) {
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

	for index, album := range models.Albums {
		if album.ID == id {
			if updateAlbum.Title != "" {
				models.Albums[index].Title = updateAlbum.Title
			}

			if updateAlbum.Artist != "" {
				models.Albums[index].Artist = updateAlbum.Artist
			}

			if updateAlbum.Price != 0 {
				models.Albums[index].Price = updateAlbum.Price
			}

			c.IndentedJSON(http.StatusOK, gin.H{ "message": "Updated album successfully", "album": models.Albums[index] })
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{ "message": "Album not found" })
}