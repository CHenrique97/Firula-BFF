package controllers

import (
	"log"

	connectDB "github.com/firula-bff/connect"
	"github.com/firula-bff/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Check struct {
	Result bool
}

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

	var check Check

	err := connectDB.DB.Raw("SELECT EXISTS(SELECT 1 FROM `users` WHERE `email` = ?) as result", body.Email).Scan(&check).Error
	log.Println(check)
	log.Println(err)
	if err != nil {
		// handle error
		c.JSON(400, gin.H{
			"message": "User could not be verified",
		})
		return
	}

	if check.Result {
		c.JSON(400, gin.H{
			"message": "User could not be created",
		})
		return
	}
	// use the `result` variable
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
