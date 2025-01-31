package app

import (
	"path"
	"server/config"
	"server/internal/consts"
	"server/internal/controller"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	New(gin *gin.Engine)
}

func New(r *gin.Engine) {
	r.Static(consts.UploadFilePathName, path.Join(config.Config.FileUploadDir))

	controller.NewSwaggerController(r)

	controller.NewHomeController(r)

	controller.NewUserController(r)
	controller.NewAuthController(r)
	controller.NewCaptchaController(r)

	controller.NewAdminUserController(r)
	controller.NewAdminAuthController(r)
	controller.NewAdminRoleController(r)

	controller.NewNewsController(r)

}
