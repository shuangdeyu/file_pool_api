package controller

import (
	"file_pool_api/conf"
	"file_pool_api/service"
	"github.com/yvasiyarov/php_session_decoder/php_serialize"
	"helper_go/comhelper"
	"helper_go/verhelper"
)

/**
 * 登录 100
 */
func Login(AppInitParam *AppParam) map[string]interface{} {
	param := AppInitParam.RequestParam
	name := comhelper.DefaultParam(param["name"], "")
	password := comhelper.DefaultParam(param["password"], "")
	if name == "" || password == "" {
		return map[string]interface{}{
			"e":   conf.INVALID_PARAM_VALUE_ERROR,
			"msg": "用户名或密码不能为空",
		}
	}

	// 执行登录
	ret := service.Login(name, password)
	if ret["error"] != service.SERVICE_SUCCESS {
		return map[string]interface{}{
			"e":   conf.INVALID_PARAM_VALUE_ERROR,
			"msg": ret["data"],
		}
	}
	info := ret["data"].(map[string]interface{})
	// 用户缓存
	data := php_serialize.PhpArray{
		"user_id":   info["id"],
		"token":     param["token"],
		"timestamp": param["timestamp"],
		"ip":        AppInitParam.Ip,
	}
	encoder := php_serialize.NewSerializer()
	val, _ := encoder.Encode(data)
	service.Save(service.CACHE_APP_TOKEN, param["token"], val)

	return nil
}

/**
 * 注册 101
 */
func Register(AppInitParam *AppParam) map[string]interface{} {
	param := AppInitParam.RequestParam
	name := comhelper.DefaultParam(param["name"], "")
	password := comhelper.DefaultParam(param["password"], "")
	email := comhelper.DefaultParam(param["email"], "")
	if name == "" || password == "" {
		return map[string]interface{}{
			"e":   conf.INVALID_PARAM_VALUE_ERROR,
			"msg": "用户名或密码不能为空",
		}
	}
	if email != "" {
		if !verhelper.IsEmail(email) {
			return map[string]interface{}{
				"e":   conf.INVALID_PARAM_VALUE_ERROR,
				"msg": "邮箱格式不正确",
			}
		}
	}

	// 验证用户名是否已被注册
	ret := service.Get_user_info_by_name(name)
	if ret["error"] == service.SERVICE_SUCCESS {
		return map[string]interface{}{
			"e": conf.USER_REG_ALREADY_ERROR,
		}
	}
	// 执行注册
	ret = service.Register(name, password, email)
	if ret["error"] != service.SERVICE_SUCCESS {
		return map[string]interface{}{
			"e": conf.SERVER_ERROR,
		}
	}
	// 注册成功登录
	ret = Login(AppInitParam)
	if ret != nil && ret["e"] != service.SERVICE_SUCCESS {
		return map[string]interface{}{
			"e":   conf.INVALID_PARAM_VALUE_ERROR,
			"msg": ret["data"],
		}
	}

	return nil
}

/**
 * 退出登录 102
 */
func LoginOut(AppInitParam *AppParam) map[string]interface{} {
	param := AppInitParam.RequestParam
	token := param["token"]
	service.Delete(service.CACHE_APP_TOKEN + token)
	return nil
}

/**
 * 获取用户信息 103
 */
func UserInfo(AppInitParam *AppParam) map[string]interface{} {
	param := AppInitParam.RequestParam
	user_id := param["user_id"]

	ret := service.Get_user_info(user_id)
	info, ok := ret["data"].(map[string]interface{})
	if ret["error"] != service.SERVICE_SUCCESS || !ok || !(len(info) > 0) {
		return map[string]interface{}{
			"e": conf.USER_NOT_EXIST_ERROR,
		}
	}
	return info
}
