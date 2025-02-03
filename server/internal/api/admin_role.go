package api

import (
	"server/dal/model"
	"server/internal/types"
	"server/internal/utils"
)

// 角色列表
type (
	AdminRoleListReq struct {
		types.Pagination
		types.Order
		Keyword string `json:"keyword" form:"keyword"` // 查询关键字
	}
	AdminRoleListRes struct {
		utils.ListResult[[]*model.AdminRole]
	}
)

// 新增角色
type (
	AdminRoleCreateReq struct {
		Name string `json:"name" binding:"required" label:"角色名称"` // 角色名称
		Code string `json:"code" binding:"required" label:"角色编码"` // 角色编码
		Desc string `json:"desc"`                                 // 描述
	}
	AdminRoleCreateRes struct {
		ID uint `json:"id" label:"角色 ID"` // 角色 ID
	}
)

// 修改角色
type (
	AdminRoleUpdateReq struct {
		ID uint `json:"id"  binding:"required" label:"角色 ID" required:"true"` // 角色 ID
		AdminRoleCreateReq
	}
	AdminRoleUpdateRes struct {
		ID uint `json:"id" label:"角色 ID"` // 角色 ID
	}
)

// 删除用户角色
type (
	AdminRoleDeleteReq struct {
		ID uint `json:"id"  binding:"required" label:"角色 ID"` // 角色 ID
	}
	AdminRoleDeleteRes struct {
	}
)

// 修改角色权限码
type (
	AdminRoleUpdateCodesReq struct {
		ID    uint     `json:"id"  binding:"required" label:"角色 ID"` // 角色 ID
		Codes []string `json:"codes" binding:"required" label:"权限码"` // 权限码
	}
	AdminRoleUpdateCodesRes struct {
	}
)

// 权限码列表
type (
	AdminPermissionCodeListReq struct {
	}
	AdminPermissionCodeListRes = []utils.PermissionInfo
)
