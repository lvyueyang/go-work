package controller

import (
	"errors"
	"fmt"
	"io"
	"path"
	"server/config"
	"server/dal/model"
	"server/internal/api"
	"server/internal/consts"
	"server/internal/lib/valid"
	"server/internal/middleware"
	"server/internal/service"
	"server/internal/utils"
	"server/internal/utils/resp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type AdminUserController struct {
	service *service.AdminUserService
}

func NewAdminUserController(e *gin.Engine) {
	c := &AdminUserController{
		service: service.NewAdminUserService(),
	}
	admin := e.Group("/api/admin/user")

	admin.GET("/current", middleware.AdminAuth(), c.CurrentInfo)
	admin.POST("/list", middleware.AdminRole(utils.CreatePermission("admin:user:find:list", "查询管理员列表")), c.FindList)
	admin.POST("/create", middleware.AdminRole(utils.CreatePermission("admin:user:create", "创建管理员")), c.Create)
	admin.POST("/update/info", middleware.AdminRole(utils.CreatePermission("admin:user:update:info", "修改管理员基本信息")), c.Update)
	admin.POST("/update/status", middleware.AdminRole(utils.CreatePermission("admin:user:update:status", "修改管理员状态")), c.UpdateStatus)
	admin.POST("/update/role", middleware.AdminRole(utils.CreatePermission("admin:user:update:role", "修改管理员角色")), c.UpdateRole)
	admin.POST("/reset-password", middleware.AdminRole(utils.CreatePermission("admin:user:update:password", "修改管理员密码")), c.ResetPassword)
	admin.POST("/delete", middleware.AdminRole(utils.CreatePermission("admin:user:delete", "删除管理员")), c.Delete)
	admin.POST("/upload", middleware.AdminRole(utils.CreatePermission("admin:user:upload:file", "上传文件到本地")), c.Upload)
}

