package vo

type LoginResp struct {
	Nickname   string `json:"nickname"`
	Headimgurl string `json:"headimgurl"`
}

// 	userinfo := bo.GetWechatAccessToken(logincode.Code)
