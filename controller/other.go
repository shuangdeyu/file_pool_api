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
	head_pic := ""
	user_name := ""
	desc := ""
	info := service.GetByMap(service.CACHE_APP_TOKEN + token)
	user_id, ok := info["user_id"].(string)
	if info != nil && ok && comhelper.StringToInt(user_id) > 0 {
		is_login = 1
		// 获取用户信息
		ret := service.GetUserInfo(user_id)
		info, ok := ret["data"].(map[string]interface{})
		if ok && len(info) > 0 {
			user_name = info["name"].(string)
			head_pic = info["head_pic"].(string)
			desc = info["desc"].(string)
		}
	}
	return map[string]interface{}{
		"is_login":  is_login,
		"user_name": user_name,
		"head_pic":  head_pic,
		"desc":      desc,
	}
}
