package vo

import (
	"fmt"
	"log"
	"net/url"

	"github.com/airdb/passport/model/po"
	"github.com/imroc/req"
)

type ProviderSecret struct {
	Provider     string `json:"provider"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectURI  string `json:"redirect_uri"`
	URL          string `json:"url"`
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

type GithubResp struct {
	AccessToken      string `json:"access_token"`
	Scope            string `json:"scope"`
	TokenType        string `json:"token_type"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	ErrorUri         string `json:"error_uri"`
}

func GithubUserInfo(provider string, code, state string) error {
	p := QueryProvider(provider)

	param := req.Param{
		"client_id":     p.ClientID,
		"client_secret": p.ClientSecret,
		"code":          code,
		"state":         state,
		"redirect_uri":  p.RedirectURI,
	}

	apiurl := "https://github.com/login/oauth/access_token"
	r, err := req.Post(apiurl, "", param)
	if err != nil {
		log.Fatal(err)
	}
	var resp GithubResp
	m, err := url.ParseQuery(r.String())

	resp.AccessToken = m.Get("access_token")
	resp.Error = m.Get("error")
	fmt.Println("access_token:", resp.AccessToken)
	fmt.Println("error", resp.Error)
	return err
}
