package main

import (
	"fmt"

	"github.com/Aman-Pailwan/go-crud/initializers"
	"github.com/Aman-Pailwan/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	fmt.Println("Migration Process Started...")
	initializers.DB.AutoMigrate(&models.Book{})
	fmt.Println("Migration Completed")
}
