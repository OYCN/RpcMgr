package auth

import (
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.RouterGroup) {
	e.GET("/token", HandleLoginGetCsrfToken)
	e.POST("/login", HandleLogin)
	e.GET("/logout", HandleLogout)
	if gin.Mode() == gin.DebugMode {
		e.POST("/register", HandleRegister)
	}
}
