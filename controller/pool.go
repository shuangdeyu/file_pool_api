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
	// 判断权限并删除池
	if info["is_manager"].(string) != "Y" {
		return map[string]interface{}{
			"e": conf.POOL_USER_CANT_DELETE_ERROR,
		}
	}
	ret = service.DeleteUserPoolById(pool_user_id)
	if ret["error"] != service.SERVICE_SUCCESS {
		return map[string]interface{}{
			"e": conf.INVALID_PARAM_VALUE_ERROR,
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

	if permit == "" {
		return map[string]interface{}{
			"e":   conf.INVALID_PARAM_VALUE_ERROR,
			"msg": "权限值不能为空",
		}
	}

	// 修改权限值

	return nil
}
