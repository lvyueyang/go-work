package api

// 用户登录
type (
	AdminUserLoginReq struct {
		Username string `json:"username"` // 用户名
		Password string `json:"password"` // 密码
	}
	AdminUserLoginRes struct {
		Token string `json:"token"`
	}
)

// 忘记密码
type (
	AdminUserForgetPasswordReq struct {
		Email    string `json:"email" binding:"required" label:"邮箱"`    // 邮箱
		Password string `json:"password" binding:"required" label:"密码"` // 密码
		Captcha  string `json:"captcha" binding:"required" label:"验证码"` // 邮箱验证码
	}
	AdminUserForgetPasswordRes struct {
	}
)

// 根用户初始化
type (
	AdminInitRootUserReq struct {
		Name     string `json:"name" binding:"required" label:"用户名"`     // 昵称
		Username string `json:"username" binding:"required" label:"用户名"` // 用户名
		Password string `json:"password" binding:"required" label:"密码"`  // 密码
		Email    string `json:"email" binding:"required" label:"邮箱"`     // 邮箱
	}
	AdminInitRootUserRes struct {
		Token string `json:"token"`
	}
)
