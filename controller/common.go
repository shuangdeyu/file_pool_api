package controller

import (
	"file_pool_api/conf"
	"file_pool_api/helpers"
	"github.com/gin-gonic/gin"
	"helper_go/comhelper"
	"helper_go/pwdhelper"
	"log"
	"strings"
)

/**
 * 获取 token
 */
func GetToken(c *gin.Context) {
	uuid := comhelper.Get_uuid(4)
	token := comhelper.Md5(uuid)
	c.JSON(200, gin.H{
		"e":    0,
		"msg":  "成功",
		"data": token,
	})
}

/**
 * 获取 encrypt
 */
func GetEncrypt(c *gin.Context) {
	action := strings.TrimSpace(c.PostForm("action"))
	timestamp := strings.TrimSpace(c.PostForm("timestamp"))
	encrypt := comhelper.Md5(comhelper.Md5("file_pool{z*2!H&akQ$cM|0_0com") + comhelper.Md5("webapp"+action) + timestamp)
	c.JSON(200, gin.H{
		"e":    0,
		"msg":  "成功",
		"data": encrypt,
	})
}

/**
 * 获取 sign
 */
func GetSign(c *gin.Context) {
	param := strings.TrimSpace(c.PostForm("param"))
	sign := comhelper.Md5(param)
	c.JSON(200, gin.H{
		"e":    0,
		"msg":  "成功",
		"data": sign,
	})
}

/**
 * 加密参数
 */
func EncryptParam(c *gin.Context) {
	param := strings.TrimSpace(c.PostForm("param"))
	if param == "" {
		c.JSON(200, gin.H{
			"e":   conf.INVALID_PARAM_ERROR,
			"msg": "出错啦，请稍后再试",
		})
		return
	}

	pub := helpers.ReadFile(conf.PUBLIC_KEY_PATH)
	encode, err := pwdhelper.RsaEncrypt([]byte(param), pub)
	if err != nil {
		log.Println(err.Error())
		c.JSON(200, gin.H{
			"e":   conf.INVALID_PARAM_ERROR,
			"msg": "出错啦，请稍后再试",
		})
		return
	}
	c.JSON(200, gin.H{
		"e":    0,
		"msg":  "成功",
		"data": encode,
	})
}
