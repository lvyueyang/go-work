package api

import (
	"server/dal/model"
	"server/internal/consts"
	"server/internal/types"
	"server/internal/utils"
)

// 用户列表
type (
	UserListReq struct {
		types.Pagination
		types.Order
		Keyword string `json:"keyword" form:"keyword"` // 查询关键字
	}
	UserListRes struct {
		utils.ListResult[[]*model.User]
	}
)

// 修改用户状态
type (
	UserUpdateStatusReq struct {
		ID     uint              `json:"id" binding:"required"`                  // 用户 ID
		Status consts.UserStatus `json:"status" binding:"required" enums:"-1,1"` // 状态 -1封禁 1-正常
	}
	UserUpdateStatusRes struct {
	}
)
