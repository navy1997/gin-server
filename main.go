package main

import (
	"bubble/dao"
	"bubble/models"
	"bubble/routers"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


func main() {
	// 创建数据库
	// sql : create database bubble;
	// 连接数据库
	err := dao.InitMySQL()
	if err != nil{
		panic(err)
	}
	defer dao.DB.Close()	//数据库关闭
	// 绑定模型
	dao.DB.AutoMigrate(&models.Todo{})

	r := routers.SetupRouter()
	r.Run()
}
