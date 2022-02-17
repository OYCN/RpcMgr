package main

import (
	"webbk/db"
	"webbk/middleware"
	"webbk/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	r := gin.New()
	e := r.Group("/api")
	middleware.Load(e)
	routers.Load(e)
	r.Run("0.0.0.0:8000")
}
