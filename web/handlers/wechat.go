package handlers

import (
	"fmt"
	"github.com/airdb/passport/model/bo"
	"github.com/airdb/passport/model/vo"
	"github.com/airdb/sailor/gin/middlewares"
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
	var logincode vo.LoginReq
	if err := c.ShouldBindQuery(&logincode); err != nil {
		fmt.Println("xxxx", err)
	}

	fmt.Println("code: ", logincode.Code)
	if logincode.Code != "" {
		bo.GetWechatAccessToken(logincode.Code)
	}

	middlewares.SetResp(
		c,
		vo.LoginResp{
			Nickname: "John",
		},
	)
}
