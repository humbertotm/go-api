package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRouter Sets up router
func SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Test route grouping
	G1 := r.Group("/group1")
	{
		// Ping test
		G1.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong mofos")
		})
	}

	return r
}
