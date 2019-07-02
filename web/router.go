package web

import (
	"github.com/airdb/passport/web/handlers"
	"log"
	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func NewRouter() *Router {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	v1API := router.Group("/auth/v1")

	wechatAPI := v1API.Group("/wechat")
	wechatAPI.POST("/login", handlers.WechatLogin)
	wechatAPI.GET("/", handlers.Auth)

	return &Router {
		router,
	}
}

func (r *Router) Run() {
	err := r.Engine.Run()  // listen and serve on 0.0.0.0:8080
	if err != nil {
		log.Fatal("start failed")
	}
}


func Run() {
	NewRouter().Run()
}

