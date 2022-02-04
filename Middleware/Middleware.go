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
