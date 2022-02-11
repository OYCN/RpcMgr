package main

import (
	"webbk/middleware"
	"webbk/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	middleware.Load(r)
	routers.Load(r)
	r.Run("0.0.0.0:8000")
}
