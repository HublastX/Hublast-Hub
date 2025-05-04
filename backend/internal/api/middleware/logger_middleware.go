package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// LoggerMiddleware logs the details of incoming requests
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Process the request
		c.Next()

		// Log the request details
		log.Printf("Request: %s %s | Duration: %v | Status: %d",
			c.Request.Method,
			c.Request.URL.Path,
			time.Since(start),
			c.Writer.Status())
	}
}