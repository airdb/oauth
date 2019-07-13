package handlers

import (
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	u := "https://open.weixin.qq.com/connect/qrconnect?appid=wxb2a310377819aebd&redirect_uri=https://mina.baobeihuijia.com/auth/v1/wechat/login&response_type=code&scope=snsapi_login&state=bbhj"
	c.Redirect(307, u)
}
