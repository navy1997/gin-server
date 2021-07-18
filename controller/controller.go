package controller

import (
	"bubble/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexController(context *gin.Context) {
	context.HTML(http.StatusOK,"index.html",nil)
}

func CreateTodo(context *gin.Context) {
	// 前端页面填写代办事项 点击提交到后端接口获取数据
	// 从请求获取数据
	var todo models.Todo
	context.BindJSON(&todo)
	// 存入数据库
	err := models.CreateTodo(&todo)
	if err!=nil{
		context.JSON(http.StatusOK,gin.H{
			"error":err.Error(),
		})
	}else{
		context.JSON(http.StatusOK,todo)
	}
}

func GetTodoList(context *gin.Context) {
	// 获取todo列表
	todoList,err := models.GetTodoList()
	if  err != nil{
		context.JSON(http.StatusOK,gin.H{
			"error":err.Error(),
		})
	}else{
		context.JSON(http.StatusOK,todoList)
	}
}

func GetTodo(context *gin.Context) {
	id,ok := context.Params.Get("id")
	if !ok {
		context.JSON(http.StatusOK,gin.H{
			"msg":"id未传",
		})
		return
	}
	// 根据id获取todo
	todo,err := models.GetTodoById(id)
	if err!=nil{
		context.JSON(http.StatusOK,gin.H{
			"mes":"id不存在",
		})
		return
	}
	context.JSON(http.StatusOK,gin.H{
		"msg":"查询成功",
		"data":todo,
	})

}

func UpdateTodo(context *gin.Context) {
	// 从请求获取id
	id,_ := context.Params.Get("id")
	todo,err := models.GetTodoById(id)
	if err!=nil{
		context.JSON(http.StatusOK,gin.H{
			"error":err.Error(),
		})
	}
	// 根据id更新todo状态
	context.BindJSON(&todo)
	if err = models.UpdateTodo(todo);err!=nil{
		context.JSON(http.StatusOK,gin.H{
			"error":err.Error(),
		})
	}else{
		context.JSON(http.StatusOK,todo)
	}

}

func DeleteTodo(context *gin.Context) {
	// 获取id
	id,ok := context.Params.Get("id")
	if !ok{
		context.JSON(http.StatusOK,gin.H{
			"error":"无效的id",
		})
	}else{
		// 删除todo
		if err := models.DeleteTodoById(id);err != nil{
			context.JSON(http.StatusOK,gin.H{
				"error":err.Error(),
			})
		}else{
			context.JSON(http.StatusOK,gin.H{
				"id":"deleted",
			})
		}
	}
}