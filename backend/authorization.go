package main

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		claims := jwt.ExtractClaims(c)
		_ = claims
		//		user, _ := c.Get()
		c.Next()
	}
}
