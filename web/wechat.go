package web

import (
	"log"
	"net/http"
	"os"

	"github.com/airdb/passport/model/vo"
	"github.com/airdb/sailor"
	"github.com/airdb/sailor/enum"
	"github.com/esap/wechat"
	"github.com/gin-gonic/gin"
)

func WechatLogin(c *gin.Context) {
	var req vo.WechatLoginReq
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, &sailor.HTTPAirdbResponse{
			Code:    enum.AirdbUndefined,
			Success: false,
			Data:    nil,
			Error:   "不合法的请求参数",
		})

		return
	}

	wechat.Debug = true
	// wechat.GetAccessToken()
	ss, err := wechat.New(
		os.Getenv("WECHAT_TOKEN"),
		os.Getenv("WECHAT_APPID"),
		os.Getenv("WECHAT_APPSECRET"),
	).Jscode2Session(req.Code)

	log.Println(ss.UserId, err)

	c.JSON(http.StatusOK, sailor.HTTPAirdbResponse{
		Code:    enum.AirdbSuccess,
		Success: true,
		Data:    "xxx",
	})
}
