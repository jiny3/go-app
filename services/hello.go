package services

import "github.com/gin-gonic/gin"

func Hello(c *gin.Context) {
	// This is a simple handler that returns a greeting message.
	c.String(200, "Hello, World!")
}
