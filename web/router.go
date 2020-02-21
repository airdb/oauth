package web

import (
	"fmt"
	"io"
	"log"
	"net/http/httptest"

	"github.com/airdb/sailor/config"
	"github.com/gin-gonic/gin"
)

func Run() {
	address := config.GetDefaultBindAddress()
	log.Println("web server start at", address)
	err := NewRouter().Run(address)
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

	v1API.GET("/wechat/login", WechatLogin)
	// v1API.GET("/:provider", Login)
	// v1API.GET("/:provider/callback", Callback)

	return router
}

func APIRequest(uri, method string, param io.Reader) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, uri, param)

	if method == "GET" {
		req.Header.Set("Content-Type", "application/json")
	} else if method == "POST" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	w := httptest.NewRecorder()
	NewRouter().ServeHTTP(w, req)

	return w
}
