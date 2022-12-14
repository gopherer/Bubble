# Bubble清单

一个基于gin+gorm开发的备忘录

一个基于gorm增删改查的web小程序

## 下载
`git clone https://github.com/gopherer/Bubble.git`

## 配置Mysql
1.在Mysql中执行一下代码创建项目所需的数据库：

`CREATE DATABASE bubble DEFAULT CHARSET=utf8mb4;`

2.在bubble/conf/config.ini文件中按如下提示配置数据库连接信息。
`

    port = 80
    release = false
    [mysql]
    user = 你的数据库用户名
    password = 你的数据库密码
    host = 你的数据库host地址
    port = 你的数据库端口
    db = bubble
`

新增`bubble/conf/config.json`使用json文件保持项目配置

## 编译

`go build -o bubble.exe .\main.go`

## 执行

`./bubble conf/config.ini`

亦或

`./bubble conf/config.json` 但需要启动注释中的代码

## 启动
在浏览器输入 localhost

浏览器默认访问80端口  localhost 是127.0.0.1的别名

![image ](https://github.com/gopherer/Bubble/blob/main/photo/bubble.png)

## 其他
[项目源地址](https://github.com/Q1mi/bubble)

[项目学习地址 最新Go Web开发教程基于gin框架和gorm的web开发实战 (七米出品)](https://www.bilibili.com/video/BV1gJ411p7xC?p=25&vd_source=41d370dbb2b396b1210612c7fce0fa21)
