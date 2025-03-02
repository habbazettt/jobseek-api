package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role.(string) != "admin" {
			c.JSON(http.StatusForbidden, gin.H{
				"status":  "error",
				"message": "Forbidden: Admin access only",
				"data":    nil,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
