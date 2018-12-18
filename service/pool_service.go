package service

import "helper_go/comhelper"

/**
 * 获取用户文档池列表
 * @param user_id string 用户id
 * @param search string 搜索关键字
 * @param offset string 偏移量
 * @param limit string 请求数
 * @param p_type string 类型
 */
func GetUserPoolList(user_id, search, offset, limit, p_type string) map[string]interface{} {
	data := _call(FilePoolService, "GetUserPoolList", map[string]interface{}{
		"UserId": comhelper.StringToInt(user_id),
		"Search": search,
		"Offset": comhelper.StringToInt(offset),
		"Limit":  comhelper.StringToInt(limit),
		"Type":   p_type,
	})
	return data
}

/**
 * 根据用户文档池id获取文档池信息
 * @param pool_user_id string 用户池id
 */
func GetPoolInfoById(pool_user_id string) map[string]interface{} {
	data := _call(FilePoolService, "GetPoolInfoById", map[string]interface{}{
		"PoolUserId": comhelper.StringToInt(pool_user_id),
	})
	return data
}

/**
 * 删除池
 * @param pool_user_id string 用户池id
 */
func DeleteUserPoolById(pool_user_id string) map[string]interface{} {
	data := _call(FilePoolService, "DeleteUserPoolById", map[string]interface{}{
		"PoolUserId": comhelper.StringToInt(pool_user_id),
	})
	return data
}

/**
 * 创建池
 * @param name string 池名称
 * @param desc string 描述
 * @param icon string 图标
 * @param permit string 权限
 */
func CreateNewPool(name, desc, icon, permit, user_id string) map[string]interface{} {
	data := _call(FilePoolService, "CreateNewPool", map[string]interface{}{
		"Name":   name,
		"Desc":   desc,
		"Icon":   icon,
		"Permit": permit,
		"UserId": comhelper.StringToInt(user_id),
	})
	return data
}

/**
 * 修改文档池信息
 * @param pool_id string 池id
 * @param permit string 权限值
 */
func EditPoolInfo(pool_id, permit string) map[string]interface{} {

}
