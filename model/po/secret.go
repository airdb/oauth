package po

import (
	"github.com/jinzhu/gorm"
	"github.com/airdb/sailor/dbutils"
)

const (
	DBName  = "passport_db"
)

type Secret struct {
	gorm.Model
	typ string `gorm:"type:varchar("16")"`
	appid string `gorm:"type:varchat("64")"`
	secret string `gorm:"type:varchat("128")`
}

func Hello() {
	var aa Secret
	dbutils.WriteDB(DBName).First(&aa)
}
