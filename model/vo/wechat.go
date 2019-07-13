package vo

type LoginResp struct {
	Nickname string
	Pic      string
}

type LoginReq struct {
	Code string `form:"code"`
}
