package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func Auth(c *gin.Context) {
	u := "https://open.weixin.qq.com/connect/qrconnect?appid=wxb2a310377819aebd&redirect_uri=https://mina.baobeihuijia.com/auth/v1/wechat/login&response_type=code&scope=snsapi_login&state=bbhj"
	fmt.Println("xxx rewrite")
	fmt.Fprintln(os.Stderr, "hello  rewrite")
	c.Redirect(307, u)
}
