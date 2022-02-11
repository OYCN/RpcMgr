package routers

import (
	"github.com/gin-gonic/gin"

	"webbk/apps/auth"
	"webbk/apps/nodes"
	"webbk/apps/rpc"
)

func Load(e *gin.Engine) {
	nodes.Routers(e.Group("/nodes"))
	auth.Routers(e.Group("/auth"))
	rpc.Routers(e.Group("/rpc"))
}