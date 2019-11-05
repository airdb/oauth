package handlers

import (
	"fmt"
	"log"
	"net/url"

	"github.com/airdb/passport/model/vo"
	"github.com/gin-gonic/gin"
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
	// Retrieve query params for state and code
	state := c.Query("state")
	code := c.Query("code")

	// Handle callback and check for errors
	user, token, err := NewDispatcher().Handle(state, code)
	if err != nil {
		fmt.Println("err", err, token)
		c.String(200, err.Error())
		return
	}

	// Print in terminal user information
	fmt.Printf("%#v", token)
	fmt.Printf("%#v", user)

	// If no errors, show provider name
	c.String(200, "HI, "+user.FullName)
}

func Redirect(c *gin.Context) {
	provider := c.Param("provider")

	providerData := vo.QueryProvider()
	fmt.Println("xxx", providerData)

	reqURL := url.URL{
		Scheme: c.Request.URL.Scheme,
		Host:   c.Request.Host,
		Path:   c.Request.RequestURI + "/callback",
	}

	fmt.Println("xxxxxlllll", reqURL.String())
	authURL, err := NewDispatcher().New().Driver(provider).Redirect(
		providerData.ClientID,
		providerData.ClientSecret,
		reqURL.String(),
	)
	// providerData.RedirectURI,

	// Check for errors (usually driver not valid)
	if err != nil {
		c.String(200, err.Error())
		return
	}

	fmt.Println(authURL)
	c.Redirect(http.StatusFound, authURL)
}
