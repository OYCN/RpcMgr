package middleware

import (
	"github.com/gin-gonic/gin"
)

func Load(e *gin.RouterGroup) {
	e.Use(HandleCors())
	e.Use(HandleSession())
	e.Use(HandleCsrf())

	// e.Use(HandleNeedLogin())
}
