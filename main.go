package main

import (
	"fmt"

	"github.com/Aman-Pailwan/go-crud/controllers"
	"github.com/Aman-Pailwan/go-crud/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {
	r := gin.Default()
	r.GET("/books", controllers.GetBook)
	r.DELETE("/books/:id", controllers.DeleteBook)
	r.PUT("/books/:id", controllers.UpdateBook)
	r.GET("/books/:id", controllers.GetBookById)
	r.POST("/books", controllers.BookCreate)
	fmt.Println("The Server is started on localhost:8080")
	r.Run()
}
