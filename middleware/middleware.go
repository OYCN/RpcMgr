package middleware

import (
	"github.com/gin-gonic/gin"
)

func Load(e *gin.Engine) {
	e.Use(HandleCors())
	e.Use(HandleSession())
	e.Use(HandleCsrf())

	e.Use(HandleNeedLogin())
}
