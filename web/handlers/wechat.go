package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
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
	code := c.PostForm("code")
	fmt.Fprintln(os.Stderr, "hello world", code)

	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func Auth(c *gin.Context) {
	u := "https://open.weixin.qq.com/connect/qrconnect?appid=wxb2a310377819aebd&redirect_uri=https://mina.baobeihuijia.com/auth/v1/wechat/login&response_type=code&scope=snsapi_login&state=bbhj"
	fmt.Println("xxx rewrite")
	fmt.Fprintln(os.Stderr, "hello world", code)
	c.Redirect(307, u)
}
