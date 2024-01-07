package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Define struct of data
type album struct {
	ID       string  `json:"id"`
	Title    string  `json:"title"`
	Artist   string  `json:"artist"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

// Use as memery storage
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99, Quantity: 12},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99, Quantity: 3},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99, Quantity: 99},
}

// helpers

func albumById(id string) (*album, error) {
	for i, a := range albums {
		if a.ID == id {
			return &albums[i], nil
		}
	}
	return nil, errors.New("Album not found")
}

// Route funcitons

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")
	album, err := albumById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album by ID not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, album)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	res := map[string]string{"message": "created"}

	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Body"})
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, res)
}

func checkoutAlbum(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if ok == false {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing query ID param"})
	}

	album, err := albumById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album by ID not found"})
		return
	}

	if album.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Abbum is out of stock"})
		return
	}

	album.Quantity -= 1
	c.IndentedJSON(http.StatusOK, album)

}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums", postAlbums)
	router.PATCH("/albums", checkoutAlbum)

	router.Run("localhost:8080")
}
