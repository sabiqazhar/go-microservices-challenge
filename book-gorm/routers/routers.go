package routers

import (
	"book-api/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StartServer(db *gorm.DB) *gin.Engine  {
	router := gin.Default()

	router.Use(func(ctx *gin.Context) {
		ctx.Set("db", db)
	})

	router.GET("/books", controllers.GetAll)
	router.GET("/books/:id", controllers.GetByID)
	router.POST("/books", controllers.CreateBook)
	router.PUT("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)

	return router
}