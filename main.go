package main

import (
	"Bubble/models"
	"Bubble/routers"
	"Bubble/tools"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		//fmt.Println("Usage: ./bubble config/config.ini   or  ./bubble config/config.json")
		fmt.Println("Usage: ./bubble config/config.ini")
		return
	}
	//加载配置文件
	fmt.Println(os.Args[1])
	if err := tools.Init(os.Args[1]); err != nil {
		fmt.Println("load config from file failed, err:", err)
		return
	}
	//使用 config.json
	//if err := tools.ParseJson(os.Args[1]); err != nil {
	//	fmt.Println("load config from file failed, err:", err)
	//	return
	//}

	//创建数据库
	//sql: CREATE DATABASE bubble；
	//连接数据库
	if err := tools.InitMySQL(tools.Cfg); err != nil {
		fmt.Println("init mysql failed, err:", err)
		return
	}

	//使用 config.json
	//if err := tools.InitMySQL(tools.CfgJson); err != nil {
	//	fmt.Println("init mysql failed, err:", err)
	//	return
	//}

	err := tools.DB.AutoMigrate(new(models.Todo))
	if err != nil {
		fmt.Println("DB AutoMigrate err ", err)
		return
	}
	r := routers.SetupRouter()
	if err := r.Run(fmt.Sprintf(":%v", tools.Cfg.Port)); err != nil {
		fmt.Println("server startup failed, err:", err)
	}

	//使用 config.json
	//if err := r.Run(fmt.Sprintf(":%v", tools.CfgJson.Port)); err != nil {
	//	fmt.Println("server startup failed, err:", err)
	//}
}

////创建数据库 Create Database bubble
//dsn := "root:root@tcp(127.0.0.1:3306)/bubble?charset=utf8mb4"
//DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//if err != nil {
//fmt.Println(err)
//}
//err = DB.AutoMigrate(new(models.Todo))
//if err != nil {
//return
//}
//r := gin.Default()
////告诉gin框架模板文件引用的静态文件去哪里找
//r.Static("/static", "static")
////告诉gin框架去哪里找模板文件
//r.LoadHTMLGlob("templates/*")
//
//r.GET("/", func(c *gin.Context) {
//	c.HTML(http.StatusOK, "index.html", nil)
//})
////v1
//v1Group := r.Group("v1")
//{
////待办事项
////添加
//v1Group.POST("/todo", func(c *gin.Context) {
////前端页面填写待办事项 点击提交 会发请求到这里
////1.从请求中数据拿出来
//var todo models.Todo
//err := c.BindJSON(&todo)
//if err != nil {
//return
//}
////2.存入数据库
//err = DB.Create(&todo).Error
//if err != nil {
//c.JSON(http.StatusOK, gin.H{"error": err.Error()})
//fmt.Println(err)
//return
//}
////3.返回响应
//c.JSON(http.StatusOK, todo)
//})
////查看所有的待办事项
//v1Group.GET("/todo", func(c *gin.Context) {
////查询todo这个表里的所有数据
//var todoList []models.Todo
//
//if err = DB.Find(&todoList).Error; err != nil {
//c.JSON(http.StatusOK, gin.H{"error": err.Error()})
//} else {
//c.JSON(http.StatusOK, todoList)
//}
//})
//// 查看某一个待办事项
//v1Group.GET("/todo/:id", func(c *gin.Context) {
//
//})
////修改某个待办事项
//v1Group.PUT("/todo/:id", func(c *gin.Context) {
//id, ok := c.Params.Get("id")
//if !ok {
//c.JSON(http.StatusOK, gin.H{"error": err.Error()})
//}
//var todo models.Todo
//if err = DB.Where("id=?", id).First(&todo).Error; err != nil {
//c.JSON(http.StatusOK, gin.H{"error": err.Error()})
//}
//
//if err := c.BindJSON(&todo); err != nil {
//fmt.Println(err)
//return
//}
//
//if err = DB.Save(&todo).Error; err != nil {
//c.JSON(http.StatusOK, gin.H{"error": err.Error()})
//} else {
//c.JSON(http.StatusOK, todo)
//}
//})
//// 删除某一个待办事项
//v1Group.DELETE("/todo/:id", func(c *gin.Context) {
//id, ok := c.Params.Get("id")
//if !ok {
//c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
//return
//}
//if err = DB.Where("id=?", id).Delete(models.Todo{}).Error; err != nil {
//c.JSON(http.StatusOK, gin.H{"error": err.Error()})
//} else {
//c.JSON(http.StatusOK, gin.H{id: "deleted"})
//}
//})
//}
//if err := r.Run(":80"); err != nil {
//fmt.Println(err)
//}
