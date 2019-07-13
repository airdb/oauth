package bo

import (
	"fmt"
	"github.com/airdb/passport/model/po"
	"github.com/imroc/req"
)

var (
	weapp *po.Secret
)

type WechatAccessTokenResp struct {
	AccessToken  string
	ExpiresIn    string
	RefreshToken string
	Openid       string
	Scope        string
	Unionid      string
	Errcode      int
	Errmsg       string
}

func ini() {
	weapp = po.QuerySecret("wechat")
}
func GetRewriteURI() string {

	return fmt.Sprintf("%s?appid=%s&redirect_uri=%s&&response_type=code&scope=snsapi_login&state=%s",
		weapp.URL,
		weapp.Appid,
		weapp.RedirectURI,
		weapp.State,
	)
}

func GetWechatAccessToken(code string) {
	url := fmt.Sprintf("%s?appid=%s&secret=%s&code=%s&grant_type=authorization_code",
		"https://api.weixin.qq.com/sns/oauth2/access_token",
		weapp.Appid,
		weapp.Secret,
		code,
	)

	r, _ := req.Get(url)

	weinfo := &WechatAccessTokenResp{}
	r.ToJSON(&weinfo)

	fmt.Println("xxxx", weinfo)
}

func GetWechatUserInfo() {
}
