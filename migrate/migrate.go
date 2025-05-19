package main

import (
	"github.com/Divyshekhar/go-crud/initializers"
	"github.com/Divyshekhar/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
