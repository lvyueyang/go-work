package api

import (
	"server/dal/model"
	"server/internal/types"
	"server/internal/utils"
)

// {{.CName}}列表
type (
	{{.Name}}ListReq struct {
		types.Pagination
		types.Order
		Keyword string `json:"keyword" form:"keyword"` // 查询关键字
	}
	{{.Name}}ListRes struct {
		utils.ListResult[[]*model.{{.Name}}]
	}
)

// {{.CName}}详情
type (
	{{.Name}}InfoReq struct {
		ID uint `json:"id"  binding:"required" label:"{{.CName}} ID"` // {{.CName}} ID
	}
	{{.Name}}InfoRes struct {
		model.{{.Name}}
	}
)

// 新增{{.CName}}
type (
	{{.Name}}CreateReq struct {
		Title     string `json:"title" binding:"required" label:"{{.CName}}名称"`
		Desc      string `json:"desc"`
		Cover     string `json:"cover"`
		Content   string `json:"content"`
		PushDate  string `json:"push_date"`  // 发布日期 YYYY-MM-DD HH:mm:ss
		IsVisible bool   `json:"is_visible"` // 可见性
		Recommend uint   `json:"recommend"`  // 推荐等级 0 为不推荐，数值越大越靠前
	}
	{{.Name}}CreateRes struct {
		ID uint `json:"id" label:"{{.CName}} ID" binding:"required"` // {{.CName}} ID
	}
)

// 修改{{.CName}}
type (
	{{.Name}}UpdateReq struct {
		ID uint `json:"id"  binding:"required" label:"{{.CName}} ID"` // {{.CName}} ID
		{{.Name}}CreateReq
	}
	{{.Name}}UpdateRes struct {
		ID uint `json:"id" label:"{{.CName}} ID"` // {{.CName}} ID
	}
)

// 删除用户{{.CName}}
type (
	{{.Name}}DeleteReq struct {
		ID uint `json:"id"  binding:"required" label:"{{.CName}} ID"` // {{.CName}} ID
	}
	{{.Name}}DeleteRes struct {
	}
)

// 修改{{.CName}}权限码
type (
	{{.Name}}UpdateCodesReq struct {
		ID    uint     `json:"id"  binding:"required" label:"{{.CName}} ID"` // {{.CName}} ID
		Codes []string `json:"codes" binding:"required" label:"权限码"` // 权限码
	}
	{{.Name}}UpdateCodesRes struct {
	}
)
