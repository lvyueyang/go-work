package api

import (
	"server/dal/model"
	"server/internal/types"
	"server/internal/utils"
)

// 新闻列表
type (
	NewsListReq struct {
		types.Pagination
		types.Order
		Keyword string `json:"keyword" form:"keyword"` // 查询关键字
	}
	NewsListRes struct {
		utils.ListResult[[]*model.News]
	}
)

// 新闻详情
type (
	NewsInfoReq struct {
		ID uint `json:"id"  binding:"required" label:"新闻 ID"` // 新闻 ID
	}
	NewsInfoRes struct {
		model.News
	}
)

// 新增新闻
type (
	NewsCreateReq struct {
		Title     string `json:"title" binding:"required" label:"新闻名称"`
		Desc      string `json:"desc"`
		Cover     string `json:"cover"`
		Content   string `json:"content"`
		PushDate  string `json:"push_date"`  // 发布日期 YYYY-MM-DD HH:mm:ss
		IsVisible bool   `json:"is_visible"` // 可见性
		Recommend uint   `json:"recommend"`  // 推荐等级 0 为不推荐，数值越大越靠前
	}
	NewsCreateRes struct {
		ID uint `json:"id" label:"新闻 ID" binding:"required"` // 新闻 ID
	}
)

// 修改新闻
type (
	NewsUpdateReq struct {
		ID uint `json:"id"  binding:"required" label:"新闻 ID"` // 新闻 ID
		NewsCreateReq
	}
	NewsUpdateRes struct {
		ID uint `json:"id" label:"新闻 ID"` // 新闻 ID
	}
)

// 删除用户新闻
type (
	NewsDeleteReq struct {
		ID uint `json:"id"  binding:"required" label:"新闻 ID"` // 新闻 ID
	}
	NewsDeleteRes struct {
	}
)

// 修改新闻权限码
type (
	NewsUpdateCodesReq struct {
		ID    uint     `json:"id"  binding:"required" label:"新闻 ID"` // 新闻 ID
		Codes []string `json:"codes" binding:"required" label:"权限码"` // 权限码
	}
	NewsUpdateCodesRes struct {
	}
)
