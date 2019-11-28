package po

import (
	"fmt"
	"github.com/airdb/sailor/dbutils"
	"github.com/jinzhu/gorm"
	"time"
)

type WechatUserInfo struct {
	gorm.Model
	Openid     string `gorm:"type:varchar(32)"`
	Nickname   string `gorm:"type:varchar(64)"`
	Sex        uint   `gorm:"size:4"`
	Language   string `gorm:"type:varchar(16)"`
	City       string `gorm:"type:varchar(16)"`
	Country    string `gorm:"type:varchar(16)"`
	Headimgurl string `gorm:"type:varchar(128)"`
	Unionid    string `gorm:"type:varchar(32)"`
}

/*
func CreateWechatUserInfoTab() {
	dbutils.WriteDB(DBName).DropTable("wechat_user_info").Debug()
	db := dbutils.WriteDB(DBName).Set("gorm:table_options", "CHARSET=utf8mb4").AutoMigrate(&WechatUserInfo{})
	if db.Error != nil {
		log.Println("error:", db.Error)
	}
}

func AddWechatUserInfo(info *WechatUserInfo) {
	dbutils.WriteDB(DBName).Save(info).Debug()
}
*/

type GitHubUserInfo struct {
	gorm.Model

	Login             string
	GithubID          int
	NodeID            string
	AvatarURL         string
	GravatarID        string
	URL               string
	HTMLURL           string
	FollowersURL      string
	FollowingURL      string
	GistsURL          string
	StarredURL        string
	SubscriptionsURL  string
	OrganizationsURL  string
	ReposURL          string
	EventsURL         string
	ReceivedEventsURL string
	Type              string
	SiteAdmin         bool
	// Name                    interface{}
	Company                 string
	Blog                    string
	Location                string
	Email                   string
	Hireable                bool
	Bio                     string
	PublicRepos             int
	PublicGists             int
	Followers               int
	Following               int
	GithubCreatedAt         time.Time
	GithubUpdatedAt         time.Time
	PrivateGists            int
	TotalPrivateRepos       int
	OwnedPrivateRepos       int
	DiskUsage               int
	Collaborators           int
	TwoFactorAuthentication bool
}

func UpsertGithubUserInfo(info *GitHubUserInfo) {
	fmt.Println("po", info.Login)
	dbutils.DefaultDB().Where(&GitHubUserInfo{
		GithubID: info.GithubID,
		NodeID:   info.NodeID,
	}).Assign(&GitHubUserInfo{
		Model:                   gorm.Model{},
		Login:                   info.Login,
		GithubID:                info.GithubID,
		NodeID:                  info.NodeID,
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
	}).FirstOrCreate(&GitHubUserInfo{})
}
