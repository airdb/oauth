package handlers

import (
	"github.com/gin-gonic/gin"
)

// Wechat login godoc
// @Summary wechat login
// @Description Wechat login
// @Tags wechat login
// @Accept  json
// @Produce  json
// @Param
// @Success 200 {object} vo.ListenerResp
// @Router /wechat/login [post]
func WechatLogin(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func Auth(c *gin.Context) {
	u := "https://open.weixin.qq.com/connect/qrconnect?appid=wxb2a310377819aebd&redirect_uri=https://mina.baobeihuijia.com/dev/lastest/wechatapi/wechat/login&response_type=code&scope=snsapi_login&state=bbhj"
	log.Println("xxx rewrite")
	c.Redirect(302, u)
}
