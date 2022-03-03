package rpc

import (
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.RouterGroup) {
	e.GET("/get", HandleDefaultRpc)
	e.POST("/ctx", HandleCreatCtx)
}
