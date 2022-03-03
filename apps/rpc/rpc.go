package rpc

import (
	"webbk/config"

	"github.com/gin-gonic/gin"
)

func HandleDefaultRpc(c *gin.Context) {
	c.JSON(200, config.Ret {
		Status: true,
		Data: "hello",
	})
}
