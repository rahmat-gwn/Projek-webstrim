package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware adalah middleware sederhana untuk memeriksa autentikasi
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		// Anda dapat menambahkan logika validasi token di sini
		c.Next()
	}
}