// FindList
//
//	@Summary	管理员列表
//	@Tags		管理后台-管理员用户
//	@Accept		json
//	@Produce	json
//	@Param		req	body		api.AdminUserListReq					true	"body"
//	@Success	200	{object}	resp.Result{data=api.AdminUserListRes}	"resp"
//	@Router		/api/admin/user/list [post]
func (c *AdminUserController) FindList(ctx *gin.Context) {
	var body api.AdminUserListReq
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
//	@Summary	新增管理员
//	@Tags		管理后台-管理员用户
//	@Accept		json
//	@Produce	json
//	@Param		req	body		api.AdminUserCreateReq						true	"body"
//	@Success	200	{object}	resp.Result{data=api.AdminUserCreateRes}	"resp"
//	@Router		/api/admin/user/create [post]
func (c *AdminUserController) Create(ctx *gin.Context) {
	var body api.AdminUserCreateReq
	utils.BindBody(ctx, &body)
	info, err := c.service.Create(model.AdminUser{
		Name:     body.Name,
		Username: body.Username,
		Password: body.Password,
		Email:    body.Email,
		Avatar:   body.Avatar,
	})

	if err != nil {
		ctx.JSON(resp.ParseErr(err))
		return
	}
	ctx.JSON(resp.Succ(info.ID))
}

// Update
//
//	@Summary	修改管理员信息
//	@Tags		管理后台-管理员用户
//	@Accept		json
//	@Produce	json
//	@Param		req	body		api.AdminUserUpdateReq	true	"body"
//	@Success	200	{object}	resp.Result				"resp"
//	@Router		/api/admin/user/update/info [post]
func (c *AdminUserController) Update(ctx *gin.Context) {
	var body api.AdminUserUpdateReq
	utils.BindBody(ctx, &body)

	if err := c.service.Update(body); err != nil {
		ctx.JSON(resp.ParseErr(err))
		return
	}
	ctx.JSON(resp.Succ(nil))
}

// UpdateStatus
//
//	@Summary	修改管理员状态(封禁/解封)
//	@Tags		管理后台-管理员用户
//	@Accept		json
//	@Produce	json
//	@Param		req	body		api.AdminUserUpdateStatusReq	true	"body"
//	@Success	200	{object}	resp.Result						"resp"
//	@Router		/api/admin/user/update/status [post]
func (c *AdminUserController) UpdateStatus(ctx *gin.Context) {
	var body api.AdminUserUpdateStatusReq
	utils.BindBody(ctx, &body)
	if err := c.service.UpdateStatus(body.ID, body.Status); err != nil {
		ctx.JSON(resp.ParseErr(err))
		return
	}
	ctx.JSON(resp.Succ(nil))
}

// Delete
//
//	@Summary	删除管理员
//	@Tags		管理后台-管理员用户
//	@Accept		json
//	@Produce	json
//	@Param		req	body		api.AdminUserDeleteReq	true	"body"
//	@Success	200	{object}	resp.Result				"resp"
//	@Router		/api/admin/user/delete [post]
func (c *AdminUserController) Delete(ctx *gin.Context) {
	var body api.AdminUserDeleteReq
	utils.BindBody(ctx, &body)
	if err := c.service.Delete(body.ID); err != nil {
		ctx.JSON(resp.ParseErr(err))
		return
	}
	ctx.JSON(resp.Succ(nil))
}

// ResetPassword
//
//	@Summary	重置管理员密码
//	@Tags		管理后台-管理员用户
//	@Accept		json
//	@Produce	json
//	@Param		req	body		api.AdminUserResetPasswordReq	true	"body"
//	@Success	200	{object}	resp.Result						"resp"
//	@Router		/api/admin/user/reset-password [post]
func (c *AdminUserController) ResetPassword(ctx *gin.Context) {
	var body api.AdminUserResetPasswordReq
	utils.BindBody(ctx, &body)
	currentUser := utils.GetCurrentAdminUser(ctx)

	if err := c.service.OnlyRootAdminUser(body.ID, currentUser.ID); err != nil {
		ctx.JSON(resp.ParamErr(err.Error()))
		return
	}

	if err := c.service.UpdatePassword(body.ID, body.Password); err != nil {
		ctx.JSON(resp.ParseErr(err))
		return
	}
	ctx.JSON(resp.Succ(nil))
}

// UpdateRole
//
//	@Summary	为管理用户更新角色
//	@Tags		管理后台-管理员用户
//	@Accept		json
//	@Produce	json
//	@Param		req	body		api.AdminUserUpdateRoleReq	true	"body"
//	@Success	200	{object}	resp.Result					"resp"
//	@Router		/api/admin/user/update/role [post]
func (c *AdminUserController) UpdateRole(ctx *gin.Context) {
	var body api.AdminUserUpdateRoleReq

	utils.BindBody(ctx, &body)

	if err := c.service.UpdateRole(body.UserId, body.RoleIds); err != nil {
		ctx.JSON(resp.ParseErr(err))
		return
	}

	ctx.JSON(resp.Succ(nil))
}

// CurrentInfo
//
//	@Summary	当前登陆者信息
//	@Tags		管理后台-用户认证
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	resp.Result{data=api.AdminUserInfoRes}	"用户详情"
//	@Router		/api/admin/user/current [get]
func (c *AdminUserController) CurrentInfo(ctx *gin.Context) {
	user := utils.GetCurrentAdminUser(ctx)
	ctx.JSON(resp.Succ(user))
}

// Upload
//
//	@Summary	文件上传
//	@Tags		管理后台-通用接口
//	@Accept		json
//	@Produce	json
//	@Param		file	formData	file						true	"文件"
//	@Success	200		{object}	resp.Result{data=string}	"文件地址"
//	@Router		/api/admin/user/upload [post]
func (c *AdminUserController) Upload(ctx *gin.Context) {
	user := utils.GetCurrentAdminUser(ctx)
	file, errF := ctx.FormFile("file")
	if errF != nil {
		ctx.JSON(resp.ParseErr(errF))
		return
	}
	filePath := strconv.Itoa(int(user.ID)) + "/" + file.Filename
	fmt.Println(config.Config.FileUploadDir)
	dst := path.Join(config.Config.FileUploadDir + "/" + filePath)
	fmt.Println(dst)
	// 上传文件至指定的完整文件路径
	if err := ctx.SaveUploadedFile(file, dst); err != nil {
		ctx.JSON(resp.ParseErr(err))
		return
	}
	ctx.JSON(resp.Succ("/" + consts.UploadFilePathName + "/" + filePath))
}

// UploadToAliOss
//
//	@Summary	文件上传至 阿里云 oss
//	@Tags		管理后台-通用接口
//	@Accept		json
//	@Produce	json
//	@Param		file	formData	file						true	"文件"
//	@Param		prefix	formData	string						true	"文件路径"
//	@Success	200		{object}	resp.Result{data=string}	"文件地址"
//	@Router		/api/admin/user/upload-oss [post]
func (c *AdminUserController) UploadToAliOss(ctx *gin.Context) {
	user := utils.GetCurrentAdminUser(ctx)
	file, errF := ctx.FormFile("file")
	prefix := strings.TrimSpace(ctx.PostForm("prefix"))
	if len(prefix) < 3 {
		ctx.JSON(resp.ParamErr(valid.ErrTransform(errors.New("prefix 长度不能小于 3"))))
		return
	}
	if prefix[0:1] == "/" {
		prefix = prefix[1:]
	}
	if prefix[len(prefix)-1:] == "/" {
		prefix = prefix[0 : len(prefix)-1]
	}
	if errF != nil {
		ctx.JSON(resp.ParseErr(errF))
		return
	}
	envPath := "test"
	if config.Config.IsProd {
		envPath = "prod"
	}
	date := time.Now().Format("2006/01/02")
	filePath := envPath + "/file/" + strconv.Itoa(int(user.ID)) + "/" + prefix + "/" + date + "/" + file.Filename
	ossConfig := config.Config.AliOss
	fmt.Printf("ossconfig %+v \n", ossConfig)
	f, errO := file.Open()
	if errF != nil {
		ctx.JSON(resp.ParseErr(errO))
		return
	}
	errU := utils.UploadFileToAliOss(utils.UploadOssOptions{
		AccessKeyID:     ossConfig.AccessKeyID,
		AccessKeySecret: ossConfig.AccessKeySecret,
		Endpoint:        ossConfig.Endpoint,
		Bucket:          ossConfig.Bucket,
		File:            io.Reader(f),
		FilePath:        path.Join(filePath),
		Options:         nil,
	})
	if errU != nil {
		ctx.JSON(resp.ParseErr(errU))
		return
	}
	ctx.JSON(resp.Succ("https://" + ossConfig.Bucket + "." + ossConfig.Endpoint + "/" + filePath))
}
