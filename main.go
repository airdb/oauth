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
	fmt.Println(config.GetString("wechat.api"))
	for _, a := range config.GetDatabases() {
		fmt.Println(a)
	}
	// fmt.Println(config.GetDatabases())
	
	// dbutils.()

	web.Run()
}
