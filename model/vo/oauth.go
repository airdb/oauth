package vo

import (
	"fmt"
	"github.com/airdb/passport/model/po"
	"github.com/imroc/req"
	"log"
	"time"
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

type GithubResp struct {
	AccessToken      string `json:"access_token"`
	Scope            string `json:"scope"`
	TokenType        string `json:"token_type"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	ErrorUri         string `json:"error_uri"`
}

func GetUserInfoFromOauth(provider string, code, state string) *UserInfo {
	switch provider {
	case ProviderGithub:
		return GithubUserInfo(code, state)
	}
	return nil
}

func GithubUserInfo(code, state string) *UserInfo {
	p := QueryProvider(ProviderGithub)

	param := req.Param{
		"client_id":     p.ClientID,
		"client_secret": p.ClientSecret,
		"code":          code,
		"state":         state,
		"redirect_uri":  p.RedirectURI,
	}

	header := req.Header{
		"Accept": "application/json",
	}

	apiurl := "https://github.com/login/oauth/access_token"
	r, err := req.Post(apiurl, header, param)
	if err != nil {
		log.Fatal(err)
	}
	var resp GithubResp
	// m, err := url.ParseQuery(r.String())

	// resp.AccessToken = m.Get("access_token")
	// resp.Error = m.Get("error")
	err = r.ToJSON(&resp)
	if err != nil {
		fmt.Println("access_token:", resp.AccessToken)
		fmt.Println("error", resp.Error)
		return GetUserInfo(resp.AccessToken)
	}

	return nil
}

func GetUserInfo(accessToken string) *UserInfo {
	token := "token " + accessToken
	header := req.Header{
		"Accept":        "application/json",
		"Authorization": token,
	}

	apiurl := "https://api.github.com/user"
	r, err := req.Get(apiurl, header)
	if err != nil {
		log.Fatal(err)
	}
	var info UserInfo
	fmt.Println("userinfo: ", r)
	err = r.ToJSON(&info)
	if err != nil {
		return &info
	}

	return nil
}

func LoginRecord(userInfo *UserInfo) {
	po.UpsertGithubUserInfo(ToPoGitHubUserInfo(userInfo))
}

func ToPoGitHubUserInfo(info *UserInfo) *po.GitHubUserInfo {
	return &po.GitHubUserInfo{
		Login:                   info.Login,
		GithubID:                info.ID,
		NodeID:                  "",
		AvatarURL:               "",
		GravatarID:              "",
		URL:                     "",
		HTMLURL:                 "",
		FollowersURL:            "",
		FollowingURL:            "",
		GistsURL:                "",
		StarredURL:              "",
		SubscriptionsURL:        "",
		OrganizationsURL:        "",
		ReposURL:                "",
		EventsURL:               "",
		ReceivedEventsURL:       "",
		Type:                    "",
		SiteAdmin:               false,
		Company:                 "",
		Blog:                    "",
		Location:                "",
		Email:                   "",
		Hireable:                false,
		Bio:                     "",
		PublicRepos:             0,
		PublicGists:             0,
		Followers:               0,
		Following:               0,
		GithubCreatedAt:         time.Time{},
		GithubUpdatedAt:         time.Time{},
		PrivateGists:            0,
		TotalPrivateRepos:       0,
		OwnedPrivateRepos:       0,
		DiskUsage:               0,
		Collaborators:           0,
		TwoFactorAuthentication: false,
	}
}
