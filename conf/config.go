package conf

import (
	"github.com/cihub/seelog"
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"log"
	"time"
)

// 配置项
type Configuration struct {
	ProxyServer   string        `yaml:"proxy_server"`   // rpc地址
	RedisHost     string        `yaml:"redis_host"`     // redis地址
	RedisPassword string        `yaml:"redis_password"` // redis密码
	RedisTimeout  time.Duration `yaml:"redis_timeout"`  // redis超时时间
	RedisDatabase int           `yaml:"redis_database"` // redis库
	LocalLog      string        `yaml:"local_log"`      // 是否开启本地日志
}

var AppConfig *Configuration

// 日志配置文件
var AppLogger seelog.LoggerInterface

// 加载配置文件
func LoadConfig(path string) error {
	// 读取文件
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	var config Configuration
	// 解析文件数据，并存入结构体
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return err
	}
	// 将解析后的数据赋值给全局变量
	AppConfig = &config
	return err
}

// 获取配置
func GetConfig() *Configuration {
	return AppConfig
}

// 初始化日志配置
func InitLog(applogPath *string) {
	logger1, err1 := seelog.LoggerFromConfigAsFile(*applogPath)
	if err1 != nil {
		log.Println("日志配置错误")
		return
	}
	AppLogger = logger1
}
