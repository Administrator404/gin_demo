package main

import (
	"gin_demo/config"
	"gin_demo/service"
	"github.com/gin-gonic/gin"
)

func main() {
	// 数据库连接
	config.OpenDatabase()
	// 启动gin引擎
	r := gin.Default()
	// 路由
	r.GET("/test", service.GetInvs)
	// 运行服务端
	r.Run("127.0.0.1:8080")
}
