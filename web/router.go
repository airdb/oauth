package web

import (
	"log"

	"github.com/airdb/passport/web/handlers"
	"github.com/gin-gonic/gin"
	"github.com/airdb/sailor/gin/middlewares"
)

type Router struct {
	*gin.Engine
}

func NewRouter() *Router {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(
		middlewares.ToJSON("v1"),
	)

	auth := router.Group("/")

	auth.GET("/", handlers.Auth)
	auth.HEAD("/", handlers.Auth)

	v1API := router.Group("/auth/v1")
	wechatAPI := v1API.Group("/wechat")
	wechatAPI.POST("/login", handlers.WechatLogin)
	wechatAPI.GET("/login", handlers.WechatLogin)
	wechatAPI.GET("/", handlers.Auth)
	wechatAPI.HEAD("/", handlers.Auth)

	wechatAPI.GET("/logout", handlers.WechatLogout)
	return &Router{
		router,
	}
}

func (r *Router) Run() {
	err := r.Engine.Run() // listen and serve on 0.0.0.0:8080
	if err != nil {
		log.Fatal("start failed")
	}
}

func Run() {
	NewRouter().Run()
}
