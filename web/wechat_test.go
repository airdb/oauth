package web

import (
	"net/http"
	"testing"
)

func TestWechatLogin(t *testing.T) {
	uri := "/apis/oauth/v1/wechat/login?code=aaa111"

	// input, _ := json.Marshal(&vo.WechatLoginReq{Code: "xxx111"})
	resp := APIRequest(uri, http.MethodGet, nil)

	if resp.Code != http.StatusOK {
		t.Error(uri, resp.Code)
	}

	t.Log(uri, resp.Code)
	t.Log("resp:", resp.Body)
}
