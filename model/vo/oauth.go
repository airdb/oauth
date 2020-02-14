package vo

import (
	"github.com/airdb/passport/model/po"
)

const (
	ProviderGithub = "github"
	ProviderWechat = "wechat"
)

type ProviderSecret struct {
	Provider     string `json:"provider"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectURI  string `json:"redirect_uri"`
	URL          string `json:"url"`
}

type LoginReq struct {
	Code  string `form:"code"`
	State string `form:"state"`
}

type LoginResp struct {
	Nickname   string `json:"nickname"`
	Headimgurl string `json:"headimgurl"`
}

type User struct {
	ID        string                 `json:"id"`
	Username  string                 `json:"username"`
	FirstName string                 `json:"first_name"`
	LastName  string                 `json:"last_name"`
	FullName  string                 `json:"full_name"`
	Email     string                 `json:"email"`
	Avatar    string                 `json:"avatar"`
	Raw       map[string]interface{} `json:"raw"` // Raw data
}

func FromPoProviderSecret(poSecret *po.Secret) *ProviderSecret {
	return &ProviderSecret{
		Provider:     poSecret.Provider,
		ClientID:     poSecret.ClientID,
		ClientSecret: poSecret.ClientSecret,
		RedirectURI:  poSecret.RedirectURI,
		URL:          poSecret.URL,
	}
}

func FromPoProviderSecrets(poSecrets []*po.Secret) (secrets []*ProviderSecret) {
	for _, secret := range poSecrets {
		secrets = append(secrets, FromPoProviderSecret(secret))
	}

	return
}

func ListProvider() []*ProviderSecret {
	ProviderSecrets := FromPoProviderSecrets(po.ListProvider())

	return ProviderSecrets
}

func QueryProvider(name string) *ProviderSecret {
	return FromPoProviderSecret(po.QueryProvider(name))
}

func GetUserInfoFromOauth(provider string, code, state string) *GithubUserInfo {
	switch provider {
	case ProviderGithub:
		return GetGithubUserInfo(code, state)
	case ProviderWechat:
		return nil
	}

	return nil
}

func GetAuthRedirectURL(provider string) (*string, error) {
	providerData := QueryProvider(provider)

	switch provider {
	case ProviderGithub:
		providerData.URL = "https://github.com/login/oauth/authorize?&response_type=code&scope=snsapi_login&scope=user"
		return GetGithubAuthRedirectURL(providerData)
	case ProviderWechat:
		providerData.URL = "https://open.weixin.qq.com/connect/qrconnect?&response_type=code&scope=snsapi_login"
		return GetWechatAuthRedirectURL(providerData)
	}

	return nil, nil
}
