package controller

import (
	"server/internal/consts"
	"server/internal/lib/valid"
	"server/internal/middleware"
	"server/internal/service"
	"server/internal/utils"
	"server/internal/utils/resp"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type UserController struct {
	service *service.UserService
}

func NewUserController(e *gin.Engine) {
	c := &UserController{
		service: service.NewUserService(),
	}
	router := e.Group("/api/user")
	router.GET("/current", middleware.Auth(), c.CurrentInfo)

	// 管理后台使用的路由
	admin := e.Group("/api/admin/c-user")
	admin.GET("", middleware.AdminRole(utils.CreatePermission("admin:c_user:find:list", "查询C端用户列表")), c.FindList)
	admin.PUT("/status", middleware.AdminRole(utils.CreatePermission("admin:c_user:update:status", "修改C端用户状态")), c.UpdateState)

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
//	@Param		current		query		number											false	"当前页"	default(1)
//	@Param		page_size	query		number											false	"每页条数"	default(20)
//	@Param		order_key	query		string											false	"需要排序的列"
//	@Param		order_type	query		string											false	"排序方式"	Enums(ase,desc)
//	@Param		keyword		query		string											false	"按名称或ID搜索"
//	@Success	200			{object}	resp.Result{data=resp.RList{list=[]model.User}}	"resp"
//	@Router		/api/admin/c-user [get]
func (c *UserController) FindList(ctx *gin.Context) {
	query := service.FindUserListOption{}
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(resp.ParamErr(valid.ErrTransform(err)))
		return
	}
	result, _ := c.service.FindList(query)
	ctx.JSON(resp.Succ(result))
}

// UpdateState
//
//	@Summary	修改用户状态
//	@Tags		管理后台-C端用户管理
//	@Accept		json
//	@Produce	json
//	@Param		req	body		UpdateUserStatusBodyDto	true	"Body"
//	@Success	200	{object}	resp.Result				"resp"
//	@Router		/api/admin/status [put]
func (c *UserController) UpdateState(ctx *gin.Context) {
	var body UpdateUserStatusBodyDto
	if err := ctx.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		ctx.JSON(resp.ParamErr(valid.ErrTransform(err)))
		return
	}

	if _, err := c.service.UpdateStatus(body.ID, body.Status); err != nil {
		ctx.JSON(resp.ParseErr(err))
		return
	}
	ctx.JSON(resp.Succ(nil))
}

type CreateUserBodyDto struct {
	Name string `json:"name" binding:"required" label:"姓名"` // 姓名
	Sex  string `json:"sex" binding:"required" label:"性别"`  // 性别
}

type UpdateUserStatusBodyDto struct {
	ID     uint              `json:"id" binding:"required"`                  // 用户 ID
	Status consts.UserStatus `json:"status" binding:"required" enums:"-1,1"` // 状态 -1封禁 1-正常
}
