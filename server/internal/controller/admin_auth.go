package controller

import (
	"server/internal/api"
	"server/internal/consts"
	"server/internal/service"
	"server/internal/utils"
	"server/internal/utils/resp"

	"github.com/gin-gonic/gin"
)

type AdminAuthController struct {
	service          *service.AdminAuthService
	adminUserService *service.AdminUserService
	captchaService   *service.CaptchaService
}

func NewAdminAuthController(e *gin.Engine) {
	c := &AdminAuthController{
		service:          service.NewAdminAuthService(),
		adminUserService: service.NewAdminUserService(),
		captchaService:   service.NewCaptchaService(),
	}
	group := e.Group("/api/admin/auth")
	{
		group.POST("/login", c.Login)
		group.POST("/init-root-user", c.InitRootUser)
		group.POST("/forget-password", c.ForgetPassword)
	}
}

// Login
//
//	@Summary	用户登录
//	@Tags		管理后台-用户认证
//	@Accept		json
//	@Produce	json
//	@Param		req	body		api.AdminUserLoginReq					true	"body"
//	@Success	200	{object}	resp.Result{data=api.AdminUserLoginRes}	"resp"
//	@Router		/api/admin/auth/login [post]
func (c *AdminAuthController) Login(ctx *gin.Context) {
	var body api.AdminUserLoginReq
	utils.BindBody(ctx, &body)

	token, err := c.service.UsernameAndPasswordLogin(body.Username, body.Password)
	if err != nil {
		ctx.JSON(resp.ParseErr(err))
	} else {
		ctx.JSON(resp.Success(api.AdminUserLoginRes{Token: token}, "登录成功"))
	}
}

// InitRootUser
//
//	@Summary	初始化超级管理员用户
//	@Tags		管理后台-用户认证
//	@Accept		json
//	@Produce	json
//	@Param		req	body		api.AdminInitRootUserReq					true	"body"
//	@Success	200	{object}	resp.Result{data=api.AdminInitRootUserRes}	"resp"
//	@Router		/api/admin/auth/init-root-user [post]
func (c *AdminAuthController) InitRootUser(ctx *gin.Context) {
	var body api.AdminInitRootUserReq
	utils.BindBody(ctx, &body)

	token, err := c.adminUserService.CreateRootUser(body.Username, body.Name, body.Password, body.Email)
	if err != nil {
		ctx.JSON(resp.ParseErr(err))
	} else {
		ctx.JSON(resp.Success(api.AdminInitRootUserRes{Token: token}, "超级管理员创建成功"))
	}
}

// ForgetPassword
//
//	@Summary	忘记密码
//	@Tags		管理后台-用户认证
//	@Accept		json
//	@Produce	json
//	@Param		req	body		api.AdminUserForgetPasswordReq	true	"body"
//	@Success	200	{object}	resp.Result						"resp"
//	@Router		/api/admin/auth/forget-password [post]
func (c *AdminAuthController) ForgetPassword(ctx *gin.Context) {
	var body api.AdminUserForgetPasswordReq
	utils.BindBody(ctx, &body)

	if ok, err := c.captchaService.Validate(body.Email, consts.CaptchaTypeEmail, body.Captcha, consts.CaptchaScenesForgetPassword); ok == false {
		ctx.JSON(resp.ParamErr(err.Error()))
		return
	}

	if err := c.adminUserService.ResetPassword(body.Email, body.Password); err != nil {
		ctx.JSON(resp.ParseErr(err))
	} else {
		ctx.JSON(resp.Success(nil, "密码重置成功"))
	}
}
