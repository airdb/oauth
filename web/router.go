package web

import (
	"fmt"

	"github.com/airdb/sailor/config"
	"github.com/gin-gonic/gin"
)

func Run() {
	config.Init()

	err := NewRouter().Run(config.GetDefaultBindAddress())
	if err != nil {
		fmt.Println("error: ", err)
	}
}

func NewRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	oauth := router.Group("/")
	oauth.GET("/", IndexHandler)

	v1API := router.Group("/apis/oauth/v1")

	v1API.GET("/:provider", Login)
	v1API.GET("/:provider/callback", Callback)

	return router
}
