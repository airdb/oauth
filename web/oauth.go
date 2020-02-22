package web

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/airdb/oauth/model/vo"
	"github.com/airdb/sailor"
	"github.com/airdb/sailor/enum"
	"github.com/gin-gonic/gin"
)

// Show homepage with login URL
func IndexHandler(c *gin.Context) {
	msg := "<html><head><title>Airdb Passport</title></head><body>"
	msg += "<a href='/apis/oauth/v1/github'><button>Login with GitHub</button></a><br>"
	msg += "<a href='/apis/oauth/v1/linkedin'><button>Login with LinkedIn</button></a><br>"
	msg += "<a href='/apis/oauth/v1/facebook'><button>Login with Facebook</button></a><br>"
	msg += "<a href='/apis/oauth/v1/google'><button>Login with Google</button></a><br>"
	msg += "<a href='/apis/oauth/v1/bitbucket'><button>Login with Bitbucket</button></a><br>"
	msg += "<a href='/apis/oauth/v1/amazon'><button>Login with Amazon</button></a><br>"
	msg += "<a href='/apis/oauth/v1/slack'><button>Login with Slack</button></a><br>"
	msg += "<a href='/apis/oauth/v1/wechat'><button>Login with Wechat</button></a><br>"
	msg += "</body></html>"

	_, err := c.Writer.Write([]byte(msg))
	if err != nil {
		log.Println(err)
	}
}

func Login(c *gin.Context) {
	provider := c.Param("provider")

	authURL, err := vo.GetAuthRedirectURL(provider)
	if err != nil {
		c.Redirect(http.StatusFound, "/")
	}

	c.Redirect(http.StatusFound, *authURL)
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
		c.JSON(http.StatusOK, sailor.HTTPAirdbResponse{
			Code:    enum.AirdbSuccess,
			Success: true,
			Data: vo.LoginResp{
				Nickname:   "xxx",
				Headimgurl: "xxx.png",
			},
		})

		return
	}

	c.JSON(http.StatusOK, sailor.HTTPAirdbResponse{
		Code:    enum.AirdbSuccess,
		Success: true,
		Data: vo.LoginResp{
			Nickname:   userInfo.Login,
			Headimgurl: userInfo.AvatarURL,
		},
	})
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Create the JWT key used to create the signature
var jwtKey = []byte("my_secret_key")

func Token(c *gin.Context) {
	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: "xxx",
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return
	}
	log.Println(tokenString)
	c.String(http.StatusOK, tokenString)
}
