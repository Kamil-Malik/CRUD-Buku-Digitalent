package router

import (
	"Belajar-Golang-7/controller"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/books", controller.InsertBook)
	router.GET("/books", controller.GetAllBook)
	router.GET("/books/:id", controller.GetBookByID)
	router.PUT("/books", controller.UpdateBook)
	router.DELETE("/books/:id", controller.DeleteBookByID)

	return router
}
