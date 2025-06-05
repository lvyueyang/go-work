package controller

import (
	"server/internal/api"
	"server/internal/consts"
	"server/internal/middleware"
	"server/internal/service"
	"server/internal/utils"
	"server/internal/utils/resp"

	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"golang.org/x/exp/slog"
)

type CaptchaController struct {
	service       *service.CaptchaService
	notifyService *service.NotifyService
}

func NewCaptchaController(e *gin.Engine) {
	c := CaptchaController{
		service:       service.NewCaptchaService(),
		notifyService: service.NewNotifyService(),
	}

	router := e.Group("/api/captcha")
	{
		router.POST("", c.Create)
		router.GET("/image", c.CreateImage)
		router.GET("/image/:key", c.ImageFile)
		router.GET("/clear", middleware.Auth(), c.Clear)
	}

	c.service.ClearExpiration()
	cr := cron.New()
	// 每隔五分钟清除一次过期验证码
	cr.AddFunc("@every 5m", func() {
		c.service.ClearExpiration()
	})
	cr.Start()
}

// Create
//
//	@Summary	发送验证码
//	@Tags		验证码
//	@Accept		json
//	@Produce	json
//	@Param		req	body		api.CaptchaSendReq	true	"body"
//	@Success	200	{object}	resp.Result{data=int}	"验证码 ID"
//	@Router		/api/captcha [post]
func (c *CaptchaController) Create(ctx *gin.Context) {
	var body api.CaptchaSendReq
	if err := utils.BindBody(ctx, &body); err != nil {
		return
	}

	// 校验验证码
	if succ := captcha.VerifyString(body.CaptchaKey, body.CaptchaValue); !succ {
		ctx.JSON(resp.ParamErr("图形验证码不正确"))
		return
	}

	// 创建邮箱验证码
	res, err := c.service.Create(body.Type, body.Value, body.Scenes)
	if err != nil {
		ctx.JSON(resp.ParamErr("验证码创建失败"))
	}

	go func() {
		if body.Type == consts.CaptchaTypeEmail {
			cru := consts.CaptchaScenesMap[body.Scenes]
			if err := c.notifyService.SendCaptchaEmail(cru.Label, cru.EmailTitle, body.Value, res.Code); err != nil {
				slog.Error("邮件发送失败", "email", body.Value, "err", err)
			}
		}
	}()

	ctx.JSON(resp.Succ(res.ID))
}

// CreateImage
//
//	@Summary	获取图片验证码的 key
//	@Tags		验证码
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	resp.Result{data=string}	"验证码图片的key"
//	@Router		/api/captcha/image [get]
func (c *CaptchaController) CreateImage(ctx *gin.Context) {
	ctx.JSON(resp.Succ(captcha.New()))
}

// ImageFile
//
//	@Summary	获取验证码图片
//	@Tags		验证码
//	@Accept		png
//	@Produce	json
//	@Param		key	path		string	true	"图片验证码 key"
//	@Success	200	{object}	string	"图片文件流"
//	@Router		/api/captcha/image/{key} [get]
func (c *CaptchaController) ImageFile(ctx *gin.Context) {
	key := ctx.Param("key")

	captcha.WriteImage(ctx.Writer, key, consts.CaptchaWidth, consts.CaptchaHeight)
}

// Clear
//
//	@Summary	清除过期验证码
//	@Tags		验证码
//	@Accept		json
//	@Produce	json
//	@Param		X-Auth-Token	header		string			true	"Authentication token"
//	@Success	200				{object}	resp.Result{}	"请求结果"
//	@Router		/api/captcha/clear [get]
func (c *CaptchaController) Clear(ctx *gin.Context) {
	c.service.ClearExpiration()
}
