package auth

import (
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.RouterGroup) {
	e.GET("/token", HandleLoginGetToken)
	e.POST("/login", HandleLogin)
}
