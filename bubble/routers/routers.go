package routers

import (
	"hellp/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// 搭建 gin框架
	r := gin.Default()
	// 解析 静态 页面
	r.Static("/static", "static")
	// 告诉 gin 去哪找 html文件  解析文件
	r.LoadHTMLGlob("templates/*")

	r.GET("/", controller.IndexHandler)

	// 功能实现
	v1Group := r.Group("v1")
	{
		// 增加信息
		v1Group.POST("/todo", controller.CreateATodo)
		// 查询所有 信息
		v1Group.GET("/todo", controller.GetTodoList)
		// 查询某一条信息
		v1Group.GET("/todo/:id", controller.GetATodo)
		// 更新某一条信息
		v1Group.PUT("/todo/:id", controller.UpdataTodo)
		// 删除某一条信息
		v1Group.DELETE("/todo/:id", controller.DeleteTodo)
	}
	return r
}
