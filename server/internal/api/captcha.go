package api

import "server/internal/consts"

type (
	CaptchaSendReq struct {
		Type         consts.CaptchaType   `json:"type" binding:"required" label:"验证码类型"`          // 验证码类型， 1-手机 2-邮箱
		Value        string               `json:"value" binding:"required" label:"手机/邮箱账号"`       // 手机/邮箱账号
		Scenes       consts.CaptchaScenes `json:"scenes" binding:"required"`                      // 使用场景， 1-注册 2-忘记密码 3-修改手机 4-修改邮箱
		CaptchaKey   string               `json:"captcha_key" binding:"required"`                 // 图形验证码的key
		CaptchaValue string               `json:"captcha_value" binding:"required" label:"图形验证码"` // 输入的图形验证码
	}
	CaptchaSendRes struct {
		ID int `json:"id"`
	}
)
