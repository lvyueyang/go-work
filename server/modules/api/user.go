package api

import (
	"github.com/gin-gonic/gin"
	"server/middleware"
	"server/modules/service"
	"server/utils"
	"server/utils/resp"
)

type UserController struct {
	service *service.UserService
}

func NewUserController(e *gin.Engine) {
	c := &UserController{
		service: service.NewUserService(),
	}
	admin := e.Group("/api/user")
	admin.GET("/current", middleware.Auth(), c.CurrentInfo)
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

type CreateUserBodyDto struct {
	Name string `json:"name" binding:"required" label:"姓名"` // 姓名
	Sex  string `json:"sex" binding:"required" label:"性别"`  // 性别
}
