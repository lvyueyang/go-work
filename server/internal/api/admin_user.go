package api

import (
	"server/dal/model"
	"server/internal/consts"
	"server/internal/types"
	"server/internal/utils"
)

// 管理员列表
type (
	AdminUserListReq struct {
		types.Pagination
		types.Order
		Keyword string `json:"keyword" form:"keyword"` // 查询关键字
	}
	AdminUserListRes struct {
		utils.ListResult[[]*model.AdminUser]
	}
)

// 管理员新增
type (
	AdminUserCreateReq struct {
		Name     string `json:"name" binding:"required" label:"姓名"`      // 姓名
		Username string `json:"username" binding:"required" label:"用户名"` // 用户名
		Password string `json:"password" binding:"required" label:"密码"`  // 密码
		Email    string `json:"email" binding:"required" label:"邮箱"`     // 邮箱
		Avatar   string `json:"avatar"`                                  // 头像
	}
	AdminUserCreateRes struct {
		ID string `json:"id"` // 管理员ID
	}
)

// 管理员信息修改
type (
	AdminUserUpdateReq struct {
		ID     uint   `json:"id" binding:"required" label:"ID"`   // ID
		Name   string `json:"name" binding:"required" label:"姓名"` // 姓名
		Avatar string `json:"avatar"`                             // 头像
	}
	AdminUserUpdateRes struct {
	}
)

// 管理员状态修改
type (
	AdminUserUpdateStatusReq struct {
		ID     uint                   `json:"id" binding:"required" label:"ID"`                    // ID
		Status consts.AdminUserStatus `json:"status" binding:"required" label:"用户状态" enums:"-1,1"` // 状态 1-解封 2-封禁
	}
	AdminUserUpdateStatusRes struct {
	}
)

// 管理员角色修改
type (
	AdminUserUpdateRoleReq struct {
		UserId  uint   `json:"user_id" binding:"required" label:"用户 ID"`  // 用户 ID
		RoleIds []uint `json:"role_ids" binding:"required" label:"角色 ID"` // 角色 ID
	}
	AdminUserUpdateRoleRes struct {
	}
)

// 管理员删除
type (
	AdminUserDeleteReq struct {
		ID uint `json:"id" binding:"required" label:"ID"` // ID
	}
	AdminUserDeleteRes struct {
	}
)

// 管理员重置密码
type (
	AdminUserResetPasswordReq struct {
		ID       uint   `json:"id" binding:"required" label:"ID"`       // ID
		Password string `json:"password" binding:"required" label:"密码"` // 密码
	}
	AdminUserResetPasswordRes struct {
	}
)

// 管理员重置密码
type (
	AdminUserInfoReq struct {
	}
	AdminUserInfoRes struct {
		model.AdminUser
	}
)
