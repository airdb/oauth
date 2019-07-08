package main

import (
	"fmt"
	"github.com/airdb/passport/web"
	"github.com/airdb/sailor/config"
	// "github.com/airdb/sailor/dbutils"
)

func main() {
	fmt.Println(config.GetEnv())
	fmt.Println(config.Get("wechat"))
	fmt.Println(config.GetString("api"))
	fmt.Println(config.AllSettings())
	cf := config.AllSettings()
	fmt.Println("xxxx", cf["wechat"])
	for _, a := range config.GetDatabases() {
		fmt.Println(a)
	}
	// fmt.Println(config.GetDatabases())
	
	// dbutils.()

	web.Run()
}
