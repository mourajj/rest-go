package router

import (
	"rest-go/controller"

	"github.com/gin-gonic/gin"
)

func StartRouter() {
	router := gin.Default()
	router.GET("/albums", controller.GetAlbums)
	router.POST("/albums", controller.PostAlbums)
	router.GET("/albums/:id", controller.GetAlbumsByID)
	router.Run("localhost:8080")
}
