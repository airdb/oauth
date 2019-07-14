package bo

import (
	"fmt"
	"github.com/airdb/passport/model/po"
	"github.com/imroc/req"
)

var (
	weapp  *po.Secret
	weinfo *WechatAccessTokenResp
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

type WechatUserInfo struct {
	Openid     string
	Nickname   string
	Sex        uint
	Province   string
	City       string
	Country    string
	Headimgurl string
	Privilege  []string
	Unionid    string
}

func init() {
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

	if weinfo.Errmsg != "" {
		fmt.Println("error_info", weinfo.Errcode, weinfo.Errmsg)
	}
	fmt.Println("xxxx", weinfo.AccessToken, weinfo.Openid)
	GetWechatUserInfo(weinfo)
}

func GetWechatUserInfo(weinfo *WechatAccessTokenResp) {
	fmt.Println("get_wechat_user_info", weinfo.Openid)
	param := req.Param{
		"access_token": weinfo.AccessToken,
		"openid":       weinfo.Openid,
	}
	r, err := req.Get("https://api.weixin.qq.com/sns/userinfo", param)
	if err != nil {
		fmt.Println("get_userinfo_failed, ", err)
	}

	var userinfo WechatUserInfo
	r.ToJSON(&userinfo)

	fmt.Println(userinfo, r)
}
