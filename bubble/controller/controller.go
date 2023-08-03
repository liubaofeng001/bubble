package controller

import (
	"hellp/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

func CreateATodo(ctx *gin.Context) {
	// 业务逻辑
	// 从前端获取 数据 json
	// 1 获取前端 json 格式数据
	var todo models.Todo
	ctx.BindJSON(&todo)
	// 2 将 这条数据 存入数据库
	err := models.CreateATodo(&todo)
	// 3 返回响应
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"err": "error",
		})
	} else {
		ctx.JSON(http.StatusOK, todo)
	}
}

// 获取 所有的 信息
func GetTodoList(ctx *gin.Context) {
	// 业务逻辑
	// 查询多条记录，Find函数返回的是一个数组。

	// 查询 多条数据 从数据库中
	// var todoList []models.Todo
	todoList, err := models.GetTodoList()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err.Error,
		})
	} else {
		ctx.JSON(http.StatusOK, todoList)
	}
}

func GetATodo(ctx *gin.Context) {
	// 业务逻辑
}

// 更新信息
func UpdataTodo(ctx *gin.Context) {
	// 业务逻辑
	// 1 根据 id 查询到 这条消息
	id := ctx.Param("id")
	// var todo models.Todo
	// 1.1 如果查不到这条数据 return
	// 1.2 查得到 继续
	todo, err := models.GetTodoById(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	// 2 从 前端获取 新的数据 bindJSON
	ctx.BindJSON(&todo)
	// 3 将这条数据 更新到数据库中
	if err := models.UpdataTodoById(todo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, todo)
	}
}

// 删除一条信息
func DeleteTodo(ctx *gin.Context) {
	// 业务逻辑
	// 删除 一条信息
	// db.Where("type = ?", 5).Delete(&Food{})
	id := ctx.Param("id")
	err := models.DeleteTodo(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err.Error,
		})
	}
}
