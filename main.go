package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()

	server.GET("/home", event)

	server.Run(":8080")
}

func event(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}

