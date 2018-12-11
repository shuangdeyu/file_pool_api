package conf

// 接口号数组
func ReadRequest() map[string]interface{} {
	var app map[string]interface{}
	app = make(map[string]interface{})

	/**
	 * 100~149 用户相关
	 * 150~199 池相关
	 * 200~249 文档相关
	 */
	app["app_request"] = map[string]interface{}{
		"100": map[string]interface{}{
			"method": "Login",              // 登录
			"valid":  [2]bool{true, false}, // 是否验签 是否登录
		},
		"101": map[string]interface{}{
			"method": "Register", // 注册
			"valid":  [2]bool{true, false},
		},
		"102": map[string]interface{}{
			"method": "LoginOut", // 退出登录
			"valid":  [2]bool{true, true},
		},
		"103": map[string]interface{}{
			"method": "UserInfo", // 用户信息
			"valid":  [2]bool{true, true},
		},
	}
	return app
}

// 根据接口号获取接口信息
func GetAppAction(action string) map[string]interface{} {
	arr := ReadRequest()
	method, ok := arr["app_request"].(map[string]interface{})[action].(map[string]interface{})
	if ok {
		return method
	} else {
		return nil
	}
}
