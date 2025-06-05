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

type {{.Name}}Controller struct {
	service *service.{{.Name}}Service
}

func New{{.Name}}Controller(e *gin.Engine) {
	c := &{{.Name}}Controller{
		service: service.New{{.Name}}Service(),
	}
	admin := e.Group("/api/admin/{{.DbName}}")
	{
		admin.POST("/list", middleware.AdminRole(utils.CreatePermission("admin:{{.ServiceName}}:find:list", "查询{{.CName}}列表")), c.FindList)
		admin.POST("/info", middleware.AdminRole(utils.CreatePermission("admin:{{.ServiceName}}:find:detail", "查询{{.CName}}详情")), c.FindDetail)
		admin.POST("/create", middleware.AdminRole(utils.CreatePermission("admin:{{.ServiceName}}:create", "创建{{.CName}}")), c.Create)
		admin.POST("/update/info", middleware.AdminRole(utils.CreatePermission("admin:{{.ServiceName}}:update:info", "修改{{.CName}}信息")), c.Update)
		admin.POST("/delete", middleware.AdminRole(utils.CreatePermission("admin:{{.ServiceName}}:delete", "删除{{.CName}}")), c.Delete)
	}
}

// FindList
//
//	@Summary	{{.CName}}列表
//	@Tags		管理后台-{{.CName}}
//	@Accept		json
//	@Produce	json
//	@Param		req	body		api.{{.Name}}ListReq						true	"Body"
//	@Success	200	{object}	resp.Result{data=api.{{.Name}}ListRes}	"resp"
//	@Router		/api/admin/{{.DbName}}/list [post]
func (c *{{.Name}}Controller) FindList(ctx *gin.Context) {
	var body api.{{.Name}}ListReq
	if err := utils.BindBody(ctx, &body); err != nil {
		return
	}

	result, err := c.service.FindList(body)
	if err != nil {
		ctx.JSON(resp.ParseErr(err))
		return
	}

	ctx.JSON(resp.Succ(result))
}

// FindDetail
//
//	@Summary	{{.CName}}详情
//	@Tags		管理后台-{{.CName}}
//	@Accept		json
//	@Produce	json
//	@Param		req	body		api.{{.Name}}ListReq					true	"Body"
//	@Success	200	{object}	resp.Result{data=api.{{.Name}}InfoRes}	"resp"
//	@Router		/api/admin/{{.DbName}}/info [post]
func (c *{{.Name}}Controller) FindDetail(ctx *gin.Context) {
	var body api.{{.Name}}InfoReq
	if err := utils.BindBody(ctx, &body); err != nil {
		return
	}
	info, err := c.service.FindDetail(body.ID)
	if err != nil {
		ctx.JSON(resp.ParseErr(err))
		return
	}
	ctx.JSON(resp.Succ(info))
}

// Create
//
//	@Summary	新增{{.CName}}
//	@Tags		管理后台-{{.CName}}
//	@Accept		json
//	@Produce	json
//	@Param		req	body		api.{{.Name}}CreateReq					true	"Body"
//	@Success	200	{object}	resp.Result{data=api.{{.Name}}CreateRes}	"resp"
//	@Router		/api/admin/{{.DbName}}/create [post]
func (c *{{.Name}}Controller) Create(ctx *gin.Context) {
	var body api.{{.Name}}CreateReq
	if err := utils.BindBody(ctx, &body); err != nil {
		return
	}
	user := utils.GetCurrentAdminUser(ctx)

	data := model.{{.Name}}{
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
//	@Summary	修改{{.CName}}
//	@Tags		管理后台-{{.CName}}
//	@Accept		json
//	@Produce	json
//	@Param		req	body		api.{{.Name}}UpdateReq	true	"Body"
//	@Success	200	{object}	resp.Result			"resp"
//	@Router		/api/admin/{{.DbName}}/update/info [post]
func (c *{{.Name}}Controller) Update(ctx *gin.Context) {
	var body api.{{.Name}}UpdateReq
	if err := utils.BindBody(ctx, &body); err != nil {
		return
	}

	data := model.{{.Name}}{
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
//	@Summary	删除{{.CName}}
//	@Tags		管理后台-{{.CName}}
//	@Accept		json
//	@Produce	json
//	@Param		id	path		number		true	"ID"
//	@Success	200	{object}	resp.Result	"resp"
//	@Router		/api/admin/{{.DbName}}/delete [post]
func (c *{{.Name}}Controller) Delete(ctx *gin.Context) {
	var body api.{{.Name}}DeleteReq
	if err := utils.BindBody(ctx, &body); err != nil {
		return
	}

	if err := c.service.Delete(body.ID); err != nil {
		ctx.JSON(resp.ParseErr(err))
		return
	}
	ctx.JSON(resp.Succ(nil))
}
