package middleware

import (
	"github.com/gin-gonic/gin"
)

func HandleNeedLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.Redirect(302, "/")
		c.Next()
	}
}
