package app

import (
	"blog/app/controller"
	"blog/app/middleware"
	"github.com/gin-gonic/gin"
)

func Routers(r *gin.Engine) {
	r.Use(middleware.Auth)
	r.GET("/api", controller.Index)
}
