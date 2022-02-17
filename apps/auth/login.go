package auth

import (
	"errors"
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"webbk/config"
	"webbk/db"
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
	var user db.User
	result := db.Db.Where(&db.User{Name: form.Username, Passwd: form.Password}).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(400, config.Ret {
			Status: false,
			Data: "User not found or passwd error",
		})
		return
	}
	session := sessions.Default(c)
	session.Set("isLogined", true)
	session.Save()
	ret := config.Ret {
		Status: true,
		Data: "success",
	}
	c.JSON(200, ret)
}
