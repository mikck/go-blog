package routers

import (
	"blog/app/controller"
	"blog/app/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Option func(*gin.Engine)

var options []Option

// 注册app的路由配置
func Include(opts ...Option) {
	options = append(options, opts...)
}

// 初始化
func Init() *gin.Engine {
	r := gin.Default()

	r.POST("/login", controller.Login)
	r.POST("/register", controller.Register)
	r.POST("/refreshToken", controller.RefreshToken)

	authorize := r.Group("/", middleware.JWTAuth())
	{
		authorize.GET("user", func(c *gin.Context) {
			claims := c.MustGet("claims").(*middleware.CustomClaims)
			fmt.Println(claims.ID)
			c.String(http.StatusOK, claims.Name)
		})
	}
	return r
}
