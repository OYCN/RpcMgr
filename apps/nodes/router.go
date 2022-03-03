package nodes

import (
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.RouterGroup) {
	e.GET("/nodes", HandleNodesGetNodeList)
	e.POST("/node", HandleAddNode)
}
