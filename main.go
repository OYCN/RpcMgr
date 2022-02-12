package main

import (
	"webbk/middleware"
	"webbk/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	e := r.Group("/api")
	middleware.Load(e)
	routers.Load(e)
	r.Run("0.0.0.0:8000")
}
