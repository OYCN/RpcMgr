package auth

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"webbk/config"
)

func HandleLogin(c *gin.Context) {
	var form LoginForm
	err := c.ShouldBindJSON(&form);
	if err != nil {
		c.JSON(400, config.Ret {
			Status: false,
			Data: err.Error(),
		})
		return
	}
	fmt.Println(form.Username)
	fmt.Println(form.Password)
	fmt.Println(form.Remember)
	session := sessions.Default(c)
	session.Set("isLogined", true)
	session.Save()
	ret := config.Ret {
		Status: true,
		Data: "success",
	}
	c.JSON(200, ret)
}
