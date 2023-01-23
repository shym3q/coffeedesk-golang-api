package controllers

import (
	"github.com/coffee-go-api/config"
	"github.com/coffee-go-api/models"
	"github.com/gin-gonic/gin"
)

func AddCoffee(c *gin.Context) {
	var body struct {
		Name        string
		Description string
		Link        string
		Image       string
	}
	err := c.Bind(&body)
	if err != nil {
		return
	}
	coffee := models.Coffee{Name: body.Name, Description: body.Description, Link: body.Link, Image: body.Image}
	result := config.DB.Create(&coffee)
	if result.Error != nil {
		c.Status(409)
		return
	}
	c.JSON(201, coffee)
}

func Get(c *gin.Context) {
	id := c.Param("id")

	var coffee models.Coffee
	result := config.DB.First(&coffee, id)
	if result.Error != nil { 
		c.Status(400)
		return
	}
	c.JSON(200, coffee)
}

func GetAll(c *gin.Context) {
	var coffees []models.Coffee
	result := config.DB.Find(&coffees)
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, coffees)
}

func DeleteByID(c *gin.Context) {
	id := c.Param("id")

	config.DB.Delete(&models.Coffee{}, id)

	c.Status(200)
}

func DeletePermamently(c *gin.Context) {
	id := c.Param("id")

	config.DB.Unscoped().Delete(&models.Coffee{}, id)

	c.Status(200)
}
