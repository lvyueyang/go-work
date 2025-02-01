package utils

import (
	"server/internal/lib/valid"
	"server/internal/utils/resp"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func BindBody(ctx *gin.Context, body any) error {
	err := ctx.ShouldBindBodyWith(body, binding.JSON)
	if err != nil {
		ctx.JSON(resp.ParamErr(valid.ErrTransform(err)))
		ctx.Abort()
		return err
	}
	return nil
}
