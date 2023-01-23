package main

import (
	"github.com/coffee-go-api/config"
	"github.com/coffee-go-api/controllers"
	"github.com/coffee-go-api/models"
	"github.com/gin-gonic/gin"
	"fmt"
)

func main() {
	config.GetDatabase()
	err := config.DB.AutoMigrate(&models.Coffee{})
	if err != nil {
		fmt.Println("error")
	}
	r := gin.Default()
	r.POST("/coffees", controllers.AddCoffee)
	r.GET("/coffees/:id", controllers.Get)
	r.GET("/coffees", controllers.GetAll)
	r.DELETE("/coffees/:id/delete", controllers.DeleteByID)
	r.DELETE("/coffees/:id/deletepermamently", controllers.DeletePermamently)
	r.Run("localhost:8080")
}
