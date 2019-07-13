package main

import (
	"github.com/airdb/passport/web"
	"github.com/airdb/passport/model/po"
)

func main() {
	po.Hello()
	web.Run()
}
