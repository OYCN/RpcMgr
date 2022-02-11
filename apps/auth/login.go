package auth

import (
	"github.com/gin-gonic/gin"
)

type loginForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Remember bool `json:"remember" binding:"required"`
	Token string `json:"token" binding:"required"`
}

type loginRet struct {
	Status bool `json:"status"`
	Msg string `json:"msg"`
}

func HandleLogin(c *gin.Context) {
	var form loginForm
	err := c.ShouldBindJSON(&form);
	if err != nil {
		c.JSON(503, gin.H{"error": err.Error()})
		return
	}
	ret := loginRet {
		Status: true,
		Msg: "success",
	}
	c.JSON(200, ret)
}
