package routers

import (
	"Bubble/controller"
	"Bubble/tools"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	if tools.Cfg.Release {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	// 告诉gin框架模板文件引用的静态文件去哪里找
	r.Static("/static", "static")
	// 告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")
	//注册启动Controller
	new(controller.IndexController).IndexController(r)
	new(controller.TodoController).TodoController(r)
	return r
}
