package utils

import (
	"server/dal/model"
	"server/internal/consts"

	"github.com/gin-gonic/gin"
)

// GetCurrentAdminUser 获取当前登录的管理员用户信息
func GetCurrentAdminUser(ctx *gin.Context) *model.AdminUser {
	return ctx.MustGet(consts.ContextUserInfoKey).(*model.AdminUser)
}

// GetCurrentUser 获取当前登录的用户信息
func GetCurrentUser(ctx *gin.Context) *model.User {
	return ctx.MustGet(consts.ContextUserInfoKey).(*model.User)
}
