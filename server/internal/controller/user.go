package controller

import (
	"server/internal/api"
	"server/internal/middleware"
	"server/internal/service"
	"server/internal/utils"
	"server/internal/utils/resp"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *service.UserService
}

func NewUserController(e *gin.Engine) {
	c := &UserController{
		service: service.NewUserService(),
	}
	router := e.Group("/api/user")
	{
		router.GET("/current", middleware.Auth(), c.CurrentInfo)
	}

	// 管理后台使用的路由
	admin := e.Group("/api/admin/c-user")
	{
		admin.POST("/list", middleware.AdminRole(utils.CreatePermission("admin:c_user:find:list", "查询C端用户列表")), c.FindList)
		admin.POST("/update/status", middleware.AdminRole(utils.CreatePermission("admin:c_user:update:status", "修改C端用户状态")), c.UpdateState)
	}

}

// CurrentInfo
//
//	@Summary	当前登陆者信息
//	@Tags		前台-用户
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	resp.Result{data=model.User}	"用户详情"
//	@Router		/api/user/current [get]
func (c *UserController) CurrentInfo(ctx *gin.Context) {
	user := utils.GetCurrentUser(ctx)
	ctx.JSON(resp.Succ(user))
}

// FindList
//
//	@Summary	用户列表
//	@Tags		管理后台-C端用户管理
//	@Accept		json
//	@Produce	json
//	@Param		req	body		api.UserListReq	true	"body"
//	@Success	200			{object}	resp.Result{data=api.UserListRes}	"resp"
//	@Router		/api/admin/c-user/list [post]
func (c *UserController) FindList(ctx *gin.Context) {
	var body api.UserListReq
	if err := utils.BindBody(ctx, &body); err != nil {
		return
	}
	result, _ := c.service.FindList(body)
	ctx.JSON(resp.Succ(result))
}

// UpdateState
//
//	@Summary	修改用户状态
//	@Tags		管理后台-C端用户管理
//	@Accept		json
//	@Produce	json
//	@Param		req	body		api.UserUpdateStatusReq	true	"Body"
//	@Success	200	{object}	resp.Result				"resp"
//	@Router		/api/admin/c-user/update/status [post]
func (c *UserController) UpdateState(ctx *gin.Context) {
	var body api.UserUpdateStatusReq
	if err := utils.BindBody(ctx, &body); err != nil {
		return
	}

	if _, err := c.service.UpdateStatus(body.ID, body.Status); err != nil {
		ctx.JSON(resp.ParseErr(err))
		return
	}
	ctx.JSON(resp.Succ(nil))
}
