package service

import "helper_go/comhelper"

/**
 * 获取文档列表
 * @param search string 搜索关键字
 * @param offset string 偏移量
 * @param limit string 请求数
 * @param pool_id string 池id
 * @param open string 是否公开
 * @param delete string 是否删除
 */
func GetFileList(search, offset, limit, pool_id, open, delete string) map[string]interface{} {
	data := _call(FilePoolService, "GetFileList", map[string]interface{}{
		"Search":   search,
		"Offset":   comhelper.StringToInt(offset),
		"Limit":    comhelper.StringToInt(limit),
		"PoolId":   comhelper.StringToInt(pool_id),
		"Open":     open,
		"IsDelete": delete,
	})
	return data
}

/**
 * 获取文档点赞数
 * @param file_id string 文档id
 */
func GetFilePraiseCount(file_id string) string {
	data := _call(FilePoolService, "GetFilePraiseCount", map[string]interface{}{
		"FileId": comhelper.StringToInt(file_id),
	})
	return data["data"].(string)
}

/**
 * 获取文档收藏数
 * @param file_id string 文档id
 */
func GetFileCollectCount(file_id string) string {
	data := _call(FilePoolService, "GetFileCollectCount", map[string]interface{}{
		"FileId": comhelper.StringToInt(file_id),
	})
	return data["data"].(string)
}
