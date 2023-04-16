package router

import (
	"Belajar-Golang-7/controller"
	"Belajar-Golang-7/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title           Book CRUD API
// @version         1.0
// @description     Challenge digitalent FGA
// @termsOfService  http://swagger.io/terms/

// @contact.name   LeleStacia
// @contact.url    https://wa.me/628987654321
// @contact.email  km8003296@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func StartServer() *gin.Engine {
	router := gin.Default()
	docs.SwaggerInfo.Title = "Book API Digitalent Challenge"
	router.POST("/books", controller.InsertBook)
	router.GET("/books", controller.GetAllBook)
	router.GET("/books/:id", controller.GetBookByID)
	router.PUT("/books", controller.UpdateBook)
	router.DELETE("/books/:id", controller.DeleteBookByID)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")

	return router
}
