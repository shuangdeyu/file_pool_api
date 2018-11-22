package controller

import (
	"file_pool_api/service"
	"github.com/gin-gonic/gin"
	"helper_go/comhelper"
)

/**
 * 获取用户信息
 */
func GetUserInfo(c *gin.Context) {
	ret := service.Get_user_info(1)
	comhelper.Dump(ret)
}
