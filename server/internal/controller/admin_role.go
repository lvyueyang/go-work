package controller

import (
	"server/internal/api"
	"server/internal/lib/valid"
	"server/internal/middleware"
	"server/internal/service"
	"server/internal/utils"
	"server/internal/utils/resp"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type AdminRoleController struct {
	service          *service.AdminRoleService
	adminUserService *service.AdminUserService
}

func NewAdminRoleController(e *gin.Engine) {
	c := &AdminRoleController{
		service:          service.NewAdminRoleService(),
		adminUserService: service.NewAdminUserService(),
	}
	group := e.Group("/api/admin/role")
	group.POST("/list", middleware.AdminRole(utils.CreatePermission("admin:role:find:list", "查询管理员角色列表")), c.FindList)
	group.POST("/create", middleware.AdminRole(utils.CreatePermission("admin:role:create", "创建管理员角色")), c.Create)
	group.POST("/update", middleware.AdminRole(utils.CreatePermission("admin:role:update:info", "修改管理员角色信息")), c.Update)
	group.POST("/delete", middleware.AdminRole(utils.CreatePermission("admin:role:delete", "删除管理员角色")), c.Delete)
	group.POST("/update/permission-codes", middleware.AdminRole(utils.CreatePermission("admin:role:update:code", "修改管理角色权限码")), c.UpdatePermissionCodes)
	group.GET("/permission/codes", c.FindPermissionCodes)
}

// FindList
//
//	@Summary	管理员角色列表
//	@Tags		管理后台-管理员角色
//	@Accept		json
//	@Produce	json
//	@Param		req	body		api.AdminRoleListReq					true	"body"
//	@Success	200	{object}	resp.Result{data=api.AdminRoleListRes}	"resp"
//	@Router		/api/admin/role/list [post]
func (c *AdminRoleController) FindList(ctx *gin.Context) {
	var body api.AdminRoleListReq
	utils.BindBody(ctx, &body)

	result, err := c.service.FindList(body)
	if err != nil {
		ctx.JSON(resp.ParseErr(err))
		return
	}
	ctx.JSON(resp.Succ(result))
}

// Create
//
//	@Summary	新增管理员角色
//	@Tags		管理后台-管理员角色
//	@Accept		json
//	@Produce	json
//	@Param		req	body		api.AdminRoleCreateReq						true	"body"
//	@Success	200	{object}	resp.Result{data=api.AdminRoleCreateRes}	"resp"
//	@Router		/api/admin/role/create [post]
func (c *AdminRoleController) Create(ctx *gin.Context) {
	var body api.AdminRoleCreateReq
	utils.BindBody(ctx, &body)

	info, err := c.service.Create(body)
	if err != nil {
		ctx.JSON(resp.ParseErr(err))
		return
	}
	ctx.JSON(resp.Succ(info.ID))
}

// Update
//
//	@Summary	修改管理员角色
//	@Tags		管理后台-管理员角色
//	@Accept		json
//	@Produce	json
//	@Param		req	body		api.AdminRoleUpdateReq	true	"body"
//	@Success	200	{object}	resp.Result				"resp"
//	@Router		/api/admin/role/update [post]
func (c *AdminRoleController) Update(ctx *gin.Context) {
	var body api.AdminRoleUpdateReq
	utils.BindBody(ctx, &body)
	if _, err := c.service.Update(body); err != nil {
		ctx.JSON(resp.ParseErr(err))
		return
	}
	ctx.JSON(resp.Succ(nil))
}

// Delete
//
//	@Summary	删除角色
//	@Tags		管理后台-管理员角色
//	@Accept		json
//	@Produce	json
//	@Param		req	body		api.AdminRoleDeleteReq	true	"body"
//	@Success	200	{object}	resp.Result				"resp"
//	@Router		/api/admin/role/delete [post]
func (c *AdminRoleController) Delete(ctx *gin.Context) {
	var body api.AdminRoleDeleteReq
	utils.BindBody(ctx, &body)

	if err := c.service.Delete(body.ID); err != nil {
		ctx.JSON(resp.ParseErr(err))
		return
	}
	ctx.JSON(resp.Succ(nil))
}

// UpdatePermissionCodes
//
//	@Summary	修改管理员角色权限
//	@Tags		管理后台-管理员角色
//	@Accept		json
//	@Produce	json
//	@Param		req	body		api.AdminRoleUpdateCodesReq	true	"body"
//	@Success	200	{object}	resp.Result					"resp"
//	@Router		/api/admin/role/update/permission-codes [post]
func (c *AdminRoleController) UpdatePermissionCodes(ctx *gin.Context) {
	var body api.AdminRoleUpdateCodesReq
	if err := ctx.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		ctx.JSON(resp.ParamErr(valid.ErrTransform(err)))
		return
	}

	for _, code := range body.Codes {
		if utils.FindPermission(code).Code == "" {
			ctx.JSON(resp.ParamErr("无效的权限码: " + code))
			return
		}
	}

	if _, err := c.service.UpdatePermissionCode(body.ID, body.Codes); err != nil {
		ctx.JSON(resp.ParseErr(err))
		return
	}
	ctx.JSON(resp.Succ(nil))
}

// FindPermissionCodes
//
//	@Summary	管理后台权限码列表
//	@Tags		管理后台-管理员角色
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	resp.Result{data=api.AdminPermissionCodeListRes}	"resp"
//	@Router		/api/admin/role/permission/codes [get]
func (c *AdminRoleController) FindPermissionCodes(ctx *gin.Context) {
	ctx.JSON(resp.Succ(utils.PermissionList))
}
