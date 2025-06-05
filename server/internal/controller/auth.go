package controller

import (
	"server/internal/api"
	"server/internal/service"
	"server/internal/utils"
	"server/internal/utils/resp"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service *service.AuthService
}

func NewAuthController(e *gin.Engine) {
	c := &AuthController{
		service: service.NewAuthService(),
	}
	router := e.Group("/api/auth")
	{
		router.POST("/login", c.Login)
		router.POST("/wxmp/login", c.WxMpLogin)
		router.POST("/register", c.Register)
	}
}

//
//func (c *AuthController) New(e *gin.Engine) {
//	router := e.Group("/api/auth")
//
//	router.POST("/login", c.Login)
//}

// Login
//
//	@Summary	用户登录
//	@Tags		前台-用户
//	@Accept		json
//	@Produce	json
//	@Param		req	body		api.UserLoginReq					true	"body"
//	@Success	200	{object}	resp.Result{data=api.UserLoginRes}	"resp"
//	@Router		/api/auth/login [post]
func (c *AuthController) Login(ctx *gin.Context) {
	var body api.UserLoginReq
	if err := utils.BindBody(ctx, &body); err != nil {
		return
	}

	token, err := c.service.UsernameAndPasswordLogin(body.Username, body.Password)
	if err != nil {
		ctx.JSON(resp.ParseErr(err))
	} else {
		ctx.JSON(resp.Success(api.UserLoginRes{Token: token}, "登录成功"))
	}
}

// Register
//
//	@Summary	用户注册
//	@Tags		前台-用户
//	@Accept		json
//	@Produce	json
//	@Param		req	body		api.UserRegisterReq						true	"body"
//	@Success	200	{object}	resp.Result{data=api.UserRegisterRes}	"resp"
//	@Router		/api/auth/register [post]
func (c *AuthController) Register(ctx *gin.Context) {
	var body api.UserRegisterReq
	if err := utils.BindBody(ctx, &body); err != nil {
		return
	}

	token, err := c.service.UsernameAndPasswordRegister(body.Email, body.Username, body.Password)
	if err != nil {
		ctx.JSON(resp.ParseErr(err))
	} else {
		ctx.JSON(resp.Success(api.UserRegisterRes{Token: token}, "注册成功"))
	}
}

// WxMpLogin
//
//	@Summary	微信小程序登录
//	@Tags		前台-用户
//	@Accept		json
//	@Produce	json
//	@Param		req	body		api.UserLoginByWXMpReq						true	"body"
//	@Success	200	{object}	resp.Result{data=api.UserLoginByWXMpRes}	"resp"
//	@Router		/api/auth/wxmp/login [post]
func (c *AuthController) WxMpLogin(ctx *gin.Context) {
	var body api.UserLoginByWXMpReq
	if err := utils.BindBody(ctx, &body); err != nil {
		return
	}
	token, err := c.service.WxMpOpenidLogin(body.Code)
	if err != nil {
		ctx.JSON(resp.ParseErr(err))
	} else {
		ctx.JSON(resp.Success(api.UserLoginByWXMpRes{Token: token}, "登录成功"))
	}
}
