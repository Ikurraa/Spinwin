package Middleware

import (
	"awesomespinner/Controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := Controller.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Tidak valid")
			c.Abort()
			return
		}
		c.Next()
	}
}

func JwtAuthMiddlewareTicket() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := Controller.TokenValidTicket(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Tidak valid")
			c.Abort()
			return
		}
		c.Next()
	}
}

func CORS() gin.HandlerFunc {
	// TO allow CORS
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
