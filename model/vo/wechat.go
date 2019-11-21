package vo

type LoginReq struct {
	Code  string `form:"code"`
	State string `form:"state"`
}

type LoginResp struct {
	Nickname   string `json:"nickname"`
	Headimgurl string `json:"headimgurl"`
}

// 	userinfo := bo.GetWechatAccessToken(logincode.Code)
