package controller

import (
	"server/dal/model"
	"server/internal/api"
	"server/internal/middleware"
	"server/internal/service"
	"server/internal/utils"
	"server/internal/utils/resp"
	"time"

	"github.com/gin-gonic/gin"
)

type NewsController struct {
	service *service.NewsService
}

func NewNewsController(e *gin.Engine) {
	c := &NewsController{
		service: service.NewNewsService(),
	}
	admin := e.Group("/api/admin/news")
	{
		admin.POST("/list", middleware.AdminRole(utils.CreatePermission("admin:news:find:list", "查询新闻列表")), c.FindList)
		admin.POST("/info", middleware.AdminRole(utils.CreatePermission("admin:news:find:detail", "查询新闻详情")), c.FindDetail)
		admin.POST("/create", middleware.AdminRole(utils.CreatePermission("admin:news:create", "创建新闻")), c.Create)
		admin.POST("/update/info", middleware.AdminRole(utils.CreatePermission("admin:news:update:info", "修改新闻信息")), c.Update)
		admin.POST("/delete", middleware.AdminRole(utils.CreatePermission("admin:news:delete", "删除新闻")), c.Delete)
	}
}

// FindList
//
//	@Summary	新闻列表
//	@Tags		管理后台-新闻
//	@Accept		json
//	@Produce	json
//	@Param		req	body		api.NewsListReq						true	"Body"
//	@Success	200	{object}	resp.Result{data=api.NewsListRes}	"resp"
//	@Router		/api/admin/news/list [post]
func (c *NewsController) FindList(ctx *gin.Context) {
	var body api.NewsListReq
	utils.BindBody(ctx, &body)

	result, err := c.service.FindList(body)
	if err != nil {
		ctx.JSON(resp.ParseErr(err))
		return
	}

	ctx.JSON(resp.Succ(result))
}

// FindDetail
//
//	@Summary	新闻详情
//	@Tags		管理后台-新闻
//	@Accept		json
//	@Produce	json
//	@Param		req	body		api.NewsListReq					true	"Body"
//	@Success	200	{object}	resp.Result{data=api.NewsInfoRes}	"resp"
//	@Router		/api/admin/news/info [post]
func (c *NewsController) FindDetail(ctx *gin.Context) {
	var body api.NewsInfoReq
	utils.BindBody(ctx, &body)
	info, err := c.service.FindDetail(body.ID)
	if err != nil {
		ctx.JSON(resp.ParseErr(err))
		return
	}
	ctx.JSON(resp.Succ(info))
}

// Create
//
//	@Summary	新增新闻
//	@Tags		管理后台-新闻
//	@Accept		json
//	@Produce	json
//	@Param		req	body		api.NewsCreateReq					true	"Body"
//	@Success	200	{object}	resp.Result{data=api.NewsCreateRes}	"resp"
//	@Router		/api/admin/news/create [post]
func (c *NewsController) Create(ctx *gin.Context) {
	var body api.NewsCreateReq
	utils.BindBody(ctx, &body)
	user := utils.GetCurrentAdminUser(ctx)

	data := model.News{
		Title:     body.Title,
		Desc:      body.Desc,
		Cover:     body.Cover,
		Content:   body.Content,
		Recommend: body.Recommend,
		IsVisible: body.IsVisible,
		PushDate:  time.Now(),
		AuthorID:  user.ID,
	}

	if body.PushDate != "" {
		t, terr := time.Parse("2006-01-02 15:04:05", body.PushDate)
		if body.PushDate != "" && terr != nil {
			ctx.JSON(resp.ParamErr("发布日期格式错误"))
			return
		}
		data.PushDate = t
	}

	if _, err := c.service.Create(data); err != nil {
		ctx.JSON(resp.ParseErr(err))
		return
	}
	ctx.JSON(resp.Succ(nil))
}

// Update
//
//	@Summary	修改新闻
//	@Tags		管理后台-新闻
//	@Accept		json
//	@Produce	json
//	@Param		req	body		api.NewsUpdateReq	true	"Body"
//	@Success	200	{object}	resp.Result			"resp"
//	@Router		/api/admin/news/update/info [post]
func (c *NewsController) Update(ctx *gin.Context) {
	var body api.NewsUpdateReq
	utils.BindBody(ctx, &body)

	data := model.News{
		Title:     body.Title,
		Desc:      body.Desc,
		Cover:     body.Cover,
		Content:   body.Content,
		Recommend: body.Recommend,
		IsVisible: body.IsVisible,
		PushDate:  time.Now(),
	}

	if body.PushDate != "" {
		t, terr := time.Parse("2006-01-02 15:04:05", body.PushDate)
		if body.PushDate != "" && terr != nil {
			ctx.JSON(resp.ParamErr("发布日期格式错误"))
			return
		}
		data.PushDate = t
	}
	if _, err := c.service.Update(body.ID, data); err != nil {
		ctx.JSON(resp.ParseErr(err))
		return
	}
	ctx.JSON(resp.Succ(nil))
}

// Delete
//
//	@Summary	删除新闻
//	@Tags		管理后台-新闻
//	@Accept		json
//	@Produce	json
//	@Param		id	path		number		true	"ID"
//	@Success	200	{object}	resp.Result	"resp"
//	@Router		/api/admin/news/delete [post]
func (c *NewsController) Delete(ctx *gin.Context) {
	var body api.NewsDeleteReq
	utils.BindBody(ctx, &body)

	if err := c.service.Delete(body.ID); err != nil {
		ctx.JSON(resp.ParseErr(err))
		return
	}
	ctx.JSON(resp.Succ(nil))
}
