package main

import (
	"InsTest/bootstrap"
	"InsTest/global"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	// 初始化配置
	bootstrap.InitializeConfig()

	// 初始化日志
	global.App.Log = bootstrap.InitializeLog()
	global.App.Log.Info("log init success!")
	// 启动服务器
	bootstrap.RunServer()
	fmt.Println("服务启动成功")
}
