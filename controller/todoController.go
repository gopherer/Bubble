package controller

import (
	"Bubble/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TodoController struct {
}

func (controller TodoController) TodoController(context *gin.Engine) {
	v1Group := context.Group("v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", controller.CreateTodo)
		// 查看所有的待办事项
		v1Group.GET("/todo", controller.GetTodoList)
		// 修改某一个待办事项
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		// 删除某一个待办事项
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}
}

func (controller *TodoController) CreateTodo(c *gin.Context) {
	// 前端页面填写待办事项 点击提交 会发请求到这里
	// 1. 从请求中把数据拿出来
	var todo *models.Todo
	err := c.BindJSON(&todo)
	if err != nil {
		return
	}
	// 2. 存入数据库
	err = todo.CreateATodo()
	if err != nil {
		return
	}
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
		//c.JSON(http.StatusOK, gin.H{
		//	"code": 2000,
		//	"msg": "success",
		//	"data": todo,
		//})
	}
}

func (controller *TodoController) GetTodoList(c *gin.Context) {
	// 查询todo这个表里的所有数据
	var todo *models.Todo
	todoList, err := todo.GetAllTodo()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

func (controller *TodoController) UpdateATodo(c *gin.Context) {
	var todo *models.Todo
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	newTodo, err := todo.GetATodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	err = c.BindJSON(newTodo)
	if err != nil {
		return
	}
	if err = newTodo.UpdateATodo(); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func (controller *TodoController) DeleteATodo(c *gin.Context) {
	var todo *models.Todo
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	if err := todo.DeleteATodo(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{id: "deleted"})
	}
}
