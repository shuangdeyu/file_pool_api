package controller

import (
	"file_pool_api/service"
	"helper_go/comhelper"
	"helper_go/timehelper"
)

/**
 * 获取公共文档列表 200
 */
func FileList(AppInitParam *AppParam) map[string]interface{} {
	param := AppInitParam.RequestParam
	search := comhelper.DefaultParam(param["search"], "")
	offset := comhelper.DefaultParam(param["offset"], "0")
	limit := comhelper.DefaultParam(param["limit"], "12")

	ret := service.GetFileList(search, offset, limit, "0", "", "")
	data := ret["data"].(map[string]interface{})
	list := data["list"].([]interface{})
	count := data["count"]
	tmp := []interface{}{}
	if len(list) > 0 {
		for _, v := range list {
			d := v.(map[string]interface{})
			// 获取点赞数
			parise_count := service.GetFilePraiseCount(d["id"].(string))
			// 获取收藏数
			collect_count := service.GetFileCollectCount(d["id"].(string))

			tmp = append(tmp, map[string]interface{}{
				"id":            d["id"],
				"name":          d["name"],
				"content":       d["content"],
				"create_time":   d["create_time"],
				"create_date":   timehelper.TimeToDate(d["create_time"].(string)),
				"praise_count":  parise_count,
				"collect_count": collect_count,
			})
		}
	}
	return map[string]interface{}{
		"count": count,
		"list":  tmp,
	}
}
