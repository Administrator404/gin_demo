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
	// 根据page和size拿到发票信息
	r.GET("/test", service.GetInvs)
	// 根据发票号码拿到发票信息、收货人信息、商品信息
	r.GET("/test1/:invNumber", service.GetInvsALL)
	// 插入记录,
	r.PUT("/put1", service.PutInvs)
	// 运行服务端
	r.Run("127.0.0.1:8080")
}
