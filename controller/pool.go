package controller

import (
	"file_pool_api/conf"
	"file_pool_api/service"
	"helper_go/comhelper"
	"strings"
)

/**
 * 获取用户池列表 150
 */
func UserPoolList(AppInitParam *AppParam) map[string]interface{} {
	param := AppInitParam.RequestParam
	search := comhelper.DefaultParam(param["search"], "")
	offset := comhelper.DefaultParam(param["offset"], "0")
	limit := comhelper.DefaultParam(param["limit"], "12")
	p_type := comhelper.DefaultParam(param["type"], "ALL") // all 所有,mine 我创建的,join 我参与的,delete 回收站
	p_type = strings.ToUpper(p_type)
	user_id := param["user_id"]

	ret := service.GetUserPoolList(user_id, search, offset, limit, p_type)
	data := ret["data"].(map[string]interface{})
	list := data["list"].([]interface{})
	count := data["count"]

	return map[string]interface{}{
		"count": count,
		"list":  list,
	}
}

/**
 * 新建池 151
 */
func CreatePool(AppInitParam *AppParam) map[string]interface{} {
	param := AppInitParam.RequestParam
	name := comhelper.DefaultParam(param["name"], "")
	desc := comhelper.DefaultParam(param["desc"], "")
	icon := comhelper.DefaultParam(param["icon"], "")
	permit := comhelper.DefaultParam(param["permit"], "10000")
	user_id := param["user_id"]

	if name == "" {
		return map[string]interface{}{
			"e":   conf.INVALID_PARAM_VALUE_ERROR,
			"msg": "名称不能为空",
		}
	}

	// 创建新池
	ret := service.CreateNewPool(name, desc, icon, permit, user_id)
	if ret["error"] != service.SERVICE_SUCCESS {
		return map[string]interface{}{
			"e":   conf.INVALID_PARAM_VALUE_ERROR,
			"msg": ret["data"],
		}
	}

	return nil
}

/**
 * 删除池 152
 */
func DeletePool(AppInitParam *AppParam) map[string]interface{} {
	param := AppInitParam.RequestParam
	pool_user_id := comhelper.DefaultParam(param["pool_user_id"], "0")

	// 先获取池信息
	ret := service.GetPoolInfoById(pool_user_id)
	if ret["error"] != service.SERVICE_SUCCESS {
		return map[string]interface{}{
			"e": conf.POOL_NOT_EXIST_ERROR,
		}
	}
	info := ret["data"].(map[string]interface{})
	// 判断权限并删除池/退出池
	if info["is_manager"].(string) != "Y" {
		// 管理员删除池
		ret = service.DeleteUserPoolById(pool_user_id, info["pool_id"].(string), true)
		if ret["error"] != service.SERVICE_SUCCESS {
			return map[string]interface{}{
				"e": conf.INVALID_PARAM_VALUE_ERROR,
			}
		}
	} else {
		// 非管理员退出池
		ret = service.DeleteUserPoolById(pool_user_id, info["pool_id"].(string), false)
		if ret["error"] != service.SERVICE_SUCCESS {
			return map[string]interface{}{
				"e": conf.INVALID_PARAM_VALUE_ERROR,
			}
		}
	}

	return nil
}

/**
 * 获取池信息 153
 */
func PoolInfo(AppInitParam *AppParam) map[string]interface{} {
	param := AppInitParam.RequestParam
	pool_user_id := comhelper.DefaultParam(param["pool_user_id"], "0")

	// 获取池信息
	ret := service.GetPoolInfoById(pool_user_id)
	if ret["error"] != service.SERVICE_SUCCESS {
		return map[string]interface{}{
			"e": conf.POOL_NOT_EXIST_ERROR,
		}
	}
	info := ret["data"].(map[string]interface{})

	return info
}

/**
 * 修改池权限 154
 */
func EditPoolPermit(AppInitParam *AppParam) map[string]interface{} {
	param := AppInitParam.RequestParam
	permit := comhelper.DefaultParam(param["permit"], "")
	pool_id := comhelper.DefaultParam(param["pool_id"], "0")
	user_id := param["user_id"]

	if permit == "" {
		return map[string]interface{}{
			"e":   conf.INVALID_PARAM_VALUE_ERROR,
			"msg": "权限值不能为空",
		}
	}

	// 权限判断
	ret := service.GetPoolInfo(pool_id)
	if ret["error"] != service.SERVICE_SUCCESS {
		return map[string]interface{}{
			"e": conf.POOL_NOT_EXIST_ERROR,
		}
	}
	info := ret["data"].(map[string]interface{})
	manager_id := info["manager_id"].(string)
	if user_id != manager_id {
		return map[string]interface{}{
			"e": conf.POOL_USER_CANT_EDIT_ERROR,
		}
	}
	// 修改权限值
	ret = service.EditPoolInfo(pool_id, permit)
	if ret["error"] != service.SERVICE_SUCCESS {
		return map[string]interface{}{
			"e": conf.INVALID_PARAM_ERROR,
		}
	}

	return nil
}

/**
 * 获取池成员列表 155
 */
func PoolMembers(AppInitParam *AppParam) map[string]interface{} {
	param := AppInitParam.RequestParam
	pool_id := comhelper.DefaultParam(param["pool_id"], "0")

	ret := service.GetPoolMembers(pool_id)
	list, ok := ret["data"].([]interface{})
	if len(list) < 1 || !ok {
		return map[string]interface{}{
			"count": 0,
			"list":  []interface{}{},
		}
	} else {
		return map[string]interface{}{
			"count": len(list),
			"list":  list,
		}
	}
}

/**
 * 添加池成员 156
 */
func AddPoolMembers(AppInitParam *AppParam) map[string]interface{} {

	return nil
}

/**
 * 删除池成员 157
 */
func DeletePoolMember(AppInitParam *AppParam) map[string]interface{} {
	param := AppInitParam.RequestParam
	pool_id := comhelper.DefaultParam(param["pool_id"], "0")
	pool_user_id := comhelper.DefaultParam(param["pool_user_id"], "0")
	user_id := param["user_id"]

	// 权限判断
	ret := service.GetPoolInfo(pool_id)
	if ret["error"] != service.SERVICE_SUCCESS {
		return map[string]interface{}{
			"e": conf.POOL_NOT_EXIST_ERROR,
		}
	}
	info := ret["data"].(map[string]interface{})
	manager_id := info["manager_id"].(string)
	if user_id != manager_id {
		return map[string]interface{}{
			"e":   conf.POOL_USER_CANT_EDIT_ERROR,
			"msg": "你没有删除成员的权限",
		}
	}
	// 删除成员
	ret = service.DeletePoolMembers(pool_user_id, pool_id)
	if ret["error"] != service.SERVICE_SUCCESS {
		return map[string]interface{}{
			"e": conf.INVALID_PARAM_ERROR,
		}
	}

	return nil
}
