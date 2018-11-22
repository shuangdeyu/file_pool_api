package main

import (
	"file_pool_api/conf"
	"file_pool_api/controller"
	"flag"
	"github.com/gin-gonic/gin"
	"log"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	configPath := flag.String("c", "conf/config.yaml", "config file path") // 系统配置文件

	// 初始化日志配置
	applogPath := flag.String("la", "conf/log-app.xml", "app log config file path")

	flag.Parse()
	conf.InitLog(applogPath)
	// 加载系统配置文件
	if err := conf.LoadConfig(*configPath); err != nil {
		log.Println("加载系统配置出错！")
	}

	gin.SetMode(gin.ReleaseMode) // 设置运行环境
	r := gin.Default()           // 启动gin

	r.POST("/user", controller.GetUserInfo)

	r.Run(":6001")
}
