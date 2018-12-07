package controller

import (
	"file_pool_api/conf"
	"file_pool_api/helpers"
	"github.com/gin-gonic/gin"
	"helper_go/comhelper"
	"runtime"
	"strings"
)

// app参数
type AppParam struct {
	Ip            string                 // 客户端ip
	Action        string                 // 接口号
	RequestConfig map[string]interface{} // 访问的接口版本信息
	RequestParam  map[string]string      // 请求参数
}

// 接口返回
type Out struct {
	E    int         `json:"e"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

/**
 * 主函数
 */
func AppIndex(c *gin.Context) {
	// app 参数初始化
	var AppInitParam *AppParam
	var OutData *Out
	AppInitParam = new(AppParam)
	AppInitParam.RequestParam = map[string]string{}
	AppInitParam.Ip = c.ClientIP()
	AppInitParam.Action = strings.TrimSpace(c.PostForm("action"))
	OutData = new(Out)
	OutData.E = 0    // 默认错误码
	OutData.Msg = "" // 默认错误信息

	// 判断接口是否存在
	req_config := conf.GetAppAction(AppInitParam.Action)
	if req_config != nil {
		AppInitParam.RequestConfig = req_config
	} else {
		OutPut(conf.INVALID_PARAM_ERROR, c, OutData, AppInitParam)
		return
	}

}

/**
 * 接口返回
 * @param e int 状态码
 */
func OutPut(e int, c *gin.Context, OutData *Out, AppInitParam *AppParam) {
	OutData.E = e
	if OutData.Msg == "" {
		if conf.GetAppConst(e) != "" {
			OutData.Msg = conf.GetAppConst(e)
		} else {
			OutData.Msg = "系统繁忙，请稍后再试"
		}
	}
	data, ok := OutData.Data.(map[string]interface{})
	if OutData.Data == nil || !ok || len(data) <= 0 {
		helpers.WriteAppErrorLog("[rpc-data:]" + comhelper.JsonEncode(OutData.Data))
		if e == conf.SUCCESS {
			OutData.Data = map[string]string{"res": "success"}
		} else {
			OutData.Data = nil
		}
	}
	result := map[string]interface{}{
		"e":    e,
		"msg":  OutData.Msg,
		"data": OutData.Data,
	}
	helpers.WriteAppErrorLog(AppInitParam.Action + " " + comhelper.JsonEncode(result))
	if e != conf.SUCCESS && OutData.Data == nil {
		c.JSON(200, gin.H{
			"e":   e,
			"msg": OutData.Msg,
		})
	} else {
		c.JSON(200, gin.H{
			"e":    e,
			"msg":  OutData.Msg,
			"data": OutData.Data,
		})
	}
	runtime.Gosched() // 出让时间片
}
