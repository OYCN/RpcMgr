package auth

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"webbk/config"
)

func HandleLogout(c *gin.Context) {
	session := sessions.Default(c)
	isLogined := session.Get("isLogined")
	if isLogined == true {
		session.Set("isLogined", false)
		session.Save()
		ret := config.Ret {
			Status: true,
			Data: "success",
		}
		c.JSON(200, ret)
	} else {
		ret := config.Ret {
			Status: false,
			Data: "Not login",
		}
		c.JSON(400, ret)
	}
}
