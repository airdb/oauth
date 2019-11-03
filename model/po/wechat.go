package po

import (
	"github.com/jinzhu/gorm"
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
