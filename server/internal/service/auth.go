package service

import (
	"fmt"
	"server/config"
	"server/dal/dao"
	"server/dal/model"
	"server/internal/consts"
	"server/internal/lib/errs"
	"server/internal/lib/wx_mp"
	"server/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userService *UserService
}

var authService = &AuthService{
	userService: NewUserService(),
}

func NewAuthService() *AuthService {
	return authService
}

type LoginOptions struct {
	Username string
}

type UsernameAndPasswordRegisterOptions struct {
	Username string
	Email    string
	Code     string
	Password string
}

// UsernameAndPasswordRegister 使用用户名邮箱和密码注册
func (s *AuthService) UsernameAndPasswordRegister(email, username, password string) (string, error) {
	var opt = map[string]string{
		"email":    email,
		"username": username,
		"password": password,
	}

	// 验证用户名和邮箱是否已被使用
	if _, err := dao.User.Where(dao.User.Email.Eq(email)).Take(); err == nil {
		return "", &errs.ClientError{Msg: "邮箱已存在", Info: nil}
	}
	if _, err := dao.User.Where(dao.User.Username.Eq(username)).Take(); err == nil {
		return "", &errs.ClientError{Msg: "用户名已存在", Info: nil}
	}

	var hashPassword []byte
	// 密码加盐
	if result, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost); err != nil {
		return "", &errs.ServerError{Msg: "密码加盐失败", Err: err, Info: opt}
	} else {
		hashPassword = result
	}

	var userInfo model.User
	// 创建账号
	if result, err := s.userService.UseEmailCreate(username, email, string(hashPassword)); err != nil {
		return "", &errs.ServerError{Msg: "账号创建失败", Err: err, Info: opt}
	} else {
		userInfo = *result
	}

	token, err := utils.CreateUserToken(userInfo, consts.EmailLoginType, config.Config.Auth.TokenSecret)

	if err != nil {
		errInfo := map[string]any{"userInfo": userInfo, "opt": opt}
		return "", errs.CreateServerError("Token 生成失败", err, errInfo)
	}
	return token, nil
}

// UsernameAndPasswordLogin 使用用户名和密码登录
func (s *AuthService) UsernameAndPasswordLogin(username string, password string) (string, error) {
	info, _ := s.userService.UseUsernameFindOne(username)
	if info.ID == 0 {
		return "", &errs.ClientError{Msg: "用户名未注册", Info: nil}
	}
	err := bcrypt.CompareHashAndPassword([]byte(info.Password), []byte(password))
	if err != nil {
		return "", errs.CreateClientError("密码错误", nil)
	}
	token, err := utils.CreateUserToken(*info, consts.UsernameLoginType, config.Config.Auth.TokenSecret)
	if err != nil {
		return "", errs.CreateServerError("Token 生成失败", err, nil)
	}
	return token, nil
}

// WxMpOpenidLogin 使用微信小程序 code 登陆
func (s *AuthService) WxMpOpenidLogin(code string) (string, error) {
	fmt.Println("code: ", code)
	wxClient := wx_mp.WxMp{
		AppID:  config.Config.WxMp.AppID,
		Secret: config.Config.WxMp.AppSecret,
	}
	res, err := wxClient.Login(code)
	if err != nil {
		return "", errs.CreateServerError("请求微信登陆接口失败", err, code)
	}

	userinfo, qerr := dao.User.Where(dao.User.WxOpenId.Eq(res.Openid)).Take()
	if qerr != nil {
		// 不存在则注册一个
		user, aerr := s.userService.UseWxMpCreate(res.Openid)
		if aerr != nil {
			return "", errs.CreateServerError("注册失败", aerr, res)
		}
		userinfo = user
	}
	// openid 已注册，校验是否为正常账号

	if userinfo.Status != consts.UserStatusNormal {
		return "", errs.CreateClientError("用户状态非法，禁止登陆", *userinfo)
	}
	token, terr := utils.CreateUserToken(*userinfo, consts.WxMpLoginType, config.Config.Auth.TokenSecret)
	if terr != nil {
		return "", errs.CreateServerError("Token 生成失败", terr, nil)
	}
	return token, nil
}

func (s *AuthService) Register() {
}
