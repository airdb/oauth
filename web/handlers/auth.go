package handlers

import (
	"github.com/airdb/passport/model/bo"
	"github.com/airdb/passport/model/vo"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth(c *gin.Context) {
	c.Redirect(307, bo.GetRewriteURI())
}

// Show homepage with login URL
func IndexHandler(c *gin.Context) {
	msg := "<html><head><title>Airdb Passport</title></head><body>"
	msg += "<a href='/auth/v1/github'><button>Login with GitHub</button></a><br>"
	msg += "<a href='/auth/v1/linkedin'><button>Login with LinkedIn</button></a><br>"
	msg += "<a href='/auth/v1/facebook'><button>Login with Facebook</button></a><br>"
	msg += "<a href='/auth/v1/google'><button>Login with Google</button></a><br>"
	msg += "<a href='/auth/v1/bitbucket'><button>Login with Bitbucket</button></a><br>"
	msg += "<a href='/auth/v1/amazon'><button>Login with Amazon</button></a><br>"
	msg += "<a href='/auth/v1/slack'><button>Login with Slack</button></a><br>"
	msg += "<a href='/auth/v1/wechat'><button>Login with Wechat</button></a><br>"
	msg += "</body></html>"
	c.Writer.Write([]byte(msg))
}

// Handle callback of provider
func Callback(c *gin.Context) {
	/*
		// Retrieve query params for state and code
		state := c.Query("state")
		code := c.Query("code")
		// provider := c.Param("provider")

		// Handle callback and check for errors
		user, token, err := gocial.Handle(state, code)
		if err != nil {
			c.Writer.Write([]byte("Error: " + err.Error()))
			return
		}

		// Print in terminal user information
		fmt.Printf("%#v", token)
		fmt.Printf("%#v", user)

		// If no errors, show provider name
		c.Writer.Write([]byte("Hi, " + user.FullName))
	*/
}

func Redirect(c *gin.Context) {
	provider := c.Param("provider")

	providerScopes := map[string][]string{
		"github":    {"public_repo"},
		"linkedin":  {},
		"facebook":  {},
		"google":    {},
		"bitbucket": {},
		"amazon":    {},
		"slack":     {},
		"asana":     {},
		"wechat":    {},
	}

	providerData := vo.ProviderSecrets[provider]
	actualScopes := providerScopes[provider]

	authURL, err := NewDispatcher().New().Driver(provider).Scopes(actualScopes).Redirect(
		providerData["clientID"],
		providerData["clientSecret"],
		providerData["redirectURL"],
	)

	// Check for errors (usually driver not valid)
	if err != nil {
		c.Writer.Write([]byte("Error: " + err.Error()))
		return
	}

	// Redirect with authURL
	c.Redirect(http.StatusFound, authURL)
}
