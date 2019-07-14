package handlers

import (
	"fmt"
	"github.com/airdb/passport/model/bo"
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
	var logincode vo.LoginReq
	if err := c.ShouldBindQuery(&logincode); err != nil {
		fmt.Println("xxxx", err)
	}

	fmt.Println("code: ", logincode.Code)
	if logincode.Code != "" {
		bo.GetWechatAccessToken(logincode.Code)
	}

	c.JSON(200, vo.LoginResp{
		Nickname: "john",
	},
	)

}

func WechatLogout(c *gin.Context) {
	SetResp(
		c,
		vo.LoginResp{
			Nickname: "john",
		},
	)
}

func SetResp(c *gin.Context, value interface{}) {
	fmt.Print("xxxx", value)
	c.Set("resp", value)
}
