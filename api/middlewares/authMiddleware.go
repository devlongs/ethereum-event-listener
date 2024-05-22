package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        apiKey := c.GetHeader("X-API-Key")
        if apiKey != "your-api-key" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid API key"})
            return
        }
        c.Next()
    }
}