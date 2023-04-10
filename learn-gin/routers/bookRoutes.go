package routers

import (
	"learn-gin/controllers"
	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine  {
	router := gin.Default()

	router.GET("/books", controllers.GetAll)
	router.GET("/books/:bookID", controllers.GetByID)
	router.POST("/books", controllers.CreateBook)
	router.PUT("/books/:bookID", controllers.UpdateBook)
	router.DELETE("/books/:bookID", controllers.DeleteBook)

	return router
}