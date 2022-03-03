package routers

import (
	"github.com/gin-gonic/gin"

	"webbk/apps/auth"
	"webbk/apps/nodes"
	"webbk/apps/rpc"

	"webbk/middleware"
)

func Load(e *gin.RouterGroup) {
	nodes.Routers(e.Group("/nodes", middleware.HandleNeedLogin()))
	auth.Routers(e.Group("/auth"))
	rpc.Routers(e.Group("/rpc", middleware.HandleNeedLogin()))
}