package helpers

import (
	"file_pool_api/conf"
	"github.com/cihub/seelog"
)

// app日志
func WriteAppInfoLog(msg string) {
	if conf.GetConfig().LocalLog == "1" {
		seelog.UseLogger(conf.AppLogger)
		seelog.Info(msg)
		seelog.Flush()
	}
}
func WriteAppErrorLog(msg string) {
	if conf.GetConfig().LocalLog == "1" {
		seelog.UseLogger(conf.AppLogger)
		seelog.Error(msg)
		seelog.Flush()
	}
}
func WriteAppDebugLog(msg string) {
	if conf.GetConfig().LocalLog == "1" {
		seelog.UseLogger(conf.AppLogger)
		//seelog.Debug(msg)
		seelog.Critical(msg)
		seelog.Flush()
	}
}
