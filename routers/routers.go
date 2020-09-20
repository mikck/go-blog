package routers

import (
	"blog/app/middleware"
	"fmt"
	"github.com/dgrijalva/jwt-go"
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
	r.GET("/jwt", func(c *gin.Context) {
		j := &middleware.JWT{
			SigningKey: []byte("test"),
		}
		claims := middleware.CustomClaims{
			ID:    1,
			Name:  "awh521",
			Email: "1044176017@qq.com",
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: 15000, //time.Now().Add(24 * time.Hour).Unix()
				Issuer:    "test",
			},
		}
		token, err := j.CreateToken(claims)
		if err != nil {
			c.String(http.StatusOK, err.Error())
			c.Abort()
		}
		c.String(http.StatusOK, token+"---------------<br>")
		res, err := j.ParseToken(token)
		fmt.Println(err)
		if err != nil {
			if err == middleware.TokenExpired {
				newToken, err := j.RefreshToken(token)
				if err != nil {
					c.String(http.StatusOK, err.Error())
				} else {
					c.String(http.StatusOK, newToken)
				}
			} else {
				c.String(http.StatusOK, err.Error())
			}
		} else {
			c.JSON(http.StatusOK, res)
		}
	})
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
