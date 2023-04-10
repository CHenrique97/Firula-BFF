package main

import (
	"fmt"

	connectDB "github.com/firula-bff/connect"
	"github.com/firula-bff/controllers"
	"github.com/firula-bff/initializers"
	"github.com/gin-gonic/gin"
)

// Main function
func init() {
	initializers.LoadEnv()
	connectDB.InitConnector()
}

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

	r.GET("/migrate", controllers.Migrate)
	r.POST("/postUser", controllers.PostCreate)
	r.Run(":8080")
	fmt.Println("Hello World")
}
