package vo

type LoginReq struct {
	Code string `form:"code"`
}

type LoginResp struct {
	Nickname   string `json:"nickname"`
	Headimgurl string `json:"headimgurl"`
}


// 	userinfo := bo.GetWechatAccessToken(logincode.Code)