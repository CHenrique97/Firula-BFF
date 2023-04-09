package controllers

import (
	"log"

	connectDB "github.com/firula-bff/connect"
	"github.com/firula-bff/models"
	"github.com/gin-gonic/gin"
)

func Migrate(c *gin.Context) {
	log.Println("Migrating")
	connectDB.DB.AutoMigrate(&models.User{})
	c.JSON(200, gin.H{
		"message": "Database Migrated",
	})
}

func PostCreate(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
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
