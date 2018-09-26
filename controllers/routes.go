package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/humbertotm/go-api/oauth"
)

// SetupRouter Sets up router
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Test group for JWT tokens
	TestJWT := r.Group("/test-jwt")
	{
		// JWT issuing test endpoint
		TestJWT.GET("/jwt-encode", func(c *gin.Context) {
			// Return a JWT in JSON response
			token, err := oauth.IssueToken()
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err})
			}
			c.JSON(http.StatusOK, gin.H{"jwtToken": token})
		})
		TestJWT.POST("/jwt-decode", func(c *gin.Context) {
			token := c.Request.Header.Get("Authorization")
			fmt.Println(token)
			claims, err := oauth.DecodeToken(token)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err})
			}
			c.JSON(http.StatusOK, gin.H{"claims": claims})
		})
	}

	return r
}
