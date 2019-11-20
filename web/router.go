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

	oauth := router.Group("/")
	oauth.GET("/", handlers.IndexHandler)

	v1API := router.Group("/oauth/v1")
	v1API.Use(
		middlewares.Jsonifier(),
	)
	v1API.GET("/:provider", handlers.Redirect)
	v1API.GET("/:provider/callback", handlers.Callback)

	return router
}
