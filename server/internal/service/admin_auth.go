package service

import (
	"fmt"
	"server/config"
	"server/dal/dao"
	"server/internal/lib/errs"
	"server/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

type AdminAuthService struct {
}

var adminAuthService = new(AdminAuthService)

func NewAdminAuthService() *AdminAuthService {
	return adminAuthService
}

func (s *AdminAuthService) UsernameAndPasswordLogin(username string, password string) (string, error) {
	fmt.Println("username", username)
	info, err := dao.AdminUser.Where(dao.AdminUser.Username.Eq(username)).Take()
	if err != nil {
		return "", &errs.ClientError{Msg: "用户未注册", Info: nil}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(info.Password), []byte(password)); err != nil {
		return "", errs.CreateClientError("密码错误", nil)
	}
	token, err := utils.CreateAdminUserToken(info.ID, config.Config.Auth.AdminTokenSecret)
	if err != nil {
		return "", err
	}
	return token, nil
}
