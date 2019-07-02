package handlers

import (
    "github.com/gin-gonic/gin"
    "net/http"
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
    u := "https://open.weixin.qq.com/connect/qrconnect?appid=wxbdc5610cc59c1631&redirect_uri=https%3A%2F%2Fpassport.yhd.com%2Fwechat%2Fcallback.do&response_type=code&scope=snsapi_login&state=3d6be0a4035d839573b04816624a415e#wechat_redirect"
    c.Redirect(http.StatusMovedPermanently, u)
}