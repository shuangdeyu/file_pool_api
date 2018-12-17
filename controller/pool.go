package controller

import (
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

}
