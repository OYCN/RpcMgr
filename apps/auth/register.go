package auth

import (
	"fmt"
	"webbk/config"
	"webbk/db"

	"github.com/gin-gonic/gin"
)

func HandleRegister(c *gin.Context) {
	var form RegisterForm
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
	var count int64
	db.Db.Model(&db.User{}).Where(&db.User{Name: form.Username}).Count(&count)
	if count != 0 {
		c.JSON(400, config.Ret {
			Status: false,
			Data: "The user is already registered",
		})
		return
	}
	user := db.User {
		Name: form.Username,
		Passwd: form.Password,
	}
	result := db.Db.Create(&user)
	if result.Error != nil {
		c.JSON(400, config.Ret {
			Status: false,
			Data: result.Error.Error(),
		})
		return
	}
	if result.RowsAffected != 1 {
		panic("Affected rows is not 1, maybe an error")
	}
	ret := config.Ret {
		Status: true,
		Data: "success",
	}
	c.JSON(200, ret)
}