package auth

import (
	"webbk/config"

	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

func HandleLoginGetCsrfToken(c *gin.Context) {
	c.JSON(200, config.Ret{
		Status: true,
		Data: csrf.GetToken(c),
	})
}
