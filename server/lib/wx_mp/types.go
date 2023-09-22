package wx_mp

type StableToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
}

type LoginResult struct {
	SessionKey string `json:"session_key"`
	Openid     string `json:"openid"`
	Unionid    string `json:"unionid"`
	Errmsg     string `json:"errmsg"`
	Errcode    string `json:"errcode"`
}
