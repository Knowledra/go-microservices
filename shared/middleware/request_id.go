package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID := c.GetHeader("X-Request-Id")
		if reqID == "" {
			reqID = uuid.New().String()
		}
		
		// Set in Gin context so it can be retrieved by logger
		c.Set("X-Request-Id", reqID)
		
		// Set in request header so it propagates if we re-use standard http logic
		c.Request.Header.Set("X-Request-Id", reqID)
		
		// Set in response header for the client
		c.Header("X-Request-Id", reqID)
		
		c.Next()
	}
}
