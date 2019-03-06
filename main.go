package main

import (
	"file_pool_api/conf"
	"file_pool_api/controller"
	"flag"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"regexp"
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

	// 跨域中间件
	r.Use(CorsMiddleware())
	// 业务接口入口
	r.POST("/app/*action", controller.AppIndex)
	r.POST("/app", controller.AppIndex)
	// 其它接口
	r.POST("/get_token", controller.GetToken)         // 获取token
	r.POST("/get_encrypt", controller.GetEncrypt)     // 获取encrypt
	r.POST("/get_sign", controller.GetSign)           // 获取sign
	r.POST("/encrypt_param", controller.EncryptParam) // 参数加密

	r.Run(":6001")
}

// 跨域中间件
func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		// filterHost 做过滤器，防止不合法的域名访问
		var filterHost = [...]string{"http://localhost.*"}
		var isAccess = false
		for _, v := range filterHost {
			match, _ := regexp.MatchString(v, origin)
			if match {
				isAccess = true
			}
		}
		if isAccess {
			// 核心处理方式
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
			c.Header("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
			c.Set("content-type", "application/json")
		}
		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}

		c.Next()
	}
}
