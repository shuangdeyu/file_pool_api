package controller

import (
	"file_pool_api/conf"
	"file_pool_api/helpers"
	"file_pool_api/service"
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

	/** 请求处理开始 **/
	var va [2]bool
	valid := req_config["valid"]
	va = valid.([2]bool)
	need_check := va[0] // 是否需要验签标志
	need_login := va[1] // 是否需要登录标志

	// 参数解密
	token := strings.TrimSpace(c.PostForm("token"))     // 客户端唯一码
	param := strings.TrimSpace(c.PostForm("param"))     // 加密的参数
	encrypt := strings.TrimSpace(c.PostForm("encrypt")) // 请求合法验证码
	e, ret := service.DecryptParam(param, token, AppInitParam.Ip)
	if e != service.SERVICE_SUCCESS {
		OutPut(conf.DECRYPT_ERROR, c, OutData, AppInitParam)
		return
	}
	AppInitParam.RequestParam = ret["param_arr"].(map[string]string)
	// 记录请求日志
	delete(ret["param_write_log"].(map[string]string), "password")
	helpers.WriteAppInfoLog(AppInitParam.Action +
		", ip:" + AppInitParam.Ip +
		", request:" + comhelper.JsonEncode(ret["param_write_log"]) +
		", and user_id:" + AppInitParam.RequestParam["user_id"] +
		" token: " + token)

	/*~!@#$% 本地测试，关闭验证 ~!@#$%*/
	if local := strings.TrimSpace(c.PostForm("local")); local == "1" {
		need_check = false // 测试用
		need_login = false // 测试用
	}
	// 检测参数是否正确
	if need_check {
		sign := ""
		if AppInitParam.RequestParam["sign"] != "" {
			sign = AppInitParam.RequestParam["sign"]
		}
		tp := AppInitParam.RequestParam["timestamp"]
		e = _valid_param(ret["param_write_sign"].(map[string]string), ret["param_decrypt_key"].([]string), encrypt, sign, tp, AppInitParam)
		if e != conf.SUCCESS {
			OutPut(e, c, OutData, AppInitParam)
			return
		}
	}
	// 检测是否需要登录
	if need_login {
		if ok := _valid_login(AppInitParam); !ok {
			OutPut(conf.USER_NOT_LOGIN_ERROR, c, OutData, AppInitParam)
			return
		}
	}

	// 调用处理函数
	/*~!@#$% 本地测试，重置param参数 ~!@#$%*/
	if local := strings.TrimSpace(c.PostForm("local")); local == "1" {
		AppInitParam.RequestParam["user_id"] = c.PostForm("user_id")
		AppInitParam.RequestParam["token"] = c.PostForm("token")
		AppInitParam.RequestParam["encrypt"] = c.PostForm("encrypt")
		AppInitParam.RequestParam["action"] = c.PostForm("action")
		AppInitParam.RequestParam["search"] = c.PostForm("search")
		AppInitParam.RequestParam["offset"] = c.PostForm("offset")
		AppInitParam.RequestParam["limit"] = c.PostForm("limit")
		AppInitParam.RequestParam["pool_id"] = c.PostForm("pool_id")
		AppInitParam.RequestParam["type"] = c.PostForm("type")
		AppInitParam.RequestParam["pool_user_id"] = c.PostForm("pool_user_id")
		AppInitParam.RequestParam["name"] = c.PostForm("name")
		AppInitParam.RequestParam["desc"] = c.PostForm("desc")
		AppInitParam.RequestParam["icon"] = c.PostForm("icon")
		AppInitParam.RequestParam["permit"] = c.PostForm("permit")
	}
	data, code := hook(AppInitParam)

	if data != nil {
		if data["e"] != nil {
			code = data["e"].(int)
		}
		if data["msg"] != nil {
			OutData.Msg = data["msg"].(string)
			OutData.Data = nil
		} else {
			if data["data"] != nil && data["e"] != nil {
				code = data["e"].(int)
				OutData.Data = data["data"]
			} else {
				if data["e"] == nil {
					delete(data, "e")
					OutData.Data = data
				}
			}
		}
	}
	// 请求日志处理入库

	OutPut(code, c, OutData, AppInitParam)
	return
}

/**
 * 调用接口处理函数
 */
var FuncsMap = map[string]func(param *AppParam) map[string]interface{}{
	"Login":            Login,            // 登录 100
	"Register":         Register,         // 注册 101
	"LoginOut":         LoginOut,         // 退出登录 102
	"UserInfo":         UserInfo,         // 用户信息 103
	"UserPoolList":     UserPoolList,     // 获取用户池列表 150
	"CreatePool":       CreatePool,       // 新建池 151
	"DeletePool":       DeletePool,       // 删除池 152
	"PoolInfo":         PoolInfo,         // 获取池信息 153
	"EditPoolPermit":   EditPoolPermit,   // 修改池权限 154
	"PoolMembers":      PoolMembers,      // 获取池成员列表 155
	"AddPoolMembers":   AddPoolMembers,   // 添加池成员列表 156
	"DeletePoolMember": DeletePoolMember, // 删除池成员 157
	"FileList":         FileList,         // 获取公共文档列表 200
	"CheckTokenLogin":  CheckTokenLogin,  // 检查客户端是否登录 400
}

func hook(AppInitParam *AppParam) (map[string]interface{}, int) {
	method, ok := AppInitParam.RequestConfig["method"].(string)
	if !ok || method == "" {
		helpers.WriteAppInfoLog("config:" + comhelper.JsonEncode(AppInitParam.RequestConfig) + " class or method not set")
		return nil, conf.SERVER_ERROR
	}
	data := FuncsMap[method](AppInitParam)
	return data, conf.SUCCESS
}

/**
 * 校验是否需要登录
 */
func _valid_login(AppInitParam *AppParam) bool {
	user_id := AppInitParam.RequestParam["user_id"]
	if user_id == "" || comhelper.StringToInt(user_id) <= 0 {
		return false
	}
	return true
}

/**
 * 校验参数
 * @param valid_arr map[string]string 解密的参数
 * @param encrypt string 请求合法码
 * @param sign string 签名
 * @param timestamp string 时间戳
 * @return e int 状态码
 */
func _valid_param(valid_arr map[string]string, valid_arr_key []string, encrypt, sign, timestamp string, AppInitParam *AppParam) int {
	// 校验请求头信息
	en := comhelper.Md5(comhelper.Md5("file_pool{z*2!H&akQ$cM|0_0com") + comhelper.Md5("webapp"+AppInitParam.Action) + timestamp)
	if en != encrypt {
		// 记录日志
		helpers.WriteAppInfoLog("encrypt valid failure;post encrypt: " + encrypt + ", " + en)
		return conf.ACCESS_DENY_ERROR
	}
	// 校验参数签名信息
	delete(valid_arr, "sign")
	if len(valid_arr) > 0 {
		sign_mobile := comhelper.Md5(comhelper.Array2UrlString(valid_arr, valid_arr_key, "sign"))
		if sign != sign_mobile {
			helpers.WriteAppInfoLog("valid param failure; sign:" + sign + ";sign_app:" + sign_mobile)
			return conf.INVALID_PARAM_ERROR
		}
	}
	return 0
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
	result := map[string]interface{}{
		"e":    e,
		"msg":  OutData.Msg,
		"data": OutData.Data,
	}
	helpers.WriteAppErrorLog(AppInitParam.Action + " " + comhelper.JsonEncode(result))
	//if e != conf.SUCCESS && OutData.Data == nil {
	if OutData.Data == nil {
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
