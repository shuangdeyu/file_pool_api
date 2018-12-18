package conf

const (
	// 系统级别错误
	SUCCESS                          = 0      // 成功
	INTERNAl_ERROR                   = 100000 // 系统繁忙，请稍后再试
	INVALID_PARAM_ERROR              = 100001 // 出错啦，请稍后再试
	HOST_IS_SPAM_ERROR               = 100002 // 访问被限制
	SERVER_ERROR                     = 100003 // 系统错误，请稍后再试
	SYSTEM_CLOSE_ERROR               = 100004 // 系统关闭
	USER_NOT_LOGIN_ERROR             = 100005 // 用户没有登录
	ACCESS_DENY_ERROR                = 100006 // 账号没有访问权限
	REPEAT_REQUEST_ERROR             = 100007 // 重复请求数据
	DECRYPT_ERROR                    = 100008 // 解密失败
	INVALID_PARAM_VALUE_ERROR        = 100009 // 参数值有误
	CURRENT_VERSION_NO_SUPPORT_ERROR = 100010 // 当前版本不再维护

	// User相关
	USER_AUTH_FAILED_ERROR              = 200100 // 账号或者密码错误
	USER_ACCOUNT_FROZEN_ERROR           = 200101 // 用户登录失败次数过多，帐户将被冻结24小时
	INVALID_MOBILE_ERROR                = 200102 // 手机号码不合法
	CAPTCHA_EXPIRE_ERROR                = 200103 // 验证码过期了
	INVALID_CAPTCHA_ERROR               = 200104 // 验证码错误
	MOBILE_REG_ALREADY_ERROR            = 200105 // 手机号码已经注册
	PASSWORD_RESET_FAILURE_ERROR        = 200106 // 修改密码失败
	USER_NOT_EXIST_ERROR                = 200107 // 用户不存在
	SEND_SMS_TOO_FAST_ERROR             = 200108 // 发送短信太频繁
	SEND_SMS_MORE_THAN_MAX_TIMES_ERROR  = 200109 // 短信发送频率超过限制，系统将冻结你的账号
	USER_ALREADY_REALNAME_AUTH          = 200110 // 用户已经实名认证
	USER_IDCARD_ALREADY_REGISTER        = 200111 // 该身份证号已经被注册
	USER_PASSWORD_ERROR                 = 200112 // 密码错误
	USER_PASSWORD_OLD_EQUAL_NEW_ERROR   = 200113 // 新旧密码一致
	USER_PASSWORD_NOT_EQUAL_ERROR       = 200313 // 两次输入密码不一致
	USER_INVITE_CODE_INVALID            = 200114 // 拜年红包无效
	IMAGE_CAPTCHA_INVALID_ERROR         = 200115 // 图像验证码错误
	USER_PASSWORD_INVALID_ERROR         = 200116 // 输入的密码不合法
	USER_ALREADY_LOGGED_ERROR           = 200117 // 您已经处于登录状态
	USER_RECOMMEND_USER_NOT_EXIST_ERROR = 200118 // 推荐用户不存在
	USER_ACCOUNT_NOT_EXIST_ERROR        = 200119 // 账户不存在
	USER_PASSWORD_CHECK_FAILED_ERROR    = 200120 // 校验密码错误
	USER_PAY_PASSWORD_NOT_EXIST_ERROR   = 200121 // 支付密码未设置
	USER_REG_ALREADY_ERROR              = 200122 // 用户名已被注册

	// 池相关
	POOL_NOT_EXIST_ERROR        = 200200 // 池不存在
	POOL_USER_CANT_DELETE_ERROR = 200201 // 您没有删除该池的权限
)

var ERR_MSG_MAP map[int]string

