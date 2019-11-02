package web

import (
	"fmt"

	"github.com/airdb/passport/web/handlers"
	"github.com/airdb/sailor/config"
	"github.com/airdb/sailor/gin/middlewares"
	"github.com/gin-gonic/gin"
)

func Run() {
	config.Init()
	fmt.Printf("Env: %s, Port: %s\n", config.GetEnv(), config.GetPort())
	err := NewRouter().Run("0.0.0.0:" + config.GetPort())
	if err != nil {
		fmt.Println("error: ", err)
	}
}

func NewRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(
		middlewares.Jsonifier(),
	)

	auth := router.Group("/")
	auth.GET("/", handlers.IndexHandler)
	auth.HEAD("/", handlers.Auth)

	v1API := router.Group("/auth/v1")
	v1API.GET("/:provider", handlers.Redirect)
	v1API.GET("/:provider/callback", handlers.Callback)

	/*
		authAPI := v1API.Group("/wechat")

		authAPI.POST("/login", handlers.WechatLogin)
		authAPI.GET("/login", handlers.WechatLogin)
		authAPI.GET("/", handlers.Auth)
		authAPI.HEAD("/", handlers.Auth)
		authAPI.GET("/logout", handlers.WechatLogout)
	*/
	return router
}
