package wx_mp

import (
	"encoding/json"
	"errors"
	"github.com/go-resty/resty/v2"
)

const (
	ApiUrl = "https://api.weixin.qq.com"
)

type WxMp struct {
	AppID  string
	Secret string
	token  StableToken
}

//func (v *WxMp) Run() error {
//	url := ApiUrl + "/cgi-bin/stable_token"
//	data, _ := json.Marshal(map[string]any{
//		"grant_type":    "client_credential",
//		"appid":         v.AppID,
//		"secret":        v.Secret,
//		"force_refresh": false,
//	})
//	client := resty.New()
//	var result StableToken
//	resp, err := client.R().SetBody(string(data)).Post(url)
//	if resp.StatusCode() != 200 {
//		return err
//	}
//	if jerr := json.Unmarshal(resp.Body(), &result); jerr != nil {
//		return errors.New("请求结果解析失败")
//	}
//	v.token = result
//	return nil
//}

func (v *WxMp) Login(code string) (*LoginResult, error) {
	url := ApiUrl + "/sns/jscode2session"
	data := map[string]string{
		"appid":      v.AppID,
		"secret":     v.Secret,
		"js_code":    code,
		"grant_type": "authorization_code",
	}
	client := resty.New()
	var result LoginResult
	resp, err := client.R().SetQueryParams(data).Get(url)

	if resp.StatusCode() != 200 {
		return &result, err
	}
	if jerr := json.Unmarshal(resp.Body(), &result); jerr != nil {
		return &result, errors.New("请求结果解析失败")
	}
	if result.Openid == "" {
		return &result, errors.New("获取openid失败")
	}
	return &result, nil
}
