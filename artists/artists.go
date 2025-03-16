package artists

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"album_search_go_service/models"
)

func Index(c *gin.Context) {
	var artists = []string{}
	for	_, album := range models.Albums {
		artists = append(artists, album.Artist)
	}

	c.IndentedJSON(http.StatusOK, gin.H{ "artists": artists })
}
