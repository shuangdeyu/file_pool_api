package service

/**
 * 获取用户基本信息
 * @param user_id int 用户id
 */
func Get_user_info(user_id int) map[string]interface{} {
	data := _call(FilePoolService, "GetUserInfo", map[string]interface{}{
		"UserId": user_id,
	})
	return data
}
