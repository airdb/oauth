package vo

import (
	"crypto/rand"
	"encoding/base64"
	"log"

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
	}
	return nil
}

func GetOauthRedirectURL(provider string) (*string, error) {
	// 	providerData := vo.QueryProvider(provider)
	providerData := QueryProvider(provider)
	switch provider {
	case ProviderGithub:
		return GetGithubAuthRedirectURL(providerData)
	}
	return nil, nil
}

// Generate a random token
func randToken() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		log.Println(err)
	}
	return base64.StdEncoding.EncodeToString(b)
}
