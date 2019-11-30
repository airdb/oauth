package handlers

import (
	"fmt"
	"github.com/airdb/passport/model/vo"
	"github.com/airdb/sailor/enum"
	"github.com/airdb/sailor/gin/middlewares"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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
	provider := c.Param("provider")

	var logincode vo.LoginReq
	if err := c.ShouldBindQuery(&logincode); err != nil {
		fmt.Println("xxxx", err)
	}

	fmt.Println("provider", provider, logincode)

	userInfo := vo.GetUserInfoFromOauth(provider, logincode.Code, logincode.State)
	fmt.Println("get user info", userInfo)
	if userInfo == nil {
		middlewares.SetResp(
			c,
			enum.AirdbFailed,
			vo.LoginResp{
				Nickname:   "null",
				Headimgurl: "null",
			},
		)

		return
	}

	middlewares.SetResp(
		c,
		enum.AirdbSuccess,
		vo.LoginResp{
			Nickname:   userInfo.Login,
			Headimgurl: userInfo.AvatarURL,
		},
	)
}

func Login(c *gin.Context) {
	provider := c.Param("provider")
	/*
		authURL, err := NewDispatcher().New().Driver(provider).Redirect(
			providerData.ClientID,
			providerData.ClientSecret,
			providerData.RedirectURI,
		)
	*/
	authURL, err := vo.GetAuthRedirectURL(provider)
	fmt.Print("authURL", err, authURL)
	if err != nil {
		c.Redirect(http.StatusFound, "/")
	}
	// authURL := fmt.Sprintf("%s?appid=%s&redirect_uri=%s&response_type=code&scope=snsapi_login&state=bbhj", providerData.URL, providerData.ClientID, providerData.RedirectURI)

	fmt.Println(authURL)
	c.Redirect(http.StatusFound, *authURL)
}
