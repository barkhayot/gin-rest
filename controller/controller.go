package controller

import (
	"fetch/helper"
	"fetch/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, model.Albums)
}

func GetAlbumById(c *gin.Context) {
	id := c.Param("id")
	album, err := helper.AlbumById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album by ID not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, album)
}

func PostAlbums(c *gin.Context) {
	var newAlbum model.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.JSON(400, gin.H{"message": "Invalid Body"})
		return
	}
	model.Albums = append(model.Albums, newAlbum)
	c.JSON(200, gin.H{"message": "New data has been added"})
}

func CheckoutAlbum(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if ok == false {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing query ID param"})
	}

	album, err := helper.ValidateData(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album by ID not found"})
		return
	}

	album.Quantity -= 1
	c.IndentedJSON(http.StatusOK, album)

}

func GetAlbumBack(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if ok == false {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing query ID param"})
	}

	album, err := helper.ValidateData(id) //albumById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album by ID not found"})
		return
	}

	album.Quantity += 1
	c.IndentedJSON(http.StatusOK, album)
}
