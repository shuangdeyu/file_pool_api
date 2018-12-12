package service

import (
	"file_pool_api/conf"
	"file_pool_api/helpers"
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
	var data = make(map[string]interface{})
	var param_arr = make(map[string]string)
	var param_write_log = make(map[string]string)
	var param_write_sign = make(map[string]string)
	var request_param_key []string // request_param 键值的slice

	if param != "" {
		// 开始解密
		public_key := helpers.ReadFile(conf.PUBLIC_KEY_PATH)
		private_key := helpers.ReadFile(conf.PRIVATE_KEY_PATH)
		param_decrypt, err := pwdhelper.RsaDecrypt(param, public_key, private_key)
		if len(param_decrypt) > 0 && err == nil {
			for _, v := range strings.Split(param_decrypt, "&") {
				arr_v := strings.SplitN(v, "=", 2)
				if len(arr_v) > 1 {
					// 签名校验需要使用解码前的参数
					param_write_sign[arr_v[0]] = arr_v[1]
					// 业务逻辑使用解码后的参数
					val, _ := url.PathUnescape(arr_v[1])
					param_arr[arr_v[0]] = val
					param_write_log[arr_v[0]] = val
					request_param_key = append(request_param_key, arr_v[0])
				}
			}
		} else {
			return SERVICE_ERROR_PARAM_INVALID, nil
		}
	}
	// 获取用户缓存
	ret := GetByMap(CACHE_APP_TOKEN + token)
	if ret == nil {
		param_arr["user_id"] = ""
	} else {
		param_arr["user_id"] = ret["user_id"].(string)
	}
	param_arr["token"] = token

	data["param_arr"] = param_arr                 // 使用参数
	data["param_write_log"] = param_write_log     // 用于日志处理参数
	data["param_write_sign"] = param_write_sign   // 用户签名校验参数
	data["param_decrypt_key"] = request_param_key // 参数key
	return SERVICE_SUCCESS, data
}
