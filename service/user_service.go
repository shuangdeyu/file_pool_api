package service

import (
	"helper_go/comhelper"
	"helper_go/pwdhelper"
)

/**
 * 用户登录
 * @param name string 用户名
 * @param password string 密码
 */
func Login(name, password string) map[string]interface{} {
	en_password := pwdhelper.PasswordHash(password, true)
	data := _call(FilePoolService, "Login", map[string]interface{}{
		"Name":     name,
		"Password": en_password,
	})
	return data
}

/**
 * 用户注册
 * @param name string 用户名
 * @param password string 密码
 * @param email string 邮箱
 */
func Register(name, password, email string) map[string]interface{} {
	en_password := pwdhelper.PasswordHash(password, true)
	data := _call(FilePoolService, "Register", map[string]interface{}{
		"Name":     name,
		"Password": en_password,
		"Email":    email,
	})
	return data
}

/**
 * 获取用户基本信息
 * @param user_id int 用户id
 */
func Get_user_info(user_id string) map[string]interface{} {
	data := _call(FilePoolService, "GetUserInfo", map[string]interface{}{
		"UserId": comhelper.StringToInt(user_id),
	})
	return data
}

/**
 * 根据用户名获取用户信息
 * @param name string 用户名
 */
func Get_user_info_by_name(name string) map[string]interface{} {
	data := _call(FilePoolService, "GetUserInfoByName", map[string]interface{}{
		"Name": name,
	})
	return data
}
