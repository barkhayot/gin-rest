package controller

import (
	"fetch/helper"
	"fetch/model"

	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	c.JSON(200, model.Albums)
}

func GetAlbumById(c *gin.Context) {
	id := c.Param("id")
	album, err := helper.AlbumById(id)

	if err != nil {
		c.JSON(404, gin.H{"message": "Album by ID not found"})
		return
	}
	c.JSON(200, album)
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
		c.JSON(400, gin.H{"message": "missing query ID param"})
	}

	album, err := helper.ValidateData(id)

	if err != nil {
		c.JSON(404, gin.H{"message": "Album by ID not found"})
		return
	}

	album.Quantity -= 1
	c.JSON(200, album)

}

func GetAlbumBack(c *gin.Context) {
	id, ok := c.GetQuery("id")
	title, ok2 := c.GetQuery("title")

	if ok == false || ok2 == false {
		c.JSON(400, gin.H{"message": "missing query ID param"})
		return
	}

	album, err := helper.ValidateData(id) //albumById(id)

	if err != nil {
		c.JSON(404, gin.H{"message": "Album by ID not found"})
		return
	}

	album.Quantity += 1
	album.Title = title
	c.JSON(200, album)
}