func GetAppConst(e int) string {
	if len(ERR_MSG_MAP) == 0 {
		ERR_MSG_MAP = make(map[int]string)
		// 系统级别错误
		ERR_MSG_MAP[SUCCESS] = "成功"
		ERR_MSG_MAP[INTERNAl_ERROR] = "系统繁忙，请稍后再试"
		ERR_MSG_MAP[INVALID_PARAM_ERROR] = "出错啦，请稍后再试"
		ERR_MSG_MAP[HOST_IS_SPAM_ERROR] = "账号已被锁定"
		ERR_MSG_MAP[SERVER_ERROR] = "系统错误，请稍后再试"
		ERR_MSG_MAP[SYSTEM_CLOSE_ERROR] = "系统升级中..."
		ERR_MSG_MAP[USER_NOT_LOGIN_ERROR] = "账号还没有登录"
		ERR_MSG_MAP[ACCESS_DENY_ERROR] = "账号没有访问权限"
		ERR_MSG_MAP[REPEAT_REQUEST_ERROR] = "操作太频繁"
		//ERR_MSG_MAP[DECRYPT_ERROR] = "解密失败"
		//ERR_MSG_MAP[INVALID_PARAM_VALUE_ERROR] = "参数值有误"
		ERR_MSG_MAP[CURRENT_VERSION_NO_SUPPORT_ERROR] = "您当前的版本过低，为保障您的资金安全，请您升级APP版本。"

		// 用户相关
		ERR_MSG_MAP[USER_AUTH_FAILED_ERROR] = "账号或者密码错误"
		ERR_MSG_MAP[USER_ACCOUNT_FROZEN_ERROR] = "登录失败次数过多，帐户将被冻结24小时"
		ERR_MSG_MAP[INVALID_MOBILE_ERROR] = "手机号码不合法"
		ERR_MSG_MAP[CAPTCHA_EXPIRE_ERROR] = "验证码过期了"
		ERR_MSG_MAP[INVALID_CAPTCHA_ERROR] = "验证码错误"
		ERR_MSG_MAP[MOBILE_REG_ALREADY_ERROR] = "手机号码已经注册"
		ERR_MSG_MAP[PASSWORD_RESET_FAILURE_ERROR] = "修改密码失败"
		ERR_MSG_MAP[USER_NOT_EXIST_ERROR] = "用户不存在"
		ERR_MSG_MAP[SEND_SMS_TOO_FAST_ERROR] = "发送短信太频繁,60秒内只能发送一次"
		ERR_MSG_MAP[SEND_SMS_MORE_THAN_MAX_TIMES_ERROR] = "短信发送频率超过限制，系统将冻结你的账号"
		ERR_MSG_MAP[USER_ALREADY_REALNAME_AUTH] = "用户已经实名认证"
		ERR_MSG_MAP[USER_IDCARD_ALREADY_REGISTER] = "该身份证号已经被注册"
		ERR_MSG_MAP[USER_PASSWORD_ERROR] = "原登陆密码错误"
		ERR_MSG_MAP[USER_PASSWORD_OLD_EQUAL_NEW_ERROR] = "新旧密码一致"
		ERR_MSG_MAP[USER_PASSWORD_NOT_EQUAL_ERROR] = "两次输入密码不一致"
		ERR_MSG_MAP[USER_INVITE_CODE_INVALID] = "拜年红包无效,请重新输入"
		ERR_MSG_MAP[IMAGE_CAPTCHA_INVALID_ERROR] = "图像验证码错误"
		ERR_MSG_MAP[USER_PASSWORD_INVALID_ERROR] = "输入的密码不合法"
		ERR_MSG_MAP[USER_ALREADY_LOGGED_ERROR] = "已经处于登录状态"
		ERR_MSG_MAP[USER_RECOMMEND_USER_NOT_EXIST_ERROR] = "推荐用户不存在"
		ERR_MSG_MAP[USER_ACCOUNT_NOT_EXIST_ERROR] = "账户不存在"
		ERR_MSG_MAP[USER_PASSWORD_CHECK_FAILED_ERROR] = "密码不正确"
		ERR_MSG_MAP[USER_PAY_PASSWORD_NOT_EXIST_ERROR] = "请先设置支付密码"
		ERR_MSG_MAP[USER_REG_ALREADY_ERROR] = "该用户名已被注册"

		// 文档池相关
		ERR_MSG_MAP[POOL_NOT_EXIST_ERROR] = "文档池不存在"
		ERR_MSG_MAP[POOL_USER_CANT_DELETE_ERROR] = "您没有删除该池的权限"
	}

	if msg, ok := ERR_MSG_MAP[e]; ok {
		return msg
	} else {
		return ""
	}
}
