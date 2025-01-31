package consts

type LoginType uint // 登陆类型

const (
	UsernameLoginType LoginType = 1 // 使用用户名和密码登录
	EmailLoginType    LoginType = 2 // 邮箱
	WxMpLoginType     LoginType = 3 // 微信小程序
)
