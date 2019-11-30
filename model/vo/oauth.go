package vo

import (
	"fmt"
	"github.com/airdb/passport/model/po"
)

type ProviderSecret struct {
	Provider     string `json:"provider"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectURI  string `json:"redirect_uri"`
	URL          string `json:"url"`
}

const (
	ProviderGithub = "github"
	ProviderWechat = "wechat"
)

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
		fmt.Print("get user info from: ", ProviderGithub)
		return GetGithubUserInfo(code, state)
	}
	return nil
}
