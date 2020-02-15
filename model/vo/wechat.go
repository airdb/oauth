package vo

import (
	"fmt"
	"log"

	"github.com/imroc/req"
)

type WechatAccessTokenResp struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Openid       string `json:"openid"`
	Scope        string `json:"scope"`
	Unionid      string `json:"unionid"`
	Errcode      int    `json:"errcode,omitempty"`
	Errmsg       string `json:"errmsg,omitempty"`
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

type WechatLoginReq struct {
	Code string `form:"code"`
}

type WechatLoginResp struct {
}

func GetWechatAuthRedirectURL(providerData *ProviderSecret) (*string, error) {
	state := randToken()
	authURL := fmt.Sprintf("%s&appid=%s&redirect_uri=%s&state=%s",
		providerData.URL,
		providerData.ClientID,
		providerData.RedirectURI,
		state,
	)

	return &authURL, nil
}

func GetWechatAccessToken(code, state string) *WechatUserInfo {
	p := QueryProvider(ProviderWechat)

	param := req.Param{
		"appid":        p.ClientID,
		"secret":       p.ClientSecret,
		"code":         code,
		"state":        state,
		"redirect_uri": p.RedirectURI,
		"grant_type":   "authorization_code",
	}

	header := req.Header{
		"Accept": "application/json",
	}

	apiurl := "https://api.weixin.qq.com/sns/oauth2/access_token"

	r, err := req.Get(apiurl, header, param)
	if err != nil {
		log.Fatal(err)
	}

	weinfo := &WechatAccessTokenResp{}

	fmt.Println("access_token: ", r)

	err = r.ToJSON(&weinfo)
	if err != nil {
		fmt.Println("error", weinfo)
		return nil
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

	/*
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
	*/

	return &userinfo
}
