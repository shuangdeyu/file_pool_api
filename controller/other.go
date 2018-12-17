package controller

import (
	"file_pool_api/service"
	"helper_go/comhelper"
)

/**
 * 检查客户端是否登录 400
 */
func CheckTokenLogin(AppInitParam *AppParam) map[string]interface{} {
	param := AppInitParam.RequestParam
	token := param["token"]

	is_login := 0
	info := service.GetByMap(service.CACHE_APP_TOKEN + token)
	user_id, ok := info["user_id"].(string)
	if info != nil && ok && comhelper.StringToInt(user_id) > 0 {
		is_login = 1
	}
	return map[string]interface{}{
		"is_login": is_login,
	}
}
