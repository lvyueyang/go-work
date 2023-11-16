package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"server/consts/permission"
	"server/dal/model"
	"server/lib/valid"
	"server/middleware"
	"server/modules/service"
	"server/utils"
	"server/utils/resp"
	"strconv"
	"time"
)

type {{.Name}}Controller struct {
	service *service.{{.Name}}Service
}

func New{{.Name}}Controller(e *gin.Engine) {
	c := &{{.Name}}Controller{
		service: service.New{{.Name}}Service(),
	}
	admin := e.Group("/api/admin/{{.DbName}}")
	admin.POST("/list", middleware.AdminRole(permission.Admin{{.Name}}Find), c.FindList)
	admin.GET("/:id", middleware.AdminRole(permission.Admin{{.Name}}FindDetail), c.FindDetail)
	admin.POST("", middleware.AdminRole(permission.Admin{{.Name}}Create), c.Create)
	admin.PUT("", middleware.AdminRole(permission.Admin{{.Name}}UpdateInfo), c.Update)
	admin.DELETE("/:id", middleware.AdminRole(permission.Admin{{.Name}}Delete), c.Delete)
}

// FindList
//
//	@Summary	{{.CName}}列表
//	@Tags		管理后台-{{.CName}}
//	@Accept		json
//	@Produce	json
//	@Param		req	body		service.Find{{.Name}}ListOption						true	"Body"
//	@Success	200	{object}	resp.Result{data=resp.RList{list=[]model.{{.Name}}}}	"resp"
//	@Router		/api/admin/{{.DbName}}/list [post]
func (c *{{.Name}}Controller) FindList(ctx *gin.Context) {
	query := service.Find{{.Name}}ListOption{}
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(resp.ParamErr(valid.ErrTransform(err)))
		return
	}
	result, _ := c.service.FindList(query)
	ctx.JSON(resp.Succ(result))
}

// FindDetail
//
//	@Summary	{{.CName}}详情
//	@Tags		管理后台-{{.CName}}
//	@Accept		json
//	@Produce	json
//	@Param		id	path		number							true	"ID"
//	@Success	200	{object}	resp.Result{data=model.{{.Name}}}	"resp"
//	@Router		/api/admin/{{.DbName}}/{id} [get]
func (c *{{.Name}}Controller) FindDetail(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	info, err := c.service.FindDetail(uint(id))
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
//	@Param		req	body		Create{{.Name}}BodyDto	true	"Body"
//	@Success	200	{object}	resp.Result			"resp"
//	@Router		/api/admin/{{.DbName}} [post]
func (c *{{.Name}}Controller) Create(ctx *gin.Context) {
	var body Create{{.Name}}BodyDto
	if err := ctx.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		ctx.JSON(resp.ParamErr(valid.ErrTransform(err)))
		return
	}
	user := utils.GetCurrentAdminUser(ctx)

	data := model.{{.Name}}{
		Title:     body.Title,
		Desc:      body.Desc,
		Cover:     body.Cover,
		Content:   body.Content,
		Recommend: body.Recommend,
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
//	@Param		req	body		Update{{.Name}}BodyDto	true	"Body"
//	@Success	200	{object}	resp.Result			"resp"
//	@Router		/api/admin/{{.DbName}} [put]
func (c *{{.Name}}Controller) Update(ctx *gin.Context) {
	var body Update{{.Name}}BodyDto
	if err := ctx.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		ctx.JSON(resp.ParamErr(valid.ErrTransform(err)))
		return
	}

	data := model.{{.Name}}{
		Title:     body.Title,
		Desc:      body.Desc,
		Cover:     body.Cover,
		Content:   body.Content,
		Recommend: body.Recommend,
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
//	@Router		/api/admin/{{.DbName}}/{id} [delete]
func (c *{{.Name}}Controller) Delete(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err := c.service.Delete(uint(id)); err != nil {
		ctx.JSON(resp.ParseErr(err))
		return
	}
	ctx.JSON(resp.Succ(nil))
}

type Create{{.Name}}BodyDto struct {
	Title     string `json:"title" binding:"required" label:"{{.CName}}名称"`
	Desc      string `json:"desc"`
	Cover     string `json:"cover"`
	Content   string `json:"content"`
	PushDate  string `json:"push_date"` // 发布日期 YYYY-MM-DD HH:mm:ss
	Recommend uint   `json:"recommend"` // 推荐等级 0 为不推荐，数值越大越靠前
}

type Update{{.Name}}BodyDto struct {
	ID uint `json:"id"  binding:"required" label:"角色 ID"` // 角色 ID
	Create{{.Name}}BodyDto
}
