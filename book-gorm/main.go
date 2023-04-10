package main

import (
	"book-api/database"
	"book-api/models"
	"book-api/routers"
)

func main()  {
	db := database.StartDB()
	db.AutoMigrate(&models.Book{})

	g := routers.StartServer(db)

	g.Run(":8000")
}