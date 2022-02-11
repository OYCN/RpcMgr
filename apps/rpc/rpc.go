package rpc

import (
	"github.com/gin-gonic/gin"
)

func HandleDefaultRpc(c *gin.Context) {
	c.String(200, "rpc")
}
