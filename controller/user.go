package controller

import (
	"file_pool_api/service"
	"github.com/gin-gonic/gin"
)

/**
 * 获取用户信息
 */
func GetUserInfo(c *gin.Context) {
	ret := service.Get_user_info(1)
	c.JSON(200, gin.H{
		"e":    0,
		"msg":  "success",
		"data": ret["data"],
	})

	if ret["error"] != service.SERVICE_SUCCESS {

	}
}
