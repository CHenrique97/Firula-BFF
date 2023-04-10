package controllers

import (
	"log"

	connectDB "github.com/firula-bff/connect"
	"github.com/firula-bff/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Migrate(c *gin.Context) {
	log.Println("Migrating")
	connectDB.DB.AutoMigrate(&models.User{})
	c.JSON(200, gin.H{
		"message": "Database Migrated",
	})
}

func PostCreate(c *gin.Context) {
	var body models.User
	c.BindJSON(&body)

	post := models.User{
		ID:       uuid.New().String(),
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
	}

	result := connectDB.DB.Create(&post)
	if result.Error != nil {

		c.JSON(400, gin.H{
			"message": "User could not be created",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": &post,
	})
}
func getRead(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
func putUpdate(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
func postDelete(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
