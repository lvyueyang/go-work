package service

import (
	"server/consts"
	"server/dal/dao"
	"server/dal/model"
	"server/lib/errs"
	"server/types"
	"server/utils"
	"strconv"
	"strings"
)

type UserService struct {
}

var userService = new(UserService)

func NewUserService() *UserService {
	return userService
}

type FindUserListOption struct {
	types.Pagination
	types.Order
	Keyword string `json:"keyword" form:"keyword"`
}

func (s *UserService) FindList(query FindUserListOption) (utils.ListResult[[]*model.User], error) {
	result := utils.ListResult[[]*model.User]{}
	n := dao.User
	q := n.Where(
		n.Name.Like("%" + query.Keyword + "%"),
	)
	if id, err := strconv.ParseUint(query.Keyword, 10, 64); err == nil {
		q = q.Or(n.ID.Eq(uint(id)))
	}

	if query.OrderKey != "" {
		col, _ := n.GetFieldByName(query.OrderKey)
		if strings.ToLower(query.OrderType) == "desc" {
			q = q.Order(col.Desc())
		} else {
			q = q.Order(col)
		}
	}

	if list, total, err := q.FindByPage(utils.PageTrans(query.Pagination)); err != nil {
		return result, err
	} else {
		result.List = list
		result.Total = total
	}

	return result, nil
}

type CreateUser struct {
	Name   string
	Age    uint32
	Email  string
	Avatar string
}

func (s *UserService) FindByID(id uint) (*model.User, error) {
	return dao.User.FindByID(id)
}

func (s *UserService) Create(u CreateUser) *model.User {
	info := &model.User{
		Name:   u.Name,
		Age:    u.Age,
		Email:  u.Email,
		Avatar: u.Avatar,
	}
	dao.User.Create(info)
	return info
}

// UseEmailCreate 使用用户名、邮箱、密码 创建
func (s *UserService) UseEmailCreate(username, email, password string) (*model.User, error) {
	info := &model.User{
		Email:    email,
		Status:   consts.UserStatusNormal,
		Username: username,
		Password: password,
	}

	if err := dao.User.Create(info); err != nil {
		return &model.User{}, err
	}
	return info, nil
}

// UseWxMpCreate 使用微信小程序 openid 创建
func (s *UserService) UseWxMpCreate(openid string) (*model.User, error) {
	info := &model.User{
		Status:   consts.UserStatusNormal,
		WxOpenId: openid,
	}

	if err := dao.User.Create(info); err != nil {
		return &model.User{}, err
	}
	return info, nil
}

// UseEmailFindOne 使用邮箱查账号
func (s *UserService) UseEmailFindOne(email string) (info *model.User, err error) {
	return dao.User.Where(dao.User.Email.Eq(email)).Take()
}

// UseUsernameFindOne 使用用户名查账号
func (s *UserService) UseUsernameFindOne(username string) (info *model.User, err error) {
	return dao.User.Where(dao.User.Username.Eq(username)).Take()
}

// UseWxMpOpenIDFindOne 使用微信 openid 查账号
func (s *UserService) UseWxMpOpenIDFindOne(openid string) (info *model.User, err error) {
	return dao.User.Where(dao.User.WxOpenId.Eq(openid)).Take()
}

// UpdateStatus 用户封禁/解封
func (s *UserService) UpdateStatus(id uint, status consts.UserStatus) (*model.User, error) {
	current, err := dao.User.FindByID(id)
	if err != nil {
		return current, errs.CreateServerError("用户不存在", err, nil)
	}
	data := &model.User{
		Status: status,
	}

	if _, err := dao.User.Where(dao.User.ID.Eq(id)).Updates(&data); err != nil {
		return current, err
	}
	return data, nil
}
