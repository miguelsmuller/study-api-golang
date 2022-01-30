package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

func getAlbums(cntx *gin.Context) {
	cntx.IndentedJSON(http.StatusOK, albums)
	//cntx.JSON(http.StatusOK, albums)
}

func getAlbumByID(cntx *gin.Context) {
	id := cntx.Param("id")

	for _, a := range albums {
		if a.ID == id {
			cntx.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	cntx.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func postAlbums(cntx *gin.Context) {
	var newAlbum album

	if err := cntx.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	cntx.IndentedJSON(http.StatusCreated, newAlbum)
}
