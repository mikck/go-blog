package main

import (
	"blog/app"
	"blog/app/model"
	"blog/mysql"
	"blog/routers"
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	//加载配置
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db := mysql.DbObj()
	defer db.Close()
	db.AutoMigrate(&model.User{})
	// 加载多个APP的路由配置
	routers.Include(app.Routers)
	// 初始化路由
	r := routers.Init()
	if err := r.Run(); err != nil {
		fmt.Printf("startup service failed, err:%v\n\n", err)
	}
}
