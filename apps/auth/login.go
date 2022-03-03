package auth

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"webbk/config"
	"webbk/db"
)

func HandleLogin(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("isLogined") == true {
		c.JSON(400, config.Ret {
			Status: false,
			Data: "Already logined",
		})
		return
	}
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
	fmt.Println(*form.Remember)
	sql := db.Db.Model(&db.User{}).Where(&db.User{Name: form.Username, Passwd: form.Password})
	if sql.Error != nil {
		c.JSON(400, config.Ret {
			Status: false,
			Data: sql.Error.Error(),
		})
		return
	}
	var count int64
	sql.Count(&count)
	if count == 0 {
		c.JSON(400, config.Ret {
			Status: false,
			Data: "User not found or passwd error",
		})
		return
	}
	var user db.User
	result := sql.First(&user)
	if result.Error != nil {
		c.JSON(400, config.Ret {
			Status: false,
			Data: result.Error.Error(),
		})
		return
	}
	session.Set("isLogined", true)
	session.Set("uid", user.ID)
	session.Set("name", user.Name)
	session.Save()
	ret := config.Ret {
		Status: true,
		Data: "success",
	}
	c.JSON(200, ret)
}
