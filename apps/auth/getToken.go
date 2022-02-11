package auth

import (
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

func HandleLoginGetToken(c *gin.Context) {
	c.String(200, csrf.GetToken(c))
}
