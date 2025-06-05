package api

// 前台用户登录
type (
	UserLoginReq struct {
		Username string `json:"username" binding:"required"` // 用户名
		Password string `json:"password" binding:"required"` // 密码
	}
	UserLoginRes struct {
		Token string `json:"token" binding:"required"`
	}
)

// 前台用户小程序登录
type (
	UserLoginByWXMpReq struct {
		Code string `json:"code" binding:"required"` // code
	}
	UserLoginByWXMpRes struct {
		Token string `json:"token" binding:"required"`
	}
)

// 前台用户注册
type (
	UserRegisterReq struct {
		Username string `json:"username" binding:"required" label:"用户名"`  // 用户名
		Password string `json:"password" binding:"required" label:"密码"`   // 密码
		Email    string `json:"email" binding:"required" label:"邮箱"`      // 邮箱
		Captcha  string `json:"captcha" binding:"required" label:"邮箱验证码"` // 邮箱验证码
	}
	UserRegisterRes struct {
		Token string `json:"token" binding:"required"`
	}
)
