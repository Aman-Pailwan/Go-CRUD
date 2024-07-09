package controllers

import (
	"github.com/Aman-Pailwan/go-crud/initializers"
	"github.com/Aman-Pailwan/go-crud/models"
	"github.com/gin-gonic/gin"
)

func GetBook(c *gin.Context) {
	var books []models.Book

	initializers.DB.Find(&books)

	c.JSON(200, gin.H{
		"message": books,
	})
}

func GetBookById(c *gin.Context) {
	var books models.Book
	id := c.Param("id")

	initializers.DB.First(&books, id)
	c.JSON(200, gin.H{
		"message": books,
	})

}

func BookCreate(c *gin.Context) {
	var body struct {
		ID     string
		Name   string
		Author string
	}

	c.Bind(&body)
	book := models.Book{
		ID:     body.ID,
		Name:   body.Name,
		Author: body.Author,
	}

	result := initializers.DB.Create(&book)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"Message": book,
	})
}

func UpdateBook(c *gin.Context) {
	// getting ID
	id := c.Param("id")
	// Creating Matching Struct
	var newBook struct {
		Name   string
		Author string
	}

	c.Bind(&newBook)

	//find the book to update
	var book models.Book
	initializers.DB.First(&book, id)

	initializers.DB.Model(&book).Updates(models.Book{
		Name:   newBook.Name,
		Author: newBook.Author,
	})

	c.JSON(201, gin.H{
		"Updated Book": book,
	})
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Book{}, id)
	c.Status(200)

}
