package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Author string `json:"author"`
	Price int `json:"price"`
}

var BookDatas = []Book{}

func CreateBook(ctx *gin.Context)  {
	var newBook Book

	if err :=  ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newBook.Id = fmt.Sprintf("%d", len(BookDatas)+1)
	BookDatas = append(BookDatas, newBook)

	ctx.JSON(http.StatusCreated, gin.H{
		"book": newBook,
	})
}

func UpdateBook(ctx *gin.Context)  {
	bookID := ctx.Param("bookID")
	condition := false
	var updateBook Book

	// fmt.Println(bookID)

	if err := ctx.ShouldBindJSON(&updateBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, book := range BookDatas {
		if bookID == book.Id {
			condition = true
			BookDatas[i] = updateBook
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status" : "data not found",
			"error_message" : fmt.Sprintf("book with id %v not found", bookID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message" : fmt.Sprintf("book with id %v has been updated!", bookID),
	})
}

func GetByID(ctx *gin.Context)  {
	bookID := ctx.Param("bookID")
	condition := false
	var bookData Book

	for i, book := range BookDatas {
		if bookID == book.Id {
			condition = true
			bookData = BookDatas[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status" : "data not found",
			"error_message" : fmt.Sprintf("book with id %v not found", bookID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book": bookData,
	})
}

func GetAll(ctx *gin.Context)  {
	ctx.JSON(http.StatusOK, BookDatas)
}

func DeleteBook(ctx *gin.Context)  {
	bookID := ctx.Param("bookID")
	condition := false
	var bookIndex int

	for i, book := range BookDatas {
		if bookID == book.Id {
			condition = true
			bookIndex = i
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status" : "data not found",
			"error_message" : fmt.Sprintf("book with id %v not found", bookID),
		})
		return
	}

	copy(BookDatas[bookIndex:], BookDatas[bookIndex+1:])
	BookDatas[len(BookDatas)-1] = Book{}
	BookDatas = BookDatas[:len(BookDatas)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"message" : fmt.Sprintf("book with id %v has been deleted!", bookID),
	})
}