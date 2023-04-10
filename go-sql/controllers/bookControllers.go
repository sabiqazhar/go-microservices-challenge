package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Author string `json:"author"`
	Price int `json:"price"`
}

func GetAll(ctx *gin.Context)  {
	var res = []Book{}
	var db = ctx.MustGet("db").(*sql.DB)
	
	query := "select * from book"

	rows, err := db.Query(query)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H {
			"error": err.Error(),
		})
	}

	defer rows.Close()

	for rows.Next() {
		var book = Book{}

		err = rows.Scan(&book.Id, &book.Name, &book.Author, &book.Price)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H {
				"error": err.Error(),
			})
			return
		}

		res = append(res, book)
	}

	ctx.JSON(http.StatusOK, res)
}

func GetByID(ctx *gin.Context)  {
	idString := ctx.Param("id")
	var db = ctx.MustGet("db").(*sql.DB)
	var dataBook =  Book{}

	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H {
			"error": err.Error(),
		})
		return
	}

	query := "select * from book where id = $1"

	rows := db.QueryRow(query, id)

	err = rows.Scan(&dataBook.Id, &dataBook.Name, &dataBook.Author, &dataBook.Price)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H {
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dataBook)
}

func CreateBook(ctx *gin.Context) {
	var newBook Book
	var db = ctx.MustGet("db").(*sql.DB)

	err := ctx.ShouldBindJSON(&newBook)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H {
			"error": err.Error(),
		})
	}

	query := "insert into book (name, author, price) values ($1, $2, $3) returning *"
	rows := db.QueryRow(query, newBook.Name, newBook.Author, newBook.Price)

	err = rows.Scan(&newBook.Id, &newBook.Name, &newBook.Author, &newBook.Price)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H {
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, newBook)
}


func UpdateBook(ctx *gin.Context)  {
	idString := ctx.Param("id")
	var db = ctx.MustGet("db").(*sql.DB)
	var updateBook Book

	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H {
			"error": err.Error(),
		})
		return
	}
	
	err = ctx.ShouldBindJSON(&updateBook)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H {
			"error": err.Error(),
		})
		return
	}
	
	query := "update book set name = $1, author = $2, price = $3 where id = $4 returning id"
	rows := db.QueryRow(query, updateBook.Name, updateBook.Author, updateBook.Price, id)

	err = rows.Scan(&updateBook.Id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H {
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, updateBook)
}

func DeleteBook(ctx *gin.Context) {
	idString := ctx.Param("id")
	var db = ctx.MustGet("db").(*sql.DB)
	var deleteBook Book

	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H {
			"error": err.Error(),
		})
		return
	}


	query := "delete from book where id = $1 returning *"
	rows := db.QueryRow(query, id)
	err = rows.Scan(&deleteBook.Id, &deleteBook.Name, &deleteBook.Author, &deleteBook.Price)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H {
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, deleteBook)
}