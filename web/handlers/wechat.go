package handlers

import (
	"fmt"
	"github.com/airdb/passport/model/vo"
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
	code := c.PostForm("code")
	fmt.Println("request_method_is", c.Request.Method)
	fmt.Println("code is ", code)
	var logincode vo.LoginReq
	if err := c.ShouldBindQuery(&logincode); err != nil {
		fmt.Println("xxxx", err)
	}

	c.JSON(200, gin.H{
		"message": "pong",
	})
}
