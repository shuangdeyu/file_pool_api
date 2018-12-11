package controller

import (
	"file_pool_api/service"
)

/**
 * 登录 100
 */
func Login(AppInitParam *AppParam) map[string]interface{} {
	return nil
}

/**
 * 注册 101
 */
func Register(AppInitParam *AppParam) map[string]interface{} {
	return nil
}

/**
 * 退出登录 102
 */
func LoginOut(AppInitParam *AppParam) map[string]interface{} {
	return nil
}

/**
 * 获取用户信息 103
 */
func UserInfo(AppInitParam *AppParam) map[string]interface{} {
	ret := service.Get_user_info(1)

	if ret["error"] != service.SERVICE_SUCCESS {

	}

	return ret["data"].(map[string]interface{})
}
