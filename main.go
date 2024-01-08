package main

import (
	"fetch/controller"

	"github.com/gin-gonic/gin"
)

// Routes
func main() {
	router := gin.Default()

	router.GET("/albums", controller.GetAlbums)
	router.GET("/albums/:id", controller.GetAlbumById)
	router.POST("/albums", controller.PostAlbums)
	router.PATCH("/albums/checkout", controller.CheckoutAlbum)
	router.PATCH("/albums/getback", controller.GetAlbumBack)

	router.Run("localhost:8080")
}
