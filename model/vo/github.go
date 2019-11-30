package vo

import (
	"fmt"
	"github.com/airdb/passport/model/po"
	"github.com/imroc/req"
	"log"
	"time"
)

type GithubAccessTokenResp struct {
	AccessToken      string `json:"access_token"`
	Scope            string `json:"scope"`
	TokenType        string `json:"token_type"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	ErrorUri         string `json:"error_uri"`
}

type GithubUserInfo struct {
	Login                   string      `json:"login"`
	ID                      int         `json:"id"`
	NodeID                  string      `json:"node_id"`
	AvatarURL               string      `json:"avatar_url"`
	GravatarID              string      `json:"gravatar_id"`
	URL                     string      `json:"url"`
	HTMLURL                 string      `json:"html_url"`
	FollowersURL            string      `json:"followers_url"`
	FollowingURL            string      `json:"following_url"`
	GistsURL                string      `json:"gists_url"`
	StarredURL              string      `json:"starred_url"`
	SubscriptionsURL        string      `json:"subscriptions_url"`
	OrganizationsURL        string      `json:"organizations_url"`
	ReposURL                string      `json:"repos_url"`
	EventsURL               string      `json:"events_url"`
	ReceivedEventsURL       string      `json:"received_events_url"`
	Type                    string      `json:"type"`
	SiteAdmin               bool        `json:"site_admin"`
	Name                    interface{} `json:"name"`
	Company                 string      `json:"company"`
	Blog                    string      `json:"blog"`
	Location                string      `json:"location"`
	Email                   string      `json:"email"`
	Hireable                bool        `json:"hireable"`
	Bio                     string      `json:"bio"`
	PublicRepos             int         `json:"public_repos"`
	PublicGists             int         `json:"public_gists"`
	Followers               int         `json:"followers"`
	Following               int         `json:"following"`
	CreatedAt               time.Time   `json:"created_at"`
	UpdatedAt               time.Time   `json:"updated_at"`
	PrivateGists            int         `json:"private_gists"`
	TotalPrivateRepos       int         `json:"total_private_repos"`
	OwnedPrivateRepos       int         `json:"owned_private_repos"`
	DiskUsage               int         `json:"disk_usage"`
	Collaborators           int         `json:"collaborators"`
	TwoFactorAuthentication bool        `json:"two_factor_authentication"`
	Plan                    struct {
		Name          string `json:"name"`
		Space         int    `json:"space"`
		Collaborators int    `json:"collaborators"`
		PrivateRepos  int    `json:"private_repos"`
	} `json:"plan"`
}

func GetGithubUserInfo(code, state string) *GithubUserInfo {
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
	var resp GithubAccessTokenResp
	err = r.ToJSON(&resp)
	if err != nil {
		fmt.Println("access_token:", resp)
		fmt.Println("error", resp.Error)
		return nil
	}

	return GetUserInfo(resp.AccessToken)
}

func GetUserInfo(accessToken string) *GithubUserInfo {
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
	var info GithubUserInfo
	err = r.ToJSON(&info)
	fmt.Println("access userinfo: ", err, r)
	fmt.Println("access userinfo: ", info.Login)
	if err == nil {
		return &info
	}

	return nil
}

func LoginRecord(userInfo *GithubUserInfo) {
	po.UpsertGithubUserInfo(ToPoGitHubUserInfo(userInfo))
}

func ToPoGitHubUserInfo(info *GithubUserInfo) *po.GitHubUserInfo {
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

func GetGithubAuthRedirectURL(providerData *ProviderSecret) (*string, error) {
	state := randToken()
	authURL := fmt.Sprintf("%s&client_id=%s&redirect_uri=%s&state=%s",
		"https://github.com/login/oauth/authorize?&response_type=code&scope=snsapi_login&scope=user",
		providerData.ClientID,
		providerData.RedirectURI,
		state,
	)

	return &authURL, nil
}
