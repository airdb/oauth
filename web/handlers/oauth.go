package handlers

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/airdb/passport/model/vo"
	"github.com/gin-gonic/gin"
	"github.com/airdb/sailor/config"
)

// Show homepage with login URL
func IndexHandler(c *gin.Context) {
	msg := "<html><head><title>Airdb Passport</title></head><body>"
	msg += "<a href='/oauth/v1/github'><button>Login with GitHub</button></a><br>"
	msg += "<a href='/oauth/v1/linkedin'><button>Login with LinkedIn</button></a><br>"
	msg += "<a href='/oauth/v1/facebook'><button>Login with Facebook</button></a><br>"
	msg += "<a href='/oauth/v1/google'><button>Login with Google</button></a><br>"
	msg += "<a href='/oauth/v1/bitbucket'><button>Login with Bitbucket</button></a><br>"
	msg += "<a href='/oauth/v1/amazon'><button>Login with Amazon</button></a><br>"
	msg += "<a href='/oauth/v1/slack'><button>Login with Slack</button></a><br>"
	msg += "<a href='/oauth/v1/wechat'><button>Login with Wechat</button></a><br>"
	msg += "</body></html>"
	_, err := c.Writer.Write([]byte(msg))
	if err != nil {
		log.Println(err)
	}
}

// Handle callback of provider
func Callback(c *gin.Context) {
	var logincode vo.LoginReq
	if err := c.ShouldBindQuery(&logincode); err != nil {
		fmt.Println("xxxx", err)
	}

	fmt.Println("code: ", logincode.Code)
	if logincode.Code == "" {
		fmt.Println("wechat code is null")
	}

	// userinfo := bo.GetWechatAccessToken(logincode.Code)
	c.JSON(200, vo.LoginResp{
		Nickname:   "dean",
		Headimgurl: "img",
	},
	)
}

func Redirect(c *gin.Context) {
	provider := c.Param("provider")

	fmt.Println("xxxxx", c.Request.RequestURI)
	fmt.Println("xxxxx", c.Request.Host)
	providerData := vo.QueryProvider(provider)

	scheme := "http"
	if config.GetEnv() == "live" {
		scheme = "https"
	}

	reqURL := url.URL{
		Scheme: scheme,
		Host:   c.Request.Host,
		Path:   c.Request.RequestURI + "/callback",
	}

	fmt.Println("xxxxxlllll", reqURL.String())
	/*
		authURL, err := NewDispatcher().New().Driver(provider).Redirect(
			providerData.ClientID,
			providerData.ClientSecret,
			providerData.RedirectURI,
		)
	*/
	authURL := fmt.Sprintf("%s?appid=%s&redirect_uri=%s&response_type=code&scope=snsapi_login&state=bbhj", providerData.URL, providerData.ClientID, providerData.RedirectURI)

	fmt.Println(authURL)
	c.Redirect(http.StatusFound, authURL)
}
