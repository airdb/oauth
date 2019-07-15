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
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Openid       string `json:"openid"`
	Scope        string `json:"scope"`
	Unionid      string `json:"unionid"`
	Errcode      int    `json:"errcode"`
	Errmsg       string `json:"errmsg"`
}

type WechatUserInfo struct {
	Openid     string
	Nickname   string
	Sex        uint
	Language   string
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

func GetWechatAccessToken(code string) *WechatUserInfo{
	url := fmt.Sprintf("%s?appid=%s&secret=%s&code=%s&grant_type=authorization_code",
		"https://api.weixin.qq.com/sns/oauth2/access_token",
		weapp.Appid,
		weapp.Secret,
		code,
	)

	r, _ := req.Get(url)

	weinfo := &WechatAccessTokenResp{}
	fmt.Println("access_token: ", r)
	err := r.ToJSON(&weinfo)
	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println("access_token: ", weinfo.AccessToken)
	fmt.Println("access_token: ", weinfo.RefreshToken)
	if weinfo.Errmsg != "" {
		fmt.Println("error_info", weinfo.Errcode, weinfo.Errmsg)
	}
	fmt.Println("xxxx", weinfo.AccessToken, weinfo.Openid)
	return GetWechatUserInfo(weinfo)
}

func GetWechatUserInfo(weinfo *WechatAccessTokenResp) *WechatUserInfo {
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
	err = r.ToJSON(&userinfo)
	if err != nil {
		fmt.Println("error: ", err)
	}

	info := po.WechatUserInfo{
		Openid:     userinfo.Openid,
		Nickname:   userinfo.Nickname,
		Sex:        userinfo.Sex,
		Language:   userinfo.Language,
		City:       userinfo.City,
		Country:    userinfo.Country,
		Headimgurl: userinfo.Headimgurl,
		Unionid:    userinfo.Unionid,
	}

	po.AddWechatUserInfo(&info)
	return &userinfo
}
