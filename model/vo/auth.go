package vo

// In this case we use a map to store our secrets, but you can use dotenv or your framework configuration
// for example, in revel you could use revel.Config.StringDefault(provider + "_clientID", "") etc.
var ProviderSecrets = map[string]map[string]string{
	"github": {
		"clientID":     "xxxxxxxxxxxxxx",
		"clientSecret": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		"redirectURL":  "http://localhost:9090/auth/github/callback",
	},
	"linkedin": {
		"clientID":     "xxxxxxxxxxxxxx",
		"clientSecret": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		"redirectURL":  "http://localhost:9090/auth/linkedin/callback",
	},
	"facebook": {
		"clientID":     "xxxxxxxxxxxxxx",
		"clientSecret": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		"redirectURL":  "http://localhost:9090/auth/facebook/callback",
	},
	"google": {
		"clientID":     "xxxxxxxxxxxxxx",
		"clientSecret": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		"redirectURL":  "http://localhost:9090/auth/google/callback",
	}, "bitbucket": {
		"clientID":     "xxxxxxxxxxxxxx",
		"clientSecret": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		"redirectURL":  "http://localhost:9090/auth/bitbucket/callback",
	},
	"amazon": {
		"clientID":     "xxxxxxxxxxxxxx",
		"clientSecret": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		"redirectURL":  "http://localhost:9090/auth/amazon/callback",
	},
	"slack": {
		"clientID":     "xxxxxxxxxxxxxx",
		"clientSecret": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		"redirectURL":  "http://localhost:9090/auth/slack/callback",
	},
	"asana": {
		"clientID":     "xxxxxxxxxxxxxx",
		"clientSecret": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		"redirectURL":  "http://localhost:9090/auth/asana/callback",
	},
	"wechat": {
		"clientID":     "xxxxxxxxxxxxxx",
		"clientSecret": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		"redirectURL":  "http://localhost:9090/auth/asana/callback",
	},
}
