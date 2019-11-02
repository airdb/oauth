package handlers

import (
	"fmt"
	"github.com/airdb/sailor/enum"

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
	if logincode.Code == "" {
		fmt.Println("wechat code is null")
	}

	userinfo := bo.GetWechatAccessToken(logincode.Code)
	c.JSON(200, vo.LoginResp{
		Nickname:   userinfo.Nickname,
		Headimgurl: userinfo.Headimgurl,
	},
	)

}

func WechatLogout(c *gin.Context) {
	middlewares.SetResp(
		c,
		enum.AirdbSuccess,
		vo.LoginResp{
			Nickname: "john",
		},
	)
}

func SetResp(c *gin.Context, value interface{}) {
	fmt.Print("xxxx", value)
	c.Set("resp", value)
}
