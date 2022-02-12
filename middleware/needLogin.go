package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"webbk/config"
)

type needLoginRet struct {
	Msg string `json:"msg"`
	Url string `json:"url"`
}

func HandleNeedLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		isLogined := session.Get("isLogined")
		if isLogined == nil {
			session.Set("isLogined", false)
			session.Save()
		}
		if session.Get("isLogined") == true {
			c.Next()
		} else {
			c.JSON(400, config.Ret{
				Status: false,
				Action: "redirect",
				Data: config.LoginRedirect,
			})
			c.Abort()
		}
	}
}
