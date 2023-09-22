package utils

import (
	"github.com/gin-gonic/gin"
	"server/consts"
	"server/dal/model"
)

// GetCurrentAdminUser 获取当前登录的管理员用户信息
func GetCurrentAdminUser(ctx *gin.Context) *model.AdminUser {
	return ctx.MustGet(consts.ContextUserInfoKey).(*model.AdminUser)
}

// GetCurrentUser 获取当前登录的用户信息
func GetCurrentUser(ctx *gin.Context) *model.User {
	return ctx.MustGet(consts.ContextUserInfoKey).(*model.User)
}
