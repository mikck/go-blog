package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Auth(c *gin.Context)  {
	t := time.Now()
	fmt.Println("我是自定义中间件第1种定义方式---请求之前")
	//在gin上下文中定义一个变量
	c.Set("example", "CustomRouterMiddle1")
	//请求之前
	c.Next()
	fmt.Println("我是自定义中间件第1种定义方式---请求之后")
	//请求之后
	//计算整个请求过程耗时
	t2 := time.Since(t)
	log.Println(t2)

}