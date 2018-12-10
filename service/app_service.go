package service

import (
	"helper_go/pwdhelper"
	"net/url"
	"strings"
)

const (
	// 正常接口返回
	SERVICE_SUCCESS = 0

	// 默认参数
	SERVICE_ERROR_PARAM_INVALID   = -1 // 无效参数
	SERVICE_PARAM_STRING_TYPE_ALL = ""

	// 时间格式模板
	TimeLayout = "2006-01-02 15:04:05"

	// 操作系统类型
	OS_TYPE_PC      = "normal"
	OS_TYPE_ANDROID = "android"
	OS_TYPE_IOS     = "ios"
	OS_TYPE_WP      = "wp"

	// 端类型
	ENDPOINT_WEB = "web"
	ENDPOINT_WAP = "wap"
	ENDPOINT_APP = "app"
	ENDPOINT_PC  = "pc"
)

/*
 * param 解密
 * @param param string 加密参数
 * @param token string token
 * @param ip string
 * @return int,map[string]interface{}
 */
func DecryptParam(param, token, ip string) (int, map[string]interface{}) {
	var param_arr = make(map[string]string)
	var data = make(map[string]interface{})
	var param_write_log = make(map[string]string)

	if param != "" {
		// 开始解密
		public_key := "-----BEGIN PUBLIC KEY-----\n" + "" + "\n-----END PUBLIC KEY-----"
		param_decrypt, err := pwdhelper.RsaDecrypt(param, public_key, "")
		if len(param_decrypt) > 0 && err == nil {
			for _, v := range strings.Split(param_decrypt, "&") {
				arr_v := strings.SplitN(v, "=", 2)
				if len(arr_v) > 1 {
					val, _ := url.PathUnescape(arr_v[1])
					param_arr[arr_v[0]] = val
					param_write_log[arr_v[0]] = val
				}
			}
		} else {
			return SERVICE_ERROR_PARAM_INVALID, nil
		}
	}
	//ret := Update_session(param_arr, token, agent, imei, ip, endpoint, agent_info) // 更新session
	//param_arr["u"] = ret["uid"]
	//param_arr["token"] = ret["token"]
	//param_arr["user_id"] = ret["user_id"]
	//param_arr["public_key"] = ret["public_key"]

	data["param_arr"] = param_arr             // 使用参数
	data["param_write_log"] = param_write_log // 用于日志处理参数
	return SERVICE_SUCCESS, data
}
