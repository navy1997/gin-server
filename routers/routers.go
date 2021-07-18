package routers

import (
	"bubble/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter()*gin.Engine  {

	r := gin.Default()
	// 定义gin框架模板地址
	r.LoadHTMLGlob("templates/*")
	// 静态资源文件
	r.Static("/static","static")

	r.GET("/", controller.IndexController)

	// v1版本
	v1Group := r.Group("v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", controller.CreateTodo)
		// 查看
		v1Group.GET("/todo", controller.GetTodoList)
		// 查看某一个
		v1Group.GET("/todo/:id", controller.GetTodo)
		// 修改
		v1Group.PUT("/todo/:id", controller.UpdateTodo)
		// 删除
		v1Group.DELETE("/todo/:id", controller.DeleteTodo)
	}
	return r
}