package bo

import (
	"fmt"
	"github.com/airdb/passport/model/po"
)

func GetRewriteURI() string {
	weapp := po.QuerySecret("wechat")

	return fmt.Sprintf("%s?appid=%s&redirect_uri=%s&&response_type=code&scope=snsapi_login&state=%s",
		"https://open.weixin.qq.com/connect/qrconnect",
		weapp.Appid,
		weapp.RedirectURI,
		weapp.State,
	)

}
