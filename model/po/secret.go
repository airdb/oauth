package po

import (
	"github.com/airdb/sailor/dbutils"
	"github.com/jinzhu/gorm"
)

const (
	DBName = "dev_db"
)

type Secret struct {
	gorm.Model
	Provider     string `gorm:"type:varchar(16)"`
	ClientID     string `gorm:"type:varchar(64)"`
	ClientSecret string `gorm:"type:varchar(128)"`
	URL          string `gorm:"type:varchar(64)"`
	RedirectURI  string `gorm:"type:varchar(64)"`
	State        string `gorm:"type:varchar(16)"`
}

func ListProvider() (secret []*Secret) {
	dbutils.DefaultDB().Table("secret_tab").Find(&secret)
	return
}

func QueryProvider(provider string) *Secret {
	var secret Secret

	dbutils.DefaultDB().Table("secret_tab").Where("provider = ?", provider).First(&secret)

	return &secret
}

/*
func CreateSecretTab() {
	dbutils.WriteDB(DBName).DropTable("secret").Debug()
	db := dbutils.WriteDB(DBName).Set("gorm:table_options", "CHARSET=utf8mb4").AutoMigrate(&Secret{}).Debug()
	if db.Error != nil {
		fmt.Print(db.Error)
	}
	// db := dbutils.WriteDB(DBName).Debug().FirstOrCreate(&aa)
}

func QuerySecret(typ string) *Secret {
	secret := Secret{}
	dbutils.WriteDB(DBName).First(&secret, "typ = ?", typ)

	return &secret
}
*/
